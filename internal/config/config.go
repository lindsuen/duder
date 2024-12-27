// manku - config.go
// Copyright (C) 2024 LindSuen <lindsuen@foxmail.com>
//
// Use of this source code is governed by a BSD 2-Clause license that can be
// found in the LICENSE file.

package config

import (
	"github.com/go-ini/ini"
)

type ServerConfig struct {
	ServerAddress string
	ServerPort    string
}

func (s *ServerConfig) InitServerConfig() {
	// Manku's default listening address is 0.0.0.0.
	s.ServerAddress = "0.0.0.0"
	// Manku's default listening port is 5363.
	s.ServerPort = "5363"
}

var (
	config *ini.File
	err    error
)

// ParseIni parses the config.ini file. The parameter fpath is the relative path to config.ini.
func ParseIni(fpath string) (*ServerConfig, error) {
	serverConfig := new(ServerConfig)
	serverConfig.InitServerConfig()
	config, err = ini.Load(fpath)
	if err != nil {
		return nil, err
	}
	serverConfig.ServerAddress = parseSessionKey("server", "address")
	serverConfig.ServerPort = parseSessionKey("server", "port")
	return serverConfig, nil
}

func parseSessionKey(s string, k string) string {
	return config.Section(s).Key(k).String()
}
