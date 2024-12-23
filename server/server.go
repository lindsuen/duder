// manku - server.go
// Copyright (C) 2024 LindSuen <lindsuen@foxmail.com>
//
// Use of this source code is governed by a BSD 2-Clause license that can be
// found in the LICENSE file.

package server

import (
	"github.com/labstack/echo/v4"
	"github.com/lindsuen/manku/internal/config"
	"github.com/lindsuen/manku/server/route"
)

type MankuServer struct {
	MankuServerInstance *echo.Echo
	MankuListenAddress  string
}

func NewMankuServer() *MankuServer {
	c := config.ParseIni("config/config.ini")
	s := new(MankuServer)
	s.MankuServerInstance = echo.New()
	s.MankuListenAddress = c.ServerAddress + ":" + c.ServerPort
	return s
}

func ServerStart() {
	mankuServer := NewMankuServer()
	e := mankuServer.MankuServerInstance
	route.LoadEchoRoute(e)
	e.Logger.Fatal(e.Start(mankuServer.MankuListenAddress))
}
