FROM debian:bookworm-slim

WORKDIR /usr/local/duder

RUN set -e && mkdir {config,static}

COPY ./bin/duder ./

COPY ./config/config.ini ./config/

COPY ./static/index.html ./static/

RUN set -e && chmod +x ./duder

EXPOSE 5363

CMD ["./duder"]
