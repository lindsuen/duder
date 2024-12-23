// manku - server.go
// Copyright (C) 2024 LindSuen <lindsuen@foxmail.com>
//
// Use of this source code is governed by a BSD 2-Clause license that can be
// found in the LICENSE file.

package server

import (
	"github.com/lindsuen/manku/server/middleware/log"

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
	server := new(MankuServer)
	server.MankuServerInstance = echo.New()
	server.MankuListenAddress = c.ServerAddress + ":" + c.ServerPort
	return server
}

func ServerStart() {
	mankuServer := NewMankuServer()
	inst := mankuServer.MankuServerInstance
	addr := mankuServer.MankuListenAddress
	route.LoadEchoRoute(inst)
	log.Logger(inst)
	inst.Logger.Fatal(inst.Start(addr))
}
