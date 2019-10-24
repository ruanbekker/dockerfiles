FROM alpine

RUN apk --no-cache add curl bash git openssl ncurses  \
    && curl -LO https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl \
    && curl -LO https://raw.githubusercontent.com/ahmetb/kubectx/master/kubectx \
    && curl -LO https://raw.githubusercontent.com/ahmetb/kubectx/master/kubens \
    && chmod +x ./kube* \
    && mv ./kube* /usr/local/bin/
