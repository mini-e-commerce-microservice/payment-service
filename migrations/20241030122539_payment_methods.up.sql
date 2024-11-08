CREATE TABLE payment_methods
(
    id               bigserial primary key,
    is_active        boolean            not null,
    payment_fee      numeric(9, 1)      not null,
    payment_fee_type varchar(25)        not null,
    code             varchar(10) UNIQUE NOT NULL,
    category         varchar(255)       NOT NULL,
    name             varchar(255)       NOT NULL,
    image            varchar(255)       NOT NULL
)