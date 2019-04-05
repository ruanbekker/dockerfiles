Create the network:

```
$ docker network create container-net
```

Create the ldap server:

```
$ docker run --name ldap-service --network container-net --hostname ldap-service --detach osixia/openldap:1.1.8
```

Create the ldap ui:

```
$ docker run --name phpldapadmin-service --network container-net --publish 6443:443 --hostname phpldapadmin-service --link ldap-service:ldap-host --env PHPLDAPADMIN_LDAP_HOSTS=ldap-host --detach osixia/phpldapadmin:0.7.2
```

## Using a different ldap service

If you are using a different ldap backend:

```
$ docker run -d --name ldap --network container-net -p 389:389 -e SLAPD_PASSWORD=password -e SLAPD_DOMAIN=ldap.localhost.net dinkel/openldap
```

Provision the frontend:

```
$ docker run --network container-net -p 6443:443 --env PHPLDAPADMIN_LDAP_HOSTS=ldap --detach osixia/phpldapadmin:0.7.2
```

Login with this cn:

```
cn=admin,dc=ldap,dc=localhost,dc=net
```

![image](https://user-images.githubusercontent.com/567298/55215659-33dcd200-5203-11e9-8a03-28fd260326fd.png)


References:

- https://github.com/osixia/docker-phpLDAPadmin
- https://github.com/osixia/docker-openldap
- https://github.com/dinkel/docker-openldap
- https://github.com/admiralobvious/flask-simpleldap
- https://github.com/openfrontier/docker-ldap-ssp
- https://github.com/openfrontier/openldap-docker
