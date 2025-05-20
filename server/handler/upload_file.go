// duder - upload_file.go
// Copyright (C) 2025 LindSuen <lindsuen@foxmail.com>
//
// Use of this source code is governed by a BSD 2-Clause license that can be
// found in the LICENSE file.

package handler

import (
	"crypto/sha1"
	"crypto/sha256"
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
	duder "github.com/lindsuen/duder/server/core"
)

type UploadResponse struct {
	FileList []FileInfo `json:"list"`
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
	response.FileList = []FileInfo{}

	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	formFiles := form.File["files"]
	for _, fileHeader := range formFiles {
		fileInfo := new(FileInfo)
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

		dFile := new(duder.File)
		dFile.SetFileID()
		dFile.SetFileName(fileName)
		dFile.SetFileSize(fileSize)
		dFile.SetFileCreatedTime()

		storagePath := createDateDir(cfg.Config.StoragePath) + "/" + setLocalFileName(fileName, dFile.CreatedTime)
		file, err := os.Create(storagePath)
		if err != nil {
			return err
		}
		defer file.Close()

		hash := sha256.New()
		_, err = io.Copy(file, io.TeeReader(multiFile, hash))
		if err != nil {
			return err
		}
		dFile.SetFilePath(storagePath)
		dFile.SetFileHash(fmt.Sprintf("%x", hash.Sum(nil)))

		fileInfo.ID = dFile.ID
		fileInfo.Name = dFile.Name
		fileInfo.Size = dFile.Size
		fileInfo.Path = dFile.Path
		fileInfo.CreatedTime = dFile.CreatedTime
		fileInfo.Hash = dFile.Hash

		key := []byte(fileInfo.ID)
		value, _ := json.Marshal(fileInfo)
		db.Set(key, []byte(base64.RawURLEncoding.EncodeToString(value)))

		response.FileList = append(response.FileList, *fileInfo)
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
