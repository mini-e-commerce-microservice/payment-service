CREATE TABLE payment_sources
(
    id          bigserial primary key,
    is_active   boolean            not null,
    code        varchar(10) UNIQUE NOT NULL,
    description varchar(255)       NOT NULL
)