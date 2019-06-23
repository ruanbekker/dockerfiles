# Passwordless SSH Container
Container that grants SSH to normal user for no auth

## Why not auth
purely for debugging / homelab purposes

### Auto Generated Usernames

Running the container with generated usernames:

```
$ docker run -it -p 2222:22 ruanbekker/ssh-noauth
SSH User is: user-c2ae30179dbd9f01be1a2b6c
```

SSH to the Container:

```
$ ssh -p 2222 user-c2ae30179dbd9f01be1a2b6c@localhost                                                                                                     130 ↵

+ Container Info:

    Name: de276bd33aa2

    CPU:  Intel(R) Core(TM) i5-5257U CPU @ 2.70GHz
    Memory: 1999M
    Disk: 58.4G
    Distro: Alpine:3.9.3

    CPU Load: 0.20, 0.16, 0.10
    Free Memory: 438M
    Free Disk: 58.4G

    Container Address: 172.17.0.3
    Public Address: x.x.x.x
    SSH User: user-c2ae30179dbd9f01be1a2b6c

de276bd33aa2:~$
```

### Specifying your own username

Running the container with your own username:

```
$ docker run -it -e SSH_USER=ruan -p 2222:22 ruanbekker/ssh-noauth
SSH User is: ruan
```

SSH to the container:

```
$ ssh -p 2222 ruan@localhost

+ Container Info:

    Name: 9b9b20b498d0

    CPU:  Intel(R) Core(TM) i5-5257U CPU @ 2.70GHz
    Memory: 1999M
    Disk: 58.4G
    Distro: Alpine:3.9.3

    CPU Load: 1.35, 0.42, 0.19
    Free Memory: 431M
    Free Disk: 58.4G

    Container Address: 172.17.0.3
    Public Address: xx.xx.xx.xx
    SSH User: ruan

9b9b20b498d0:~$
```
