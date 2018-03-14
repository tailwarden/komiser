FROM alpine
MAINTAINER mlabouardy <mohamed@labouardy.com>
WORKDIR /app

ENV PORT 3000
ENV DURATION 30

RUN curl -O 
CMD ["komiser", "--port", $PORT, "--duration", $DURATION]