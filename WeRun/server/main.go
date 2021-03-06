/*
 * MIT License
 *
 * Copyright (c) 2017 Lin Hao.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the 'Software'), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED 'AS IS', WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

/*
 * Revision History:
 *     Initial: 2018/01/05      Lin Hao
 */

package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/xlstudio/wxbizdatacrypt"
)

// Req - structure of the request
type Req struct {
	EncryptedData string `json:"encryptedData"`
	IV            string `json:"iv"`
	SessionKey    string `json:"sessionKey"`
}

func main() {
	// 你的小程序ID
	AppID := ""

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/decode", func(c echo.Context) error {
		key := new(Req)
		if err := c.Bind(&key); err != nil {
			return c.JSON(http.StatusForbidden, err)
		}

		pc := wxbizdatacrypt.WxBizDataCrypt{AppID: AppID, SessionKey: key.SessionKey}
		result, err := pc.Decrypt(key.EncryptedData, key.IV, true)
		if err != nil {
			return c.JSON(http.StatusForbidden, err)
		}

		return c.JSON(http.StatusOK, result)
	})

	e.Start(":8888")
}
