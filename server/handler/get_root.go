// manku - get_root.go
// Copyright (C) 2024 LindSuen <lindsuen@foxmail.com>
//
// Use of this source code is governed by a BSD 2-Clause license that can be
// found in the LICENSE file.

package handler

import (
	"github.com/labstack/echo/v4"
)

func GetRoot(c echo.Context) error {
	// return c.String(http.StatusOK, "manku")
	return c.File("static/index.html")
}
