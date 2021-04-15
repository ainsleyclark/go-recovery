// Copyright 2020 The Go Recovery Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package recovery

import (
	"net"
	"net/http"
	"os"
	"strings"
)

type Handler struct {
	Debug bool
}

type Recovery interface {
	Recover(cfg Config) []byte
	HttpRecovery(next http.Handler) http.Handler
}

// Recover defines
type Recover struct {
	cfg Config
}

type Config struct {
	Code    int
	Error   interface{}
	TplFile string
	Data    TplData

	// We need the files where the error pages are stored
	// error-400.cms
	//
}

func New(Debug bool) *Handler {
	return &Handler{}
}

// Recover - TODO
func (h *Handler) Recover(cfg Config) []byte {
	r := Recover{
		cfg: cfg,
	}

	tpl, err := r.Tpl()
	if err != nil {
		return nil
	}

	return tpl
}


func Static(w http.ResponseWriter, r *http.Request) {
	http.Handle("/static", http.FileServer(http.Dir("./static/")))
}

// HTTPRecovery hello
func (h *Handler) HTTPRecovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}
				// If the connection is dead, we can't write a status to it.
				// Otherwise we will send the recover data back.
				if brokenPipe {
					//ctx.Error(err.(error)) //nolint: errcheck
					//ctx.Abort()
					return
				} else {
					//b := h.Recover(Config{
					//	Context: ctx,
					//	Error:   err,
					//})
					http.Error(w, "Content-Type header must be application/json", http.StatusInternalServerError)

					return
				}
			}
		}()
		next.ServeHTTP(w, r)
	})
}
