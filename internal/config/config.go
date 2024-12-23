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

type ServerConfig struct {
	ServerAddress string
	ServerPort    string
}

func (s *ServerConfig) NewServerConfig() {
	s.ServerAddress = "127.0.0.1"
	s.ServerPort = "5363"
}

var (
	config *ini.File
	err    error
)

// ParseIni parses the config.ini file.
func ParseIni(file string) *ServerConfig {
	cfg := new(ServerConfig)
	cfg.NewServerConfig()
	config, err = ini.Load(file)
	if err != nil {
		log.Fatalln("Fail to read the config file: ", err)
	}
	cfg.ServerAddress = parseSessionKey("server", "address")
	cfg.ServerPort = parseSessionKey("server", "port")
	return cfg
}

func parseSessionKey(s string, k string) string {
	return config.Section(s).Key(k).String()
}
