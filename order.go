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
	"time"

	"github.com/jakeschurch/instruments/internal/timing"
)

// Order stores logic for transacting a stock.
type Order struct {
	Name string
	quotedMetric
	filled Volume

	Buy       bool
	Status    Status
	Logic     Logic
	timestamp time.Time
	ticker    *timing.OrderTicker
}

func (o *Order) timestampTx() time.Time {
	return o.timestamp.Add(o.ticker.Duration())
}

func (o *Order) Transact(p Price, v Volume) *Transaction {
	var tx *Transaction = &Transaction{
		Name:         o.Name,
		Buy:          o.Buy,
		quotedMetric: quotedMetric{o.Price, o.Volume},
		Timestamp:    o.timestampTx(),
	}
	o.filled -= v
	return tx
}

type Transaction struct {
	Name string
	Buy  bool
	quotedMetric
	Timestamp time.Time
}

// Status variables refer to a status of an order's execution.
type Status int

const (
	// Open indicates that an order has not been transacted.
	Open Status = iota // 0
	// Closed indicates that an order has been transacted.
	Closed
	// Cancelled indicates than an order was closed, but order was not transacted.
	Cancelled
)

// Logic is used to identify when the order should be executed.
type Logic int

const (
	Market Logic = iota // 0
	Limit
)
