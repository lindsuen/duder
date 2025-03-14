// manku - file.go
// Copyright (C) 2025 LindSuen <lindsuen@foxmail.com>
//
// Use of this source code is governed by a BSD 2-Clause license that can be
// found in the LICENSE file.

package core

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
)

type File struct {
	ID          string
	Name        string
	Size        int64
	Path        string
	CreatedTime int64  // timestamp
	Hash        string // sha256
}

func (d *File) SetFileID() {
	d.ID = uuid.New().String()
}

func (d *File) SetFileName(n string) {
	d.Name = n
}

func (d *File) SetFileSize(s int64) {
	d.Size = s
}

func (d *File) SetFilePath(p string) {
	d.Path = p
}

func (d *File) SetFileCreatedTime() {
	d.CreatedTime = time.Now().UnixMilli()
}

func (d *File) SetFileHash(f *os.File) {
	hash := sha256.New()
	_, err := io.Copy(hash, f)
	if err != nil {
		log.Println(err)
	}
	d.Hash = fmt.Sprintf("%x", hash.Sum(nil))
}
