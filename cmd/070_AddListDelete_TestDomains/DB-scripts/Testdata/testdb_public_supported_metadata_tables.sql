create table supported_metadata_tables
(
    "TableId"   integer default nextval('supported_metadata_tables_table_id_seq'::regclass) not null
        constraint supported_metadata_tables_pk
            primary key,
    "TableName" varchar                                                                     not null
);

alter table supported_metadata_tables
    owner to testuser;

create unique index supported_metadata_tables_table_id_uindex
    on supported_metadata_tables ("TableId");

create unique index supported_metadata_tables_tablename_uindex
    on supported_metadata_tables ("TableName");

INSERT INTO public.supported_metadata_tables ("TableId", "TableName") VALUES (0, 'Test Domains');
INSERT INTO public.supported_metadata_tables ("TableId", "TableName") VALUES (1, 'Test Plugins');
INSERT INTO public.supported_metadata_tables ("TableId", "TableName") VALUES (2, 'Original json file');