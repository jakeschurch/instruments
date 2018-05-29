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

func mockQuote() Quote {
	return Quote{
		Name:      "AAPL",
		Bid:       QuotedMetric{NewPrice(10), NewVolume(10)},
		Ask:       QuotedMetric{NewPrice(10), NewVolume(10)},
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
		{"err case", fields{q}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &tt.fields.q

			if tt.name == "err case" {
				q.Ask.Price = 0
			}

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
		{"err case", fields{mockQuote()}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &tt.fields.q

			if tt.name == "err case" {
				q.Bid.Price = 0
			}

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

func Test_quotedMetric_Total(t *testing.T) {
	tests := []struct {
		name    string
		q       *QuotedMetric
		wantA   Amount
		wantErr bool
	}{
		{"base case", &QuotedMetric{NewPrice(10), NewVolume(10)}, 100 * 100, false},
		{"err case", &QuotedMetric{NewPrice(0), NewVolume(10)}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotA, err := tt.q.Total()
			if (err != nil) != tt.wantErr {
				t.Errorf("quotedMetric.Total() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotA, tt.wantA) {
				t.Errorf("quotedMetric.Total() = %v, want %v", gotA, tt.wantA)
			}
		})
	}
}

func TestQuote_FillOrder(t *testing.T) {
	type fields struct {
		Name      string
		Bid       QuotedMetric
		Ask       QuotedMetric
		Timestamp time.Time
	}
	type args struct {
		price Price
		vol   Volume
		buy   bool
		logic Logic
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Order
	}{
		{"base case",
			fields{
				Name:      "AAPL",
				Bid:       QuotedMetric{NewPrice(10.00), NewVolume(10)},
				Ask:       QuotedMetric{NewPrice(10.00), NewVolume(10)},
				Timestamp: time.Time{}},
			args{price: NewPrice(10.00), vol: NewVolume(10), buy: true, logic: Market},
			NewOrder("AAPL", true, Market, NewPrice(10.00), NewVolume(10), time.Time{})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Quote{
				Name:      tt.fields.Name,
				Bid:       tt.fields.Bid,
				Ask:       tt.fields.Ask,
				Timestamp: tt.fields.Timestamp,
			}
			if got := q.FillOrder(tt.args.price, tt.args.vol, tt.args.buy, tt.args.logic); !(got.Volume == tt.want.Volume && got.Price == tt.want.Price && got.Name == tt.want.Name) {
				t.Errorf("Quote.FillOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewQuotedMetric(t *testing.T) {
	type args struct {
		price  float64
		volume uint32
	}
	tests := []struct {
		name string
		args args
		want QuotedMetric
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewQuotedMetric(tt.args.price, tt.args.volume); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewQuotedMetric() = %v, want %v", got, tt.want)
			}
		})
	}
}
