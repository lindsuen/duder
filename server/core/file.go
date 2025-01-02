// manku - file.go
// Copyright (C) 2025 LindSuen <lindsuen@foxmail.com>
//
// Use of this source code is governed by a BSD 2-Clause license that can be
// found in the LICENSE file.

package core

import "time"

type File struct {
	FileMetaData     MetaData
	DownloadLink     string
	TempDownloadLink string
	Tag              string
}

type MetaData struct {
	ID              string
	Name            string
	Path            string
	Size            string
	CreatedTime     time.Time // timestamp
	LastUpdatedTime time.Time // timestamp
	Hash            string    // sha256
}
