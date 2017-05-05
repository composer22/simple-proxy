# simple-proxy docker file

FROM alpine:latest

ENV BACKEND_SCHEME=http
ENV BACKEND_TARGET=

RUN apk upgrade --update \
  && apk --update-cache update \
  && apk add --update bash curl \
  && rm -rf /var/cache/apk/* /tmp/*

ADD simple-proxy-linux-amd64 /simple-proxy
ADD docker-entrypoint.sh /docker-entrypoint.sh

EXPOSE 80
CMD []
ENTRYPOINT ["/docker-entrypoint.sh"]
