FROM --platform=$BUILDPLATFORM alpine:3.16
ARG TARGETPLATFORM
ARG BUILDPLATFORM
ARG SEGMENT_WRITE_KEY
ARG VERSION
LABEL MAINTAINER="mlabouardy <mohamed@tailwarden.com>"

RUN echo "Running on $BUILDPLATFORM, building for $TARGETPLATFORM" > /log

ENV SEGMENT_WRITE_KEY $SEGMENT_WRITE_KEY
ENV VERSION $VERSION

COPY komiser /usr/bin/komiser
RUN chmod +x /usr/bin/komiser && \
    mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

EXPOSE $PORT
ENTRYPOINT ["komiser", "start"]
