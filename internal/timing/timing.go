package timing

import "time"

// orderTicker is an internal struct that allows
// for simulation of order/transaction datetimes.
type OrderTicker struct {
	start  time.Time
	ticker *time.Ticker
}

func newOrderTicker() *orderTicker {
	return &orderTicker{
		start:  time.Now(),
		ticker: time.NewTicker(time.Millisecond),
	}
}
func (oT *orderTicker) duration() time.Duration {
	t := <-oT.ticker.C
	return t.Sub(oT.start)
}
