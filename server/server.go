// manku - server.go
// Copyright (C) 2024 LindSuen <lindsuen@foxmail.com>
//
// Use of this source code is governed by a BSD 2-Clause license that can be
// found in the LICENSE file.

package server

import (
	"github.com/labstack/echo/v4"
	"github.com/lindsuen/manku/internal/config"
	"github.com/lindsuen/manku/server/middleware/logger"
	"github.com/lindsuen/manku/server/route"
)

type MankuServer struct {
	MankuServerInstance *echo.Echo
	MankuListenAddress  string
}

func NewMankuServer() (*MankuServer, error) {
	conf, err := config.ParseIni("config/config.ini")
	if err != nil {
		return nil, err
	}
	server := new(MankuServer)
	server.MankuServerInstance = echo.New()
	server.MankuListenAddress = conf.ServerAddress + ":" + conf.ServerPort
	return server, nil
}

// ServerStart can start the Manku server.
func ServerStart() error {
	mankuServer, err := NewMankuServer()
	if err != nil {
		return err
	}
	i := mankuServer.MankuServerInstance
	addr := mankuServer.MankuListenAddress
	route.LoadRoutes(i)
	logger.LoadLogger(i)
	i.Logger.Fatal(i.Start(addr))
	return nil
}
