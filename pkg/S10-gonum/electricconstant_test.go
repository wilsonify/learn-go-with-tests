// Code generated by "go generate gonum.org/v1/gonum/unit/constant”; DO NOT EDIT.

// Copyright ©2019 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gonum_examples

import (
	"fmt"
	"testing"
)

func TestElectricConstantFormat(t *testing.T) {
	t.Parallel()
	for _, test := range []struct {
		format string
		want   string
	}{
		{"%v", "8.854187817620389e-12 A^2 s^4 kg^-1 m^-3"},
		{"%.1v", "9e-12 A^2 s^4 kg^-1 m^-3"},
		{"%50.1v", "                          9e-12 A^2 s^4 kg^-1 m^-3"},
		{"%50v", "          8.854187817620389e-12 A^2 s^4 kg^-1 m^-3"},
		{"%1v", "8.854187817620389e-12 A^2 s^4 kg^-1 m^-3"},
		{"%#v", "constant.electricConstantUnits(8.854187817620389e-12)"},
		{"%s", "%!s(constant.electricConstantUnits=8.854187817620389e-12 A^2 s^4 kg^-1 m^-3)"},
	} {
		got := fmt.Sprintf(test.format, ElectricConstant)
		if got != test.want {
			t.Errorf("Format %q: got: %q want: %q", test.format, got, test.want)
		}
	}
}
