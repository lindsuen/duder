// manku - main.go
// Copyright (C) 2024 LindSuen <lindsuen@foxmail.com>
//
// Use of this source code is governed by a BSD 2-Clause license that can be
// found in the LICENSE file.

package main

import (
	"log"

	s "github.com/lindsuen/manku/server"
)

// It's the startup portal for Manku server.
func main() {
	err := s.ServerStart()
	if err != nil {
		log.Fatalln(err)
	}
}
