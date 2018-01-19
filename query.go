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

// QueryFunc is the function signature to use for setting various query
// parameters.
type QueryFunc func(*query)

// newQuery will return a new query based on options.
func newQuery(options []QueryFunc) *query {
	q := &query{}

	for _, o := range options {
		o(q)
	}

	return q
}

// FromID can be used to setting fromID to id.
func FromID(id int64) QueryFunc {
	return func(q *query) {
		q.fromID = &id
	}
}

// StartTime will set a start time for the query. The time is inclusive.
func StartTime(start Time) QueryFunc {
	return func(q *query) {
		q.startTime = &start
	}
}

// EndTime will set an end time for the query. The time is inclusive.
func EndTime(end Time) QueryFunc {
	return func(q *query) {
		q.endTime = &end
	}
}

// Limit can be used to only return limit objects from a query.
func Limit(limit int) QueryFunc {
	return func(q *query) {
		q.limit = &limit
	}
}

// params can be passed to URL builders.
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
