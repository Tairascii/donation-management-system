-- +goose Up
-- +goose StatementBegin
create table users (
    id serial primary key,
    full_name varchar(100) not null,
    email varchar(100) unique not null,
    phone varchar(20),
    created_at timestamp default now()
);

create table organizations (
    id serial primary key,
    name varchar(100) not null,
    description text,
    website varchar(255),
    created_at timestamp default now()
);

create table campaigns (
    id serial primary key,
    organization_id integer references organizations(id) on delete cascade,
    title varchar(150) not null,
    description text,
    goal_amount integer,
    start_date date,
    end_date date,
    created_at timestamp default now()
);

create table donations (
    id serial primary key,
    user_id integer references users(id) on delete set null,
    campaign_id integer references campaigns(id) on delete set null,
    amount integer not null,
    donated_at timestamp default now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table users;
drop table organizations;
drop table campaigns;
drop table donations;
-- +goose StatementEnd
