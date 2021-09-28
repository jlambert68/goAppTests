create table testdomains
(
    id               serial
        constraint testdomains_pk
            primary key,
    guid             varchar                  not null,
    name             varchar                  not null,
    description      varchar,
    ready_for_use    boolean default false    not null,
    activated        boolean default false    not null,
    deleted          boolean default false    not null,
    update_timestamp timestamp with time zone not null
);

alter table testdomains
    owner to testuser;

create unique index testdomains_guid_uindex
    on testdomains (guid);

create unique index testdomains_id_uindex
    on testdomains (id);

create unique index testdomains_name_uindex
    on testdomains (name);

INSERT INTO public.testdomains (id, guid, name, description, ready_for_use, activated, deleted, update_timestamp) VALUES (1, 'ab34afa0-6456-4163-81f5-f85c2fbea3c8', 'Custody Cash', 'Systems connected to Custody Cash', false, false, false, '2021-09-08 18:41:36.339000 +00:00');
INSERT INTO public.testdomains (id, guid, name, description, ready_for_use, activated, deleted, update_timestamp) VALUES (2, '7b0105e1-1285-4c90-9480-cade1acec31c', 'Custody Arrangement', 'Custody Arrangement - Cobol', false, false, false, '2021-09-08 18:44:42.289000 +00:00');