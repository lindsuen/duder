# manku - Makefile
# Copyright (C) 2024 LindSuen <lindsuen@foxmail.com>
#
# Use of this source code is governed by a BSD 2-Clause license that can be
# found in the LICENSE file.

APP := manku
DIR := bin

.PHONY: all clean build linux

all: linux

clean:
	@if [ -d ${DIR} ]; then rm -rf ${DIR}/*; else exit 0; fi

build:
	go build -o ${DIR}/${APP} -ldflags "-s -w" .

linux:
	@# linux-amd64:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${DIR}/${APP} -ldflags "-s -w" .
	@cd ${DIR}/ && tar -zcf ${APP}-linux_amd64.tar.gz ${APP} && rm -rf ${APP} && cd ../
	@# linux-arm64:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o ${DIR}/${APP} -ldflags "-s -w" .
	@cd ${DIR}/ && tar -zcf ${APP}-linux_arm64.tar.gz ${APP} && rm -rf ${APP} && cd ../
