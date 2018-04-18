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
)

func mockQuote() Quote {
	return Quote{
		Name:      "AAPL",
		Bid:       &quotedMetric{NewPrice(10), NewVolume(10)},
		Ask:       &quotedMetric{NewPrice(10), NewVolume(10)},
		Timestamp: time.Time{}}
}

func TestQuote_TotalAsk(t *testing.T) {
	q := mockQuote()

	type fields struct {
		q Quote
	}
	tests := []struct {
		name    string
		fields  fields
		want    Amount
		wantErr bool
	}{
		{"base case", fields{q}, 100 * 100, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &tt.fields.q

			got, err := q.TotalAsk()
			if (err != nil) != tt.wantErr {
				t.Errorf("Quote.TotalAsk() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Quote.TotalAsk() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuote_TotalBid(t *testing.T) {
	type fields struct {
		q Quote
	}
	tests := []struct {
		name    string
		fields  fields
		want    Amount
		wantErr bool
	}{
		{"base case", fields{mockQuote()}, 100 * 100, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &tt.fields.q

			got, err := q.TotalBid()
			if (err != nil) != tt.wantErr {
				t.Errorf("Quote.TotalBid() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Quote.TotalBid() = %v, want %v", got, tt.want)
			}
		})
	}
}
