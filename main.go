// manku - main.go
// Copyright (C) 2024 LindSuen <lindsuen@foxmail.com>
//
// Use of this source code is governed by a BSD 2-Clause license that can be
// found in the LICENSE file.

package main

import (
	"log"
	"net/http"

	badger "github.com/dgraph-io/badger/v4"
	"github.com/labstack/echo/v4"
)

func main() {
	// Open the Badger database located in the /tmp/badger directory.
	// It will be created if it doesn't exist.
	db, err := badger.Open(badger.DefaultOptions("data"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// Your code here…

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
