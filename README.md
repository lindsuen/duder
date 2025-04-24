# Duder

[![Commit activity](https://img.shields.io/github/commit-activity/m/lindsuen/duder)](https://github.com/lindsuen/duder/graphs/commit-activity)
[![build](https://img.shields.io/github/actions/workflow/status/lindsuen/duder/build.yml?branch=master)](https://github.com/lindsuen/duder/actions/workflows/build.yml)
[![GitHub Release](https://img.shields.io/github/v/release/lindsuen/duder)](https://github.com/lindsuen/duder/releases)
[![GitHub License](https://img.shields.io/github/license/lindsuen/duder)](https://github.com/lindsuen/duder/blob/master/README.md)

Fast File Service in Go.

## Start

```sh
$ git clone https://github.com/lindsuen/duder.git
$ cd duder/
```

### Binary

```sh
$ make build
```

```sh
$ mv bin/duder ./ && ./duder
```

### Docker

```sh
$ docker build --no-cache -t duder-server:latest .
```

```sh
$ docker run -p 5363:5363 --name duder-server -v ${TARGET_DIR}/data:/usr/local/duder/data -v ${TARGET_DIR}/upload:/usr/local/duder/upload -d duder-server:latest
```

## License

[BSD 2-Clause license](https://github.com/lindsuen/duder/blob/master/README.md)
