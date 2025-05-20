// duder - server.go
// Copyright (C) 2024 LindSuen <lindsuen@foxmail.com>
//
// Use of this source code is governed by a BSD 2-Clause license that can be
// found in the LICENSE file.

package server

import (
	"github.com/labstack/echo/v4"
	cfg "github.com/lindsuen/duder/internal/config"
	badgerDB "github.com/lindsuen/duder/internal/db"
	"github.com/lindsuen/duder/server/middleware/logger"
	"github.com/lindsuen/duder/server/route"
)

type Server struct {
	Instance      *echo.Echo
	ListenAddress string
	DataPath      string
	StoragePath   string
}

func NewServer() *Server {
	cfg.InitServerConfig()

	s := new(Server)
	s.Instance = echo.New()
	s.ListenAddress = cfg.Config.Address + ":" + cfg.Config.Port
	s.DataPath = cfg.Config.DataPath
	s.StoragePath = cfg.Config.StoragePath

	return s
}

// Start can start the Duder server.
func Start() error {
	serv := NewServer()
	inst := serv.Instance
	addr := serv.ListenAddress

	_, err := badgerDB.Open(cfg.Config.DataPath)
	if err != nil {
		return err
	}
	defer badgerDB.Close()

	route.LoadEchoRoutes(inst)
	logger.LoadEchoLogger(inst)
	inst.Logger.Fatal(inst.Start(addr))

	return nil
}
