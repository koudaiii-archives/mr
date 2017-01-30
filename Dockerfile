FROM alpine:3.5

RUN apk add --no-cache --update ca-certificates

COPY bin/mr /mr

ENTRYPOINT ["/mr"]
