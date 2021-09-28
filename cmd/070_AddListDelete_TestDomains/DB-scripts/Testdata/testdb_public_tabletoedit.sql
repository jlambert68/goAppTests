create table tabletoedit
(
    id                      integer default nextval('"tableToEdit_id_seq"'::regclass) not null,
    guid                    varchar                                                   not null,
    table_name              varchar                                                   not null,
    grpc_api_identifier     integer                                                   not null,
    valid_for_use           boolean default true                                      not null,
    gpc_api_identifier_name varchar                                                   not null
);

alter table tabletoedit
    owner to testuser;

create unique index tabletoedit_guid_uindex
    on tabletoedit (guid);

create unique index tabletoedit_id_uindex
    on tabletoedit (id);

create unique index tabletoedit_table_name_uindex
    on tabletoedit (table_name);

create unique index tabletoedit_grpc_api_identifier_uindex
    on tabletoedit (grpc_api_identifier);

INSERT INTO public.tabletoedit (id, guid, table_name, grpc_api_identifier, valid_for_use, gpc_api_identifier_name) VALUES (2, '8acacaaf-676e-4b36-abe6-c5310822ade1', 'TestDomains', 1, true, 'DomainModel');
INSERT INTO public.tabletoedit (id, guid, table_name, grpc_api_identifier, valid_for_use, gpc_api_identifier_name) VALUES (3, '81c5d008-a38a-4c47-936a-d6c3c258ae13', 'TestInstructions', 2, false, 'InstructionModel');
INSERT INTO public.tabletoedit (id, guid, table_name, grpc_api_identifier, valid_for_use, gpc_api_identifier_name) VALUES (1, '51253aba-41a9-42ef-b5f1-d8d1d7116b47', 'Orginal MagicTable', 0, true, 'TestModel');