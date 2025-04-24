// duder - config.go
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
	Config *ServerConfig

	initFile *ini.File
	err      error
)

type ServerConfig struct {
	Address     string
	Port        string
	DataPath    string
	StoragePath string
	MaxLength   string
}

func InitServerConfig() {
	Config, err = parseInitFile("config/config.ini")
	if err != nil {
		log.Fatalln(err)
	}
}

// ParseIni parses the config.ini file. The parameter fpath is the relative path to config.ini.
func parseInitFile(fpath string) (*ServerConfig, error) {
	initFile, err = ini.Load(fpath)
	if err != nil {
		return nil, err
	}

	cfg := new(ServerConfig)
	cfg.Address = parseSessionKey("server", "address")
	cfg.Port = parseSessionKey("server", "port")
	cfg.DataPath = parseSessionKey("server", "data_path")
	cfg.StoragePath = parseSessionKey("server", "storage_path")
	cfg.MaxLength = parseSessionKey("server", "max_length")

	return cfg, nil
}

func parseSessionKey(s string, k string) string {
	return initFile.Section(s).Key(k).String()
}
