CREATE TABLE payments
(
    id                  bigserial primary key,
    payment_source_code varchar(10)              not null,
    payment_method_code varchar(10)              not null,
    status              varchar(100)             not null,
    fraud_status        varchar(100)             not null,
    expired_at          timestamp with time zone not null,
    order_id            varchar(100)             not null,
    status_message      varchar(255)             not null,
    payment_type        varchar(100)             not null,
    gross_amount        float                    not null,
    signature_key       text                     not null,
    actions             jsonb                    not null,
    created_at          timestamp with time zone not null,
    updated_at          timestamp with time zone not null
)