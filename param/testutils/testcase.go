// Copyright 2019 go-hip authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package testutils

import (
	"testing"

	"github.com/pascaldekloe/goe/verify"
	"github.com/wmnsk/go-hip/param"
)

// Serializeable is just for testing Params. Don't use this.
type Serializeable interface {
	Serialize() ([]byte, error)
	Len() int
}

// TestCase is just for testing Params. Don't use this.
type TestCase struct {
	Description string
	Structured  Serializeable
	Serialized  []byte
}

// DecodeFunc is just for testing Params. Don't use this.
type DecodeFunc func([]byte) (Serializeable, error)

// Run is just for testing Params. Don't use this.
func Run(t *testing.T, cases []TestCase, decode DecodeFunc) {
	t.Helper()

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			t.Run("Decode", func(t *testing.T) {
				v, err := decode(c.Serialized)
				if err != nil {
					t.Fatal(err)
				}

				if got, want := v, c.Structured; !verify.Values(t, "", got, want) {
					t.Fail()
				}
			})

			t.Run("Serialize", func(t *testing.T) {
				b, err := c.Structured.Serialize()
				if err != nil {
					t.Fatal(err)
				}

				if got, want := b, c.Serialized; !verify.Values(t, "", got, want) {
					t.Fail()
				}
			})

			t.Run("Len", func(t *testing.T) {
				if got, want := c.Structured.Len(), len(c.Serialized); got != want {
					t.Fatalf("got %v want %v", got, want)
				}
			})

			t.Run("Interface", func(t *testing.T) {
				// Ignore *Header and Generic in this tests.
				if _, ok := c.Structured.(*param.Header); ok {
					return
				}

				decoded, err := param.Decode(c.Serialized)
				if err != nil {
					t.Fatal(err)
				}

				if got, want := decoded.ParamType(), c.Structured.(param.Param).ParamType(); got != want {
					t.Fatalf("got %v want %v", got, want)
				}
			})
		})
	}
}
