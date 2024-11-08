package models

type OutboxEvent struct {
	ID            string `db:"id"`
	AggregateType string `db:"aggregatetype"`
	AggregateID   string `db:"aggregateid"`
	Type          string `db:"type"`
	Payload       any    `db:"payload"`
	TraceParent   string `db:"trace_parent"`
}

type PaymentPayloadOutboxEvent struct {
	PaymentData Payment `json:"payment_data"`
	ErrorReason *string `json:"error_reason,omitempty"`
}
