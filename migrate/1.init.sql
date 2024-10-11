-- up

create table carts
(
    id serial,


    primary key (id)
);

create table items
(
    id       serial,
    cart_id  integer references carts not null,
    product  text                     not null,
    quantity integer                  not null,

    primary key (id)
);

-- down

drop table items;
drop table carts;
