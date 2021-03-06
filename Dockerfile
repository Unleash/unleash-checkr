FROM golang:1.14-alpine AS base

RUN apk add --no-cache bash curl docker-cli git mercurial make

ENTRYPOINT ["/entrypoint.sh"]
CMD [ "-h" ]

COPY scripts/entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

COPY unleash-checkr /bin/unleash-checkr