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
	"net/http"
)

type TplData map[string]interface{}

type (
	// Data represents the main struct for sending back data to the
	// template for recovery.
	Data struct {
		Error      error
		StatusCode int
		Request    *http.Request
		//Stack      trace.Stack
		Debug bool
	}
	// Request represents the data obtained from the context with
	// detailed information about the http request made.
	Request struct {
		URL        string
		Method     string
		IP         string
		Referer    string
		DataLength int64
		Body       string
		Headers    map[string][]string
		Query      map[string][]string
		Cookies    []*http.Cookie
	}
)

const (
	// The amount of files in the stack to be retrieved.
	StackDepth = 200
	// How many files to move up in the runtime.Caller
	// before obtaining the stack.
	StackSkip = 2
)

// getData
//
// Retrieves all the necessary template data to show
// the recovery page. If a template executor has
// been set, the template file stack will be
// retrieved,
func (re *Recover) getData(r *http.Request) *Data {
	return &Data{
		//StatusCode: re.,
		Request: r,
		//Stack:      r.getStackData(),
		// TEMPORARY
		Debug: true,
	}
}

// getStackData
//
// Check if the template exec has been set, if it has
// retrieve the file stack for the template. and
// prepend it to the stack.
//func (r *Recover) getStackData() trace.Stack {
//	stack := r.tracer.Trace(StackDepth, StackSkip)
//	if r.config.TplExec != nil && r.config.TplFile != "" {
//		root := r.config.TplExec.Config().GetRoot()
//		ext := r.config.TplExec.Config().GetExtension()
//		path := root + "/" + r.config.TplFile + ext
//
//		stack.Prepend(&trace.File{
//			File:     path,
//			Line:     tplLineNumber(r.err),
//			Function: r.config.TplFile,
//			Contents: tplFileContents(path),
//			Language: "handlebars",
//		})
//	}
//	return stack
//}
