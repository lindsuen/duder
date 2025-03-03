// manku - config.go
// Copyright (C) 2024 LindSuen <lindsuen@foxmail.com>
//
// Use of this source code is governed by a BSD 2-Clause license that can be
// found in the LICENSE file.

package config

import (
	"log"

	"github.com/go-ini/ini"
)

var (
	config *ini.File
	err    error

	Config *ServerConfig
)

func init() {
	Config, err = ParseIni("config/config.ini")
	if err != nil {
		log.Fatalln(err)
	}
}

type ServerConfig struct {
	ServerAddress     string
	ServerPort        string
	ServerDataPath    string
	ServerStoragePath string
}

func (s *ServerConfig) initServerConfig() {
	// Manku's default listening address is "0.0.0.0".
	s.ServerAddress = "0.0.0.0"

	// Manku's default listening port is "5363".
	s.ServerPort = "5363"

	// Manku's default data path is "data".
	s.ServerDataPath = "data"

	// Manku's default storage path is "upload".
	s.ServerDataPath = "upload"
}

// ParseIni parses the config.ini file. The parameter fpath is the relative path to config.ini.
func ParseIni(fpath string) (*ServerConfig, error) {
	serverConfig := new(ServerConfig)
	serverConfig.initServerConfig()
	config, err = ini.Load(fpath)
	if err != nil {
		return nil, err
	}
	serverConfig.ServerAddress = parseSessionKey("server", "address")
	serverConfig.ServerPort = parseSessionKey("server", "port")
	serverConfig.ServerDataPath = parseSessionKey("server", "data_path")
	serverConfig.ServerStoragePath = parseSessionKey("server", "storage_path")
	return serverConfig, nil
}

func parseSessionKey(s string, k string) string {
	return config.Section(s).Key(k).String()
}
