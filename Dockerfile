FROM alpine
MAINTAINER mlabouardy <mohamed@labouardy.com>

ENV PORT 3000
ENV DURATION 30
ENV AWS_ACCESS_KEY_ID access
ENV AWS_SECRET_ACCESS_KEY secret
ENV AWS_DEFAULT_REGION us-east-1

RUN apk update && apk add curl
RUN curl -L https://s3.us-east-1.amazonaws.com/komiser/1.0.0/linux/komiser -o /usr/bin/komiser && \
    chmod +x /usr/bin/komiser && \
    mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

EXPOSE $PORT
ENTRYPOINT komiser start --port $PORT --duration $DURATION
