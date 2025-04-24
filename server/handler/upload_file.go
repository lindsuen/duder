// duder - upload_file.go
// Copyright (C) 2025 LindSuen <lindsuen@foxmail.com>
//
// Use of this source code is governed by a BSD 2-Clause license that can be
// found in the LICENSE file.

package handler

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	cfg "github.com/lindsuen/duder/internal/config"
	"github.com/lindsuen/duder/internal/db"
	"github.com/lindsuen/duder/server/core"
)

type UploadResponse struct {
	Success bool       `json:"success"`
	Message []FileInfo `json:"message"`
}

type FileInfo struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Size        int64  `json:"size"`
	Path        string `json:"path"`
	CreatedTime int64  `json:"createdTime"`
	Hash        string `json:"hash"`
}

func UploadFile(c echo.Context) error {
	response := new(UploadResponse)
	response.Success = true
	response.Message = []FileInfo{}
	fileInfo := new(FileInfo)

	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	formFiles := form.File["files"]
	for _, fileHeader := range formFiles {
		multiFile, err := fileHeader.Open()
		if err != nil {
			return err
		}
		defer multiFile.Close()

		fileName := fileHeader.Filename
		fileSize := fileHeader.Size
		if fileSize > int64(parseMaxLength(cfg.Config.MaxLength)) {
			log.Println("The file " + fileName + " is too large.")
			continue
		}

		coreFile := new(core.File)
		coreFile.SetFileID()
		coreFile.SetFileName(fileName)
		coreFile.SetFileSize(fileSize)
		coreFile.SetFileCreatedTime()

		storagePath := createDateDir(cfg.Config.StoragePath) + "/" + setLocalFileName(fileName, coreFile.CreatedTime)
		file, err := os.Create(storagePath)
		if err != nil {
			log.Println(err)
		}
		defer file.Close()

		_, err = io.Copy(file, multiFile)
		if err != nil {
			return err
		}
		coreFile.SetFilePath(storagePath)
		coreFile.SetFileHash(file)

		fileInfo.ID = coreFile.ID
		fileInfo.Name = coreFile.Name
		fileInfo.Size = coreFile.Size
		fileInfo.Path = coreFile.Path
		fileInfo.CreatedTime = coreFile.CreatedTime
		fileInfo.Hash = coreFile.Hash

		key := []byte(fileInfo.ID)
		value, _ := json.Marshal(fileInfo)
		db.Set(key, []byte(base64.RawURLEncoding.EncodeToString(value)))

		response.Message = append(response.Message, *fileInfo)
	}
	return c.JSON(http.StatusOK, &response)
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

func setLocalFileName(name string, timestamp int64) string {
	nameByte := []byte(name)
	dataPrefix := fmt.Appendf(nil, "%x", sha1.Sum(nameByte))
	return string(dataPrefix[:29]) + strconv.FormatInt(timestamp, 10)
}

func parseMaxLength(s string) int {
	maxlength, _ := strconv.Atoi(s)
	return maxlength
}
