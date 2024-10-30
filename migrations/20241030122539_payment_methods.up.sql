CREATE TABLE payment_methods
(
    id        bigserial primary key,
    is_active boolean            not null,
    code      varchar(10) UNIQUE NOT NULL,
    category  varchar(255)       NOT NULL,
    name      varchar(255)       NOT NULL,
    image     varchar(255)       NOT NULL
)