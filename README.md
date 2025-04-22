# Manku

[![Commit activity](https://img.shields.io/github/commit-activity/m/lindsuen/manku)](https://github.com/lindsuen/manku/graphs/commit-activity)
[![build](https://img.shields.io/github/actions/workflow/status/lindsuen/manku/build.yml?branch=master)](https://github.com/lindsuen/manku/actions/workflows/build.yml)
[![GitHub Release](https://img.shields.io/github/v/release/lindsuen/manku)](https://github.com/lindsuen/manku/releases)
[![GitHub License](https://img.shields.io/github/license/lindsuen/manku)](https://github.com/lindsuen/manku/blob/master/README.md)

Fast File Service in Go.

## Start

```sh
$ git clone https://github.com/lindsuen/manku.git
$ cd manku/
```

### Binary

```sh
$ make build
```

```sh
$ mv bin/manku ./ && ./manku
```

### Docker

```sh
$ docker build --no-cache -t manku-server:latest .
```

```sh
$ docker run -p 5363:5363 --name manku-server -v ${TARGET_DIR}/data:/usr/local/manku/data -v ${TARGET_DIR}/upload:/usr/local/manku/upload -d manku-server:latest
```

## License

[BSD 2-Clause license](https://github.com/lindsuen/manku/blob/master/README.md)
