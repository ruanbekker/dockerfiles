## Theme: Default

Make sure that `book.json` is not in your current working directory, as it's meant for installing the FAQ theme.

File structure:

```
$ find .
.
./docker/install-docker.md
./SUMMARY.md
./README.md
./aws/deploy-docker-swarm-on-aws.md
./aws/how-to-debug-a-lambda-function.md
```

Summary Page:

```
$ cat SUMMARY.md
# Summary

* [Docker](docker/README.md)
  * [Docker Swarm on AWS](aws/deploy-docker-swarm-on-aws.md)
  * [Install Docker](docker/install-docker.md)
* [AWS](aws/README.md)
  * [Debug Lambda on AWS](aws/how-to-debug-a-lambda-function.md)
```

Initialize:

```
$ docker run --rm -v "$PWD:/srv/gitbook" -p 4000:4000 ruanbekker/gitbook:3.2.3 gitbook init
```

Serve:

```
$ docker run --rm -v "$PWD:/srv/gitbook" -p 4000:4000 ruanbekker/gitbook:3.2.3 gitbook serve
```

Home Page:

<img width="1440" alt="image" src="https://user-images.githubusercontent.com/567298/60515523-10329780-9cdc-11e9-872f-829c0e97a1e2.png">


## Theme: FAQ

File structure:

```
$ find . | grep -v _book | grep -v node_modules
.
./docker/install-docker.md
./SUMMARY.md
./book.json
./README.md
./aws/deploy-docker-swarm-on-aws.md
./aws/how-to-debug-a-lambda-function.md
```

Summary Page:

```
$ cat SUMMARY.md
# Summary

* [Introduction](README.md)
* [Docker Swarm on AWS](aws/deploy-docker-swarm-on-aws.md)
* [Debug Lambda on AWS](aws/how-to-debug-a-lambda-function.md)
* [Install Docker](docker/install-docker.md)
```

Support for related pages:

```
$ cat aws/deploy-docker-swarm-on-aws.md
---
related:
    - docker/install-docker.md

---
# Deploy Docker on AWS

Tutorial showing you how to deploy docker on aws
```

Initialize:

```
$ docker run --rm -v "$PWD:/srv/gitbook" -p 4000:4000 ruanbekker/gitbook:3.2.3 gitbook init
```

Install Plugin:

```
$ docker run --rm -v "$PWD:/srv/gitbook" -p 4000:4000 ruanbekker/gitbook:3.2.3 gitbook install
```

Serve:

```
$ docker run --rm -v "$PWD:/srv/gitbook" -p 4000:4000 ruanbekker/gitbook:3.2.3 gitbook serve
```

FAQ Home Page:

<img width="1297" alt="image" src="https://user-images.githubusercontent.com/567298/60514561-0740c680-9cda-11e9-8cf3-ecee47a4aa18.png">

Search Result:

<img width="1251" alt="image" src="https://user-images.githubusercontent.com/567298/60514605-25a6c200-9cda-11e9-90a0-86df57c8959f.png">

Related Docs:

<img width="1264" alt="image" src="https://user-images.githubusercontent.com/567298/60514648-35bea180-9cda-11e9-8c9a-fb2a7dca5210.png">

## Resources:
- https://github.com/GitbookIO?utf8=%E2%9C%93&q=faq&type=&language=
- https://github.com/GitbookIO/theme-faq
- https://github.com/billryan/algorithm-exercise/tree/master/en
