CREATE TABLE orders
(
    order_uid          char(19) not null unique,
    track_number       varchar(20) not null,
    entry              char(4) not null,
    delivery_id        char(19) not null,
    payment_id         char(19) not null,
    items_id           char(19) not null,
    locale             char(2) not null,
    internal_signature varchar(20),
    customer_id        varchar(20) not null,
    delivery_service   varchar(5) not null,
    shardkey           char(3) not null,
    sm_id              int not null,
    date_created       char(20) not null,
    oof_shard          char(2) not null,
	PRIMARY KEY(order_uid)
);

CREATE TABLE deliveries
(
    order_id   char(19) references orders (order_uid) on delete cascade not null,
    delivery_id char(19) references orders (delivery_id) on delete cascade not null,
    name        varchar(50) not null,
    phone       varchar(50) not null,
    zip         varchar(50) not null,
    city        varchar(50) not null,
    address     varchar(50) not null,
    region      varchar(50) not null,
    email       varchar(50) not null
);

-- CREATE TABLE payments
-- (
--     order_uid     char(19) references orders (order_uid) on delete cascade not null unique,
--     payment_id    char(19) references orders (payment_id) on delete cascade not null unique,
--     transaction   char(19) not null,
--     request_id    varchar(20),
--     currency      char(3) not null,
--     provider      varchar(5) not null,
--     amount        int not null,
--     payment_dt    int not null,
--     bank char(5)  not null,
--     delivery_cost int not null,
--     goods_total   int not null,
--     custom_fee    int not null
-- );

-- CREATE TABLE items
-- (
--     order_uid char(19) references orders (order_uid) on delete cascade not null unique,
--     chrt_id   serial not null unique,
--     items_id  char(19) references orders (items_id) on delete cascade not null
-- );

-- CREATE TABLE chrts
-- (
--     chrt_id     int references items (chrt_id) on delete cascade not null,
--     price       int not null,
--     rid         varchar(50) not null,
--     name        varchar(50) not null,
--     sale        int not null,
--     size        int not null,
--     total_price int not null,
--     nm_id       int not null,
--     brand       varchar(50) not null,
--     status      int not null
-- );
