// manku - upload_file.go
// Copyright (C) 2025 LindSuen <lindsuen@foxmail.com>
//
// Use of this source code is governed by a BSD 2-Clause license that can be
// found in the LICENSE file.

package handler

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	cf "github.com/lindsuen/manku/internal/config"
)

func UploadFile(c echo.Context) error {
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	formFiles := form.File["files"]
	for _, file := range formFiles {
		f, err := file.Open()
		if err != nil {
			return err
		}
		defer f.Close()

		d, err := os.Create(createDateDir(cf.Config.ServerStoragePath) + "/" + file.Filename)
		if err != nil {
			log.Println(err)
		}
		defer d.Close()

		_, err = io.Copy(d, f)
		if err != nil {
			return err
		}
	}

	return c.String(http.StatusOK, "success")
}

func createDateDir(basePath string) string {
	subFolderName := time.Now().Format("20060102")
	folderPath := fmt.Sprint(basePath + "/" + subFolderName)
	_, err := os.Stat(folderPath)
	if os.IsNotExist(err) {
		os.MkdirAll(folderPath, 0777)
		os.Chmod(folderPath, 0777)
	}
	return folderPath
}
