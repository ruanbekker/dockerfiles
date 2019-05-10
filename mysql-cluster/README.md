# mysql-cluster

Resource:
- https://github.com/toughIQ/docker-mariadb-cluster

## What does it do

1. MariaDB Cluster
2. DB Loadbalancer
3. MySQL Client that makes backups

## Demo

### Deploy

```
$ docker stack deploy -c docker-compose.yml galera
Creating network dbnet
Creating service galera_dbcluster
Creating service galera_dblb
Creating service galera_dbclient
```

Verify:

```
docker service ls
ID                  NAME                MODE                REPLICAS            IMAGE                            PORTS
jm7p70qre72u        galera_dbclient     replicated          1/1                 alpine:latest
p8kcr5y7szte        galera_dbcluster    replicated          1/1                 toughiq/mariadb-cluster:latest
1hu3oxhujgfm        galera_dblb         replicated          1/1                 toughiq/maxscale:latest          *:3306->3306/tcp
```

### Backup Client


```
$ docker exec -it $(docker ps -f name=galera_dbclient -q) find /data/
/data/
/data/2019-05-10
/data/2019-05-10/db_backup-2019-05-10T10.05.sql.gz
```

### Populate Data:

```
$ docker exec -it $(docker ps -f name=galera_dbclient -q) mysql -uroot -ppassword -h dblb
MySQL [(none)]> create table mydb.foo (name varchar(10));
MySQL [(none)]> insert into mydb.foo values('ruan');
MySQL [(none)]> exit
```

### Scale the Cluster:

```
$ docker service scale galera_dbcluster=3
galera_dbcluster scaled to 3
overall progress: 3 out of 3 tasks
1/3: running   [==================================================>]
2/3: running   [==================================================>]
3/3: running   [==================================================>]
verify: Service converged
```

Verify:

```
$ docker service ls
ID                  NAME                MODE                REPLICAS            IMAGE                            PORTS
jm7p70qre72u        galera_dbclient     replicated          1/1                 alpine:latest
p8kcr5y7szte        galera_dbcluster    replicated          3/3                 toughiq/mariadb-cluster:latest
1hu3oxhujgfm        galera_dblb         replicated          1/1                 toughiq/maxscale:latest          *:3306->3306/tcp
```

Test:

```
$ docker exec -it $(docker ps -f name=galera_dbclient -q) mysql -uroot -ppassword -h dblb -e'select * from mydb.foo;'
+------+
| name |
+------+
| ruan |
+------+
```

Simulate a node failure:

```
$ docker kill 9e336032ab52
$ docker service ls
ID                  NAME                MODE                REPLICAS            IMAGE                            PORTS
jm7p70qre72u        galera_dbclient     replicated          1/1                 alpine:latest
p8kcr5y7szte        galera_dbcluster    replicated          2/3                 toughiq/mariadb-cluster:latest
1hu3oxhujgfm        galera_dblb         replicated          1/1                 toughiq/maxscale:latest          *:3306->3306/tcp
```

While the container is provisioning, as we have 2 out of 3 running containers, read the data 3 times so test that the round robin queries dont hit the affected container (the dblb wont route traffic to the affected container):

```
$ docker exec -it $(docker ps -f name=galera_dbclient -q) mysql -uroot -ppassword -h dblb -e'select * from mydb.foo;'
+------+
| name |
+------+
| ruan |
+------+

$ docker exec -it $(docker ps -f name=galera_dbclient -q) mysql -uroot -ppassword -h dblb -e'select * from mydb.foo;'
+------+
| name |
+------+
| ruan |
+------+

$ docker exec -it $(docker ps -f name=galera_dbclient -q) mysql -uroot -ppassword -h dblb -e'select * from mydb.foo;'
+------+
| name |
+------+
| ruan |
+------+
```

Verify that the 3rd container has checked in:

```
$ docker service ls
ID                  NAME                MODE                REPLICAS            IMAGE                            PORTS
jm7p70qre72u        galera_dbclient     replicated          1/1                 alpine:latest
p8kcr5y7szte        galera_dbcluster    replicated          3/3                 toughiq/mariadb-cluster:latest
1hu3oxhujgfm        galera_dblb         replicated          1/1                 toughiq/maxscale:latest          *:3306->3306/tcp
```

### How to restore?

im deleting the db to simulate the scenario

```
$ docker exec -it $(docker ps -f name=galera_dbclient -q) sh
> mysql -uroot -ppassword -h dblb -e'drop database mydb;'
```

ensure the db is not present:

```
> mysql -uroot -ppassword -h dblb -e'select * from mydb.foo;'
ERROR 1146 (42S02) at line 1: Table 'mydb.foo' doesn't exist
```

Find the archive and extract:

```
> find /data/
/data/
/data/2019-05-10
/data/2019-05-10/db_backup-2019-05-10T10.05.sql.gz

> gunzip /data/2019-05-10/db_backup-2019-05-10T10.05.sql.gz
```

Restore to mysql:

```
> mysql -uroot -ppassword -h dblb < /data/2019-05-10/db_backup-2019-05-10T10.05.sql
```

Test:

```
> mysql -uroot -ppassword -h dblb -e'select * from mydb.foo;'
+------+
| name |
+------+
| ruan |
+------+
```
