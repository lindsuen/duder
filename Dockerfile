FROM debian:bookworm-slim

WORKDIR /usr/local/duder-server

RUN set -e && mkdir config static

COPY ./bin/duder ./

COPY ./config/duder.conf ./config/

COPY ./static/index.html ./static/

RUN set -e && chmod +x ./duder

EXPOSE 5363

CMD ["./duder"]
