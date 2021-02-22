// Copyright 2019 The Go Language Server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !gojay

package protocol

import (
	"encoding/json"
	"testing"
)

func TestProgressParams(t *testing.T) {
	testProgressParams(t, json.Marshal, json.Unmarshal)
}

func TestWorkDoneProgressBegin(t *testing.T) {
	testWorkDoneProgressBegin(t, json.Marshal, json.Unmarshal)
}

func TestWorkDoneProgressReport(t *testing.T) {
	testWorkDoneProgressReport(t, json.Marshal, json.Unmarshal)
}

func TestWorkDoneProgressEnd(t *testing.T) {
	testWorkDoneProgressEnd(t, json.Marshal, json.Unmarshal)
}
