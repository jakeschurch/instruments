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
	"reflect"
	"testing"
	"time"
)

func mockTx(buy bool) Transaction {
	return Transaction{
		"Google",
		buy,
		quotedMetric{NewPrice(15.00), NewVolume(20.00)},
		time.Time{},
	}
}
func mockHolding() *Holding {
	return &Holding{
		Name: "Google", Volume: NewVolume(20.00),
		Buy: TxMetric{NewPrice(15.00), time.Time{}},
	}
}

func TestBuy(t *testing.T) {

	type args struct {
		tx Transaction
	}
	tests := []struct {
		name    string
		args    args
		want    *Holding
		wantErr bool
	}{
		{"base case", args{mockTx(true)}, mockHolding(), false},
		{"err case", args{mockTx(false)}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Buy(tt.args.tx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Buy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Buy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrice_Avg(t *testing.T) {
	var price = NewPrice(10.00)

	type args struct {
		n          uint
		quotePrice Price
	}
	tests := []struct {
		name string
		p    *Price
		args args
		want Price
	}{
		{
			"base case",
			&price,
			args{1, NewPrice(20.00)},
			NewPrice(15.00)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.Avg(tt.args.n, tt.args.quotePrice)
		})
	}
}
