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
$ cd bin/
$ ./manku
```

### Docker

```sh
$ docker build --no-cache -t manku-server:latest .
```

```sh
$ docker run --name manku-server -d manku-server:latest -p 5363:5363
```

## License

[BSD 2-Clause license](https://github.com/lindsuen/manku/blob/master/README.md)
