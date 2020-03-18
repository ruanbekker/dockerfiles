### About

MySQL Client on Docker

### MySQL Client as a Docker Container

Helper function:

```
mysql(){
  docker run -it ruanbekker/mysql-client "$@"
}
```

## Usage

```
$ source ~/.functions
$ mysql -h example.com -u user -p dbname
mysql> 
```
