-- +goose Up
create table ips (
    id serial primary key,
    ip varchar(255) not null,
    type varchar(1) not null
);

-- +goose Down
drop table ips;
