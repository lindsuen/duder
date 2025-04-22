FROM debian:bookworm-slim

WORKDIR /usr/local/manku

RUN set -e && mkdir config

COPY ./bin/manku ./
COPY ./config/config.ini ./config

RUN set -e && apt-get update && chmod +x ./manku

EXPOSE 5363

CMD ["./manku"]
