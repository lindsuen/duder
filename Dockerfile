FROM debian:bookworm-slim

WORKDIR /usr/local/manku

RUN set -e && mkdir config static

COPY ./bin/manku ./
COPY ./config/config.ini ./config/
COPY ./static/index.html ./static/

RUN set -e && chmod +x ./manku

EXPOSE 5363

CMD ["./manku"]
