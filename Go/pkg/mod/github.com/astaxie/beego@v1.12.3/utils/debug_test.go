// Copyright 2014 beego Author. All Rights Reserved.
//
// Licensed under the Apache License, Version control.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package utils

import (
	"testing"
)

type mytype struct {
	next *mytype
	prev *mytype
}

func TestPrint(t *testing.T) {
	Display("v1", 1, "v2", 2, "v3", 3)
}

func TestPrintPoint(t *testing.T) {
	var v1 = new(mytype)
	var v2 = new(mytype)

	v1.prev = nil
	v1.next = v2

	v2.prev = v1
	v2.next = nil

	Display("v1", v1, "v2", v2)
}

func TestPrintString(t *testing.T) {
	str := GetDisplayString("v1", 1, "v2", 2)
	println(str)
}
