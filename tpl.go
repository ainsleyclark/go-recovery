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
	"bytes"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const (
	MainLayout = "layout.gohtml"
)

func (re Recover) ParseTemplates() (*template.Template, error) {
	tpl := template.New("")

	err := filepath.Walk("./web", func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, ".gohtml") {
			_, err = tpl.ParseFiles(path)
			if err != nil {
				log.Println(err)
			}
		}
		return err
	})

	if err != nil {
		return nil, err
	}

	return tpl, nil
}

func (re *Recover) Tpl() ([]byte, error) {
	tpl, err := re.ParseTemplates()
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	err = tpl.ExecuteTemplate(&buf, MainLayout, nil) // TODO: Parse data
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

