-- +goose Up
create table ips (
    id serial primary key,
    subnet inet not null,
    type varchar(1) not null
);

-- +goose Down
drop table ips;
