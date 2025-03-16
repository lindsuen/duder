// manku - db_test.go
// Copyright (C) 2025 LindSuen <lindsuen@foxmail.com>
//
// Use of this source code is governed by a BSD 2-Clause license that can be
// found in the LICENSE file.

package db

import (
	"fmt"
	"testing"
)

var key = []byte("")

func TestGet(t *testing.T) {
	_, err := Open("../../data")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(Get(key)))
}

func TestIteratorKeysAndValues(t *testing.T) {
	_, err := Open("../../data")
	if err != nil {
		fmt.Println(err)
	}
	IteratorKeysAndValues()
}
