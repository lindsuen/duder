// manku - server.go
// Copyright (C) 2024 LindSuen <lindsuen@foxmail.com>
//
// Use of this source code is governed by a BSD 2-Clause license that can be
// found in the LICENSE file.

package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lindsuen/manku/internal/config"
)

func ServerStart() {
	c := config.ParseIni("config/config.ini")
	listenAddress := c.ServerAddress + ":" + c.ServerPort
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "manku")
	})
	e.Logger.Fatal(e.Start(listenAddress))
}
