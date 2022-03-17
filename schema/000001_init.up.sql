CREATE TABLE orders
(
    order_uid          varchar(50) not null unique,
    track_number       varchar(50) not null,
    entry              char(4)     not null,
    locale             char(2)     not null,
    internal_signature varchar(50),
    customer_id        varchar(50) not null,
    delivery_service   varchar(50) not null,
    shardkey           char(3)     not null,
    sm_id              int         not null,
    date_created       timestamp   not null,
    oof_shard          char(2)     not null,
	PRIMARY KEY        (order_uid)
);

CREATE TABLE deliveries
(
    delivery_id serial      not null unique,
    order_uid   varchar(50) not null unique,
    name        varchar(50) not null,
    phone       varchar(50) not null,
    zip         varchar(50) not null,
    city        varchar(50) not null,
    address     varchar(50) not null,
    region      varchar(50) not null,
    email       varchar(50) not null,
    PRIMARY KEY (delivery_id),

    FOREIGN KEY (order_uid) REFERENCES orders (order_uid) ON DELETE CASCADE
);

CREATE TABLE payments
(
    payment_id    serial      not null unique,
    transaction   varchar(50) not null unique,
    request_id    varchar(50),
    currency      varchar(50) not null,
    provider      varchar(50) not null,
    amount        int         not null,
    payment_dt    int         not null,
    bank          varchar(50) not null,
    delivery_cost int         not null,
    goods_total   int         not null,
    custom_fee    int         not null,
    PRIMARY KEY   (transaction),

    FOREIGN KEY (transaction) REFERENCES orders (order_uid) ON DELETE CASCADE
);

CREATE TABLE items
(
    chrt_id      int         not null unique,
    order_uid    varchar(50) not null,
    track_number varchar(50) not null,
    price        int         not null,
    rid          varchar(50) not null,
    name         varchar(50) not null,
    sale         int         not null,
    size         int         not null,
    total_price  int         not null,
    nm_id        int         not null,
    brand        varchar(50) not null,
    status       int         not null,
    PRIMARY KEY  (chrt_id),

    FOREIGN KEY (order_uid) REFERENCES orders (order_uid) ON DELETE CASCADE
);
