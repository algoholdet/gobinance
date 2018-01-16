package binance

import (
	"net/url"
)

// query is used to query various API endpoints.
type query struct {
	fromID    *int64
	startTime *Time
	endTime   *Time
	limit     *int
}

func FromID(id int64) func(*query) {
	return func(q *query) {
		q.fromID = &id
	}
}

func StartTime(start Time) func(*query) {
	return func(q *query) {
		q.startTime = &start
	}
}

func EndTime(end Time) func(*query) {
	return func(q *query) {
		q.endTime = &end
	}
}

func Limit(limit int) func(*query) {
	return func(q *query) {
		q.limit = &limit
	}
}

func (q *query) params() func(url.Values) {
	return func(v url.Values) {
		if q.fromID != nil {
			param("fromId", *q.fromID)(v)
		}

		if q.startTime != nil {
			param("startTime", q.startTime.UnixNano()/1000000)(v)
		}

		if q.endTime != nil {
			param("endTime", q.endTime.UnixNano()/1000000)(v)
		}

		if q.limit != nil {
			param("limit", *q.limit)(v)
		}
	}
}
