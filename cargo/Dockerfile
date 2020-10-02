FROM debian:stretch-slim
ENV PATH ${HOME}/.cargo/bin:${PATH}

RUN apt-get update && apt-get install curl build-essential libssl-dev pkg-config -y
RUN curl https://sh.rustup.rs -sSf | sh
# cargo install --features=ssl websocat
CMD ["/bin/bash"]
