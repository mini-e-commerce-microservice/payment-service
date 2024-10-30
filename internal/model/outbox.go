package model

type OutboxEvent struct {
	ID            string `db:"id"`
	AggregateType string `db:"aggregatetype"`
	AggregateID   string `db:"aggregateid"`
	Type          string `db:"type"`
	Payload       any    `db:"payload"`
	TraceParent   string `db:"trace_parent"`
}
