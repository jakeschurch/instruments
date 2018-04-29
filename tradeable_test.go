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
func mockSellTx() Transaction {
	return Transaction{
		"Google",
		false,
		quotedMetric{NewPrice(15.00), NewVolume(10.00)},
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

func TestHolding_SellOff(t *testing.T) {
	wantedHolding := mockHolding()
	wantedHolding.Volume = NewVolume(10.00)

	type args struct {
		tx Transaction
	}
	tests := []struct {
		name    string
		h       *Holding
		args    args
		want    *Holding
		wantErr bool
	}{
		{"base case", mockHolding(), args{mockSellTx()}, wantedHolding, false},
		{"wrong Tx type", mockHolding(), args{mockTx(true)}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.h.SellOff(tt.args.tx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Holding.SellOff() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Holding.SellOff() = %v, want %v", got, tt.want)
			}
		})
	}
}

func mockSummary() *Summary {
	newPrice := NewPrice(10.00)
	metric := &SummaryMetric{newPrice, time.Time{}}
	return &Summary{
		0, NewVolume(10.00), &newPrice, &newPrice,
		metric, metric, metric, metric,
	}
}
func TestSummary_UpdateMetrics(t *testing.T) {
	type args struct {
		qBid Price
		qAsk Price
		t    time.Time
	}
	tests := []struct {
		name string
		s    *Summary
		args args
	}{
		{"Base case", mockSummary(), args{NewPrice(10.00), NewPrice(10.00), time.Time{}}},
		{"zero case", mockSummary(), args{NewPrice(0), NewPrice(10.00), time.Time{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.UpdateMetrics(tt.args.qBid, tt.args.qAsk, tt.args.t)
		})
	}
}

func mockSummaryMetric() *SummaryMetric {
	return &SummaryMetric{Price: NewPrice(10.00), Date: time.Time{}}
}
func TestSummaryMetric_Max(t *testing.T) {
	summMetric := mockSummaryMetric()

	type args struct {
		quotePrice Price
		timestamp  time.Time
	}
	tests := []struct {
		name string
		s    *SummaryMetric
		args args
		want Price
	}{
		{"no new max", summMetric, args{summMetric.Price - NewPrice(10.00), time.Time{}}, summMetric.Price},
		{"new max", summMetric, args{summMetric.Price + NewPrice(5.00), time.Time{}}, summMetric.Price + NewPrice(5.00)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Max(tt.args.quotePrice, tt.args.timestamp); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SummaryMetric.Max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSummaryMetric_Min(t *testing.T) {
	summMetric := mockSummaryMetric()

	type args struct {
		quotePrice Price
		timestamp  time.Time
	}
	tests := []struct {
		name string
		s    *SummaryMetric
		args args
		want Price
	}{
		{"no new min", summMetric, args{summMetric.Price + NewPrice(2.00), time.Time{}}, summMetric.Price},
		{"new min", summMetric, args{summMetric.Price - NewPrice(5.00), time.Time{}}, summMetric.Price - NewPrice(5.00)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Min(tt.args.quotePrice, tt.args.timestamp); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SummaryMetric.Min() = %v, want %v", got, tt.want)
			}
		})
	}
}
