# Source: https://github.com/aws/aws-cli/blob/v2/docker/Dockerfile
FROM public.ecr.aws/amazonlinux/amazonlinux:2 as installer
ARG AWS_CLI_VERSION=2.15.19
ENV AWS_CLI_VERSION=$AWS_CLI_VERSION
ARG EXE_FILENAME=awscliv2.zip

RUN curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64-${AWS_CLI_VERSION}.zip" -o "awscliv2.zip"

RUN yum update -y \
  && yum install -y unzip \
  && unzip $EXE_FILENAME \
  && ./aws/install --bin-dir /aws-cli-bin/

FROM mcr.microsoft.com/dotnet/runtime-deps:6.0-jammy
RUN apt update && \
    apt install less groff -y && \
    rm -rf /var/lib/apt/lists/*
COPY --from=installer /usr/local/aws-cli/ /usr/local/aws-cli/
COPY --from=installer /aws-cli-bin/ /usr/local/bin/

WORKDIR /aws
ENTRYPOINT ["/usr/local/bin/aws"]
