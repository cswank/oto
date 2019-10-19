// Copyright 2019 The Oto Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build !go1.13 !wasm

package oto

import (
	"syscall/js"
)

func float32SliceToTypedArray(s []float32) (js.Value, func()) {
	// Note that TypedArrayOf cannot work correcly on Wasm.
	// See https://github.com/golang/go/issues/31980

	a := js.TypedArrayOf(s)
	return a.Value, func() { a.Release() }
}

func isAudioWorkletAvailable() bool {
	// float32SliceToTypedArray's freeing function must be called. However, it is impossible to pass Float32Array 
	// to worklet without calling them. Forbid Audio Worklet on Go 1.12 or older.
	return false
}
