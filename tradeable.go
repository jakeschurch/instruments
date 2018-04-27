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

	"github.com/pkg/errors"
)

var ErrInvalidTx = errors.New("invalid transaction type given")

type Holding struct {
	Name   string
	Volume Volume
	Buy    txMetric
	Sell   txMetric
}

// Buy creates a new Holding from transaction data.
func Buy(tx Transaction) (*Holding, error) {
	if !tx.Buy {
		return nil, errors.Wrap(ErrInvalidTx, "wanted buy, got sell")
	}
	return &Holding{
		Name:   tx.Name,
		Volume: tx.Volume,
		Buy:    txMetric{Price: tx.Price, Date: tx.Timestamp},
	}, nil
}

// SellOff a number of securities from transaction data.
func (h *Holding) SellOff(tx Transaction) (*Holding, error) {
	if tx.Buy || h.Volume < tx.Volume {
		return nil, errors.Wrap(ErrInvalidTx, "wanted sell, got buy")
	}
	h.Volume -= tx.Volume
	return h, nil
}

type txMetric struct {
	Price Price
	Date  time.Time
}

// ----------------------------------------------------------------------------

type DatedMetric struct {
}

// ----------------------------------------------------------------------------

type Summary struct {
	n              uint
	Volume         Volume
	AvgBid, AvgAsk *Price

	MaxBid, MaxAsk *summaryMetric
	MinBid, MinAsk *summaryMetric
}

func (s *Summary) UpdateMetrics(qBid, qAsk Price, t time.Time) {
	s.AvgBid.Avg(s.n, qBid)
	s.AvgAsk.Avg(s.n, qAsk)
	s.n++

	s.MaxAsk.Max(qAsk, t)
	s.MaxBid.Max(qBid, t)
	s.MinBid.Min(qBid, t)
	s.MinAsk.Min(qAsk, t)
}

// ----------------------------------------------------------------------------

type summaryMetric struct {
	Price Price
	Date  time.Time
}

func (p *Price) Avg(n uint, quotePrice Price) Price {
	numerator := *p*Price(n) + quotePrice
	newAvg := numerator / (Price(n) + 1)
	p = &newAvg

	return newAvg
}

func (s *summaryMetric) Max(quotePrice Price, timestamp time.Time) Price {
	if s.Price <= quotePrice {
		s.Price = quotePrice
	}
	return s.Price
}

func (s *summaryMetric) Min(quotePrice Price, timestamp time.Time) Price {
	if s.Price >= quotePrice {
		s.Price = quotePrice
	}
	return s.Price
}
