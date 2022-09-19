FROM amazon/aws-cli:2.7.30

ENV TERRAFORM_VERSION 0.13.5
ENV HELM_VERSION 3.9.4

RUN yum install jq gzip vim tar git unzip wget -y 

RUN cd /tmp/ && wget https://releases.hashicorp.com/terraform/$TERRAFORM_VERSION/terraform_"$TERRAFORM_VERSION"_linux_amd64.zip \
    && unzip terraform_"$TERRAFORM_VERSION"_linux_amd64.zip \
    && chmod +x terraform && mv terraform /usr/local/bin/terraform \
    && wget https://get.helm.sh/helm-v$HELM_VERSION-linux-amd64.tar.gz \
    && tar -xf helm-v$HELM_VERSION-linux-amd64.tar.gz \
    && mv linux-amd64/helm /usr/bin/helm \
    && rm -rf linux-amd64 \
    && curl -LO https://storage.googleapis.com/kubernetes-release/release/`curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt`/bin/linux/amd64/kubectl \
    && chmod +x ./kubectl \
    && mv ./kubectl /usr/local/bin/kubectl \ 
    && yum clean all \
    && rm -rf /var/cache/yum

ENTRYPOINT ["sh"]
