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

package main

import (
	"fmt"
	"github.com/ainsleyclark/go-recovery"
	"html/template"
	"net/http"
)

const (
	Port = ":8080"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/static", recovery.Static)
	err := http.ListenAndServe(Port, nil)

	if err != nil {
		return
	}
	fmt.Println("Server started and listening on port: " + Port)
}

// handler
func handler(w http.ResponseWriter, r *http.Request) {
	tpl := template.New("")

	files, err := tpl.ParseGlob("tpl/*.gohtml")
	if err != nil {

		bytes := recovery.New(true).Recover(recovery.Config{
			Code:    http.StatusInternalServerError,
			Error:   err,
			TplFile: "",
			Data:    nil,
		})

		fmt.Println(files)

		w.Write(bytes)
		//w.WriteHeader(http.StatusInternalServerError)

		return
	}

	err = files.ExecuteTemplate(w, "test.gohtml", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	//if err != nil {
	//	fmt.Println(err)
	//

	//
	//}
	fmt.Println("Hit handler")
}
