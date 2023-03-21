package domain

// Represents an object used to query
type BaseQueryData struct {
	Limit int64
	Skip  int64
}

// Initialise a query data with Limit and Skip
func NewBaseQueryData() (q BaseQueryData) {
	q.Limit = int64(50)
	q.Skip = int64(0)
	return
}
