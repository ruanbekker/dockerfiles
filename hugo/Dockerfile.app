FROM ruanbekker/hugo:v0.69

LABEL description="Hugo Static Site Generator in Docker"
LABEL maintainer="Ruan Bekker <ruan@ruanbekker.com>"

ADD boot.sh /boot.sh
CMD ["sh", "/boot.sh"]
