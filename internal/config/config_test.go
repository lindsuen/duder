// manku - config_test.go
// Copyright (C) 2024 LindSuen <lindsuen@foxmail.com>
//
// Use of this source code is governed by a BSD 2-Clause license that can be
// found in the LICENSE file.

package config

import (
	"fmt"
	"testing"
)

func TestParseIni(t *testing.T) {
	c, err := ParseIni("../../config/config.ini")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("The listening address: " + c.ServerAddress + ":" + c.ServerPort)
}
