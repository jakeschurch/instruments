// Copyright (c) 2017 Jake Schurch
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package instruments

import (
	"testing"
)

func TestDivide(t *testing.T) {
	top := NewPrice(20.00)
	bottom := NewPrice(10.00)

	type args struct {
		top    Price
		bottom Price
	}
	tests := []struct {
		name string
		args args
		want Amount
	}{
		{"base case", args{top, bottom}, 200},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Divide(tt.args.top, tt.args.bottom); got != tt.want {
				t.Errorf("Divide() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestPrice_String(t *testing.T) {
	type fields struct {
		price Price
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"10", fields{10 * 100}, "$10.00"},
		{"100", fields{100 * 100}, "$100.00"},
		{"1k", fields{1000 * 100}, "$1,000.00"},
		{"10k", fields{10000 * 100}, "$10,000.00"},
		{"100k", fields{100000 * 100}, "$100,000.00"},
		{"1m", fields{1000000 * 100}, "$1,000,000.00"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.fields.price
			if got := c.String(); got != tt.want {
				t.Errorf("Price.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewPrice(t *testing.T) {
	type args struct {
		f float64
	}
	tests := []struct {
		name string
		args args
		want Price
	}{
		{"base case", args{10}, 1000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPrice(tt.args.f); got != tt.want {
				t.Errorf("NewPrice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewVolume(t *testing.T) {
	type args struct {
		f float64
	}
	tests := []struct {
		name string
		args args
		want Volume
	}{
		{"base case", args{10.00}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewVolume(tt.args.f); got != tt.want {
				t.Errorf("NewVolume() = %v, want %v", got, tt.want)
			}
		})
	}
}
