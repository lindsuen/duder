// manku - download_file.go
// Copyright (C) 2025 LindSuen <lindsuen@foxmail.com>
//
// Use of this source code is governed by a BSD 2-Clause license that can be
// found in the LICENSE file.

package handler

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/lindsuen/manku/internal/db"
	"github.com/lindsuen/manku/server/core"
)

func DownloadFile(c echo.Context) error {
	fileId := c.QueryParam("fileid")
	file := new(core.File)
	value, _ := base64.RawURLEncoding.DecodeString(string(db.Get([]byte(fileId))))
	err := json.Unmarshal(value, &file)
	if err != nil {
		return err
	}
	if !fileIsExist(file.Path) {
		return c.String(http.StatusNotFound, "The file is not found.")
	}
	return c.Attachment(file.Path, file.Name)
}

// fileIsExist
func fileIsExist(p string) bool {
	_, err := os.Stat(p)
	return err == nil
}
