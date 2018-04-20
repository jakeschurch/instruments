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
	"time"

	"github.com/jakeschurch/instruments/internal/ordering"
)

func mockOrder() *Order {
	return &Order{
		Name:         "AAPL",
		Buy:          true,
		quotedMetric: quotedMetric{Price: NewPrice(10.00), Volume: NewVolume(10.00)},
		Logic:        Market,
		Status:       Open,
		timestamp:    time.Time{},

		ticker: ordering.NewOrderTicker(),
		filled: 0,
	}
}

func testTx(tx *Transaction) bool {
	if tx.Price == NewPrice(10) && tx.Volume == NewVolume(10) && !tx.Timestamp.IsZero() {
		return true
	}
	return false
}
func TestOrder_Transact(t *testing.T) {
	type args struct {
		price  Price
		volume Volume
	}
	tests := []struct {
		name string
		o    *Order
		args args
	}{
		{"base case", mockOrder(), args{NewPrice(10), NewVolume(10)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.Transact(tt.args.price, tt.args.volume); !testTx(got) {
				t.Errorf("Order.Transact() = %v", got)
			}
		})
	}
}
