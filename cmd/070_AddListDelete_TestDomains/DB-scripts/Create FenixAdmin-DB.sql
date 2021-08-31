-- ********************************************************************************
-- Test Instructions
create table testinstructions
(
    id serial,
    guid varchar not null,
    name varchar not null,
    description varchar,
    ready_for_use bool default false not null,
    activated bool default false not null,
    deleted bool default false not null,
    update_timestamp timestamptz not null
);

create unique index testinstructions_guid_uindex
    on testinstructions (guid);

create unique index testinstructions_id_uindex
    on testinstructions (id);

create unique index testinstructions_name_uindex
    on testinstructions (name);

alter table testinstructions
    add constraint testinstructions_pk
        primary key (id);

-- ********************************************************************************
-- Test Domains
create table testdomains
(
    id serial,
    guid varchar not null,
    name varchar not null,
    description varchar,
    ready_for_use bool default false not null,
    activated bool default false not null,
    deleted bool default false not null,
    update_timestamp timestamptz not null
);

create unique index testdomains_guid_uindex
    on testdomains (guid);

create unique index testdomains_id_uindex
    on testdomains (id);

create unique index testdomains_name_uindex
    on testdomains (name);

alter table testdomains
    add constraint testdomains_pk
        primary key (id);