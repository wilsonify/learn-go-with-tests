// Code generated by "go generate gonum.org/v1/gonum/unit; DO NOT EDIT.

// Copyright ©2019 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gonum_examples

import (
	"fmt"
	"testing"
)

func TestMagneticFluxDensity(t *testing.T) {
	t.Parallel()
	for _, value := range []float64{-1, 0, 1} {
		var got MagneticFluxDensity
		err := got.From(MagneticFluxDensity(value).Unit())
		if err != nil {
			t.Errorf("unexpected error for %T conversion: %v", got, err)
		}
		if got != MagneticFluxDensity(value) {
			t.Errorf("unexpected result from round trip of %T(%v): got: %v want: %v", got, value, got, value)
		}
		if got != got.MagneticFluxDensity() {
			t.Errorf("unexpected result from self interface method call: got: %#v want: %#v", got, value)
		}
		err = got.From(ether(1))
		if err == nil {
			t.Errorf("expected error for ether to %T conversion", got)
		}
	}
}

func TestMagneticFluxDensityFormat(t *testing.T) {
	t.Parallel()
	for _, test := range []struct {
		value  MagneticFluxDensity
		format string
		want   string
	}{
		{1.23456789, "%v", "1.23456789 T"},
		{1.23456789, "%.1v", "1 T"},
		{1.23456789, "%20.1v", "                 1 T"},
		{1.23456789, "%20v", "        1.23456789 T"},
		{1.23456789, "%1v", "1.23456789 T"},
		{1.23456789, "%#v", "unit.MagneticFluxDensity(1.23456789)"},
		{1.23456789, "%s", "%!s(unit.MagneticFluxDensity=1.23456789 T)"},
	} {
		got := fmt.Sprintf(test.format, test.value)
		if got != test.want {
			t.Errorf("Format %q %v: got: %q want: %q", test.format, test.value, got, test.want)
		}
	}
}
