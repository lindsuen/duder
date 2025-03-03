// manku - server.go
// Copyright (C) 2024 LindSuen <lindsuen@foxmail.com>
//
// Use of this source code is governed by a BSD 2-Clause license that can be
// found in the LICENSE file.

package server

import (
	"github.com/labstack/echo/v4"
	cf "github.com/lindsuen/manku/internal/config"
	"github.com/lindsuen/manku/server/middleware/logger"
	"github.com/lindsuen/manku/server/route"
)

type MankuServer struct {
	MankuServerInstance *echo.Echo
	MankuListenAddress  string
	MankuDataPath       string
	MankuStoragePath    string
}

func NewMankuServer() *MankuServer {
	server := new(MankuServer)
	server.MankuServerInstance = echo.New()
	server.MankuListenAddress = cf.Config.ServerAddress + ":" + cf.Config.ServerPort
	server.MankuDataPath = cf.Config.ServerDataPath
	server.MankuStoragePath = cf.Config.ServerStoragePath
	return server
}

// ServerStart can start the Manku server.
func ServerStart() error {
	mankuServer := NewMankuServer()
	i := mankuServer.MankuServerInstance
	addr := mankuServer.MankuListenAddress
	route.LoadRoutes(i)
	logger.LoadLogger(i)
	i.Logger.Fatal(i.Start(addr))
	return nil
}
