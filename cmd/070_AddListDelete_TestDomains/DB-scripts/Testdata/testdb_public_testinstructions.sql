create table testinstructions
(
    id               serial
        constraint testinstructions_pk
            primary key,
    guid             varchar                  not null,
    name             varchar                  not null,
    description      varchar,
    ready_for_use    boolean default false    not null,
    activated        boolean default false    not null,
    deleted          boolean default false    not null,
    update_timestamp timestamp with time zone not null
);

alter table testinstructions
    owner to caxdbuser;

create unique index testinstructions_guid_uindex
    on testinstructions (guid);

create unique index testinstructions_id_uindex
    on testinstructions (id);

create unique index testinstructions_name_uindex
    on testinstructions (name);

