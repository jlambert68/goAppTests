create table testdomains
(
    id                      serial
        constraint testdomains_pk
            primary key,
    guid                    varchar                  not null,
    name                    varchar                  not null,
    description             varchar,
    ready_for_use           boolean default false    not null,
    activated               boolean default false    not null,
    deleted                 boolean default false    not null,
    update_timestamp        timestamp with time zone not null,
    replaced_by_new_version boolean default false    not null,
    domain_id               integer                  not null,
    domain_version          integer                  not null
);

alter table testdomains
    owner to caxdbuser;

create unique index testdomains_id_uindex
    on testdomains (id);

INSERT INTO public.testdomains (id, guid, name, description, ready_for_use, activated, deleted, update_timestamp, replaced_by_new_version, domain_id, domain_version) VALUES (18, 'ab34afa0-6456-4163-81f5-f85c2fbea3c8', 'Custody Cash2bb', 'Systems connected to Custody Cash2bb', false, false, false, '2021-10-11 14:25:04.378844 +00:00', true, 1, 1);
INSERT INTO public.testdomains (id, guid, name, description, ready_for_use, activated, deleted, update_timestamp, replaced_by_new_version, domain_id, domain_version) VALUES (36, 'ab34afa0-6456-4163-81f5-f85c2fbea3c8', 'Custody Cash', 'Systems connected to Custody Cash', false, false, false, '2021-10-11 14:25:04.378844 +00:00', false, 1, 2);
INSERT INTO public.testdomains (id, guid, name, description, ready_for_use, activated, deleted, update_timestamp, replaced_by_new_version, domain_id, domain_version) VALUES (17, 'ab34afa0-6456-4163-81f5-f85c2fbea3c8', 'Custody Cash2b', 'Systems connected to Custody Cash2b', false, false, false, '2021-10-07 18:30:48.494231 +00:00', true, 0, 0);
INSERT INTO public.testdomains (id, guid, name, description, ready_for_use, activated, deleted, update_timestamp, replaced_by_new_version, domain_id, domain_version) VALUES (19, '7b0105e1-1285-4c90-9480-cade1acec31c', 'Custody Arrangement2', '2Custody Arrangement - Cobol2', false, false, false, '2021-10-10 07:05:36.670262 +00:00', true, 0, 0);
INSERT INTO public.testdomains (id, guid, name, description, ready_for_use, activated, deleted, update_timestamp, replaced_by_new_version, domain_id, domain_version) VALUES (20, '828de370-2795-11ec-a0bd-9f11439f340c', 'name', 'description', true, true, false, '2021-10-07 17:39:35.413661 +00:00', true, 0, 0);
INSERT INTO public.testdomains (id, guid, name, description, ready_for_use, activated, deleted, update_timestamp, replaced_by_new_version, domain_id, domain_version) VALUES (22, '828de370-2795-11ec-a0bd-9f11439f340c', 'nameb_mmm', 'descriptionb_llll', true, true, false, '2021-10-08 19:18:34.103993 +00:00', true, 0, 0);
INSERT INTO public.testdomains (id, guid, name, description, ready_for_use, activated, deleted, update_timestamp, replaced_by_new_version, domain_id, domain_version) VALUES (23, '828de370-2795-11ec-a0bd-9f11439f340c', 'nameb_mmm_a', 'descriptionb_llll_a', true, true, false, '2021-10-08 19:19:20.630565 +00:00', true, 0, 0);
INSERT INTO public.testdomains (id, guid, name, description, ready_for_use, activated, deleted, update_timestamp, replaced_by_new_version, domain_id, domain_version) VALUES (9, 'ab34afa0-6456-4163-81f5-f85c2fbea3c8', 'Custody Cash2b', 'Systems connected to Custody Cash2b', false, false, false, '2021-10-07 18:30:48.494231 +00:00', true, 0, 0);
INSERT INTO public.testdomains (id, guid, name, description, ready_for_use, activated, deleted, update_timestamp, replaced_by_new_version, domain_id, domain_version) VALUES (15, 'ab34afa0-6456-4163-81f5-f85c2fbea3c8', 'Custody Cash2b', 'Systems connected to Custody Cash2b', false, false, false, '2021-10-07 18:30:48.494231 +00:00', true, 0, 0);
INSERT INTO public.testdomains (id, guid, name, description, ready_for_use, activated, deleted, update_timestamp, replaced_by_new_version, domain_id, domain_version) VALUES (13, 'ab34afa0-6456-4163-81f5-f85c2fbea3c8', 'Custody Cash2b', 'Systems connected to Custody Cash2b', false, false, false, '2021-10-07 18:30:48.494231 +00:00', true, 0, 0);
INSERT INTO public.testdomains (id, guid, name, description, ready_for_use, activated, deleted, update_timestamp, replaced_by_new_version, domain_id, domain_version) VALUES (2, '7b0105e1-1285-4c90-9480-cade1acec31c', 'Custody Arrangement', 'Custody Arrangement - Cobol', false, false, false, '2021-09-08 18:44:42.289000 +00:00', true, 0, 0);
INSERT INTO public.testdomains (id, guid, name, description, ready_for_use, activated, deleted, update_timestamp, replaced_by_new_version, domain_id, domain_version) VALUES (1, 'ab34afa0-6456-4163-81f5-f85c2fbea3c8', 'Custody Cash', 'Systems connected to Custody Cash', false, false, false, '2021-09-08 18:41:36.339000 +00:00', true, 0, 0);
INSERT INTO public.testdomains (id, guid, name, description, ready_for_use, activated, deleted, update_timestamp, replaced_by_new_version, domain_id, domain_version) VALUES (7, 'ab34afa0-6456-4163-81f5-f85c2fbea3c8', 'Custody Cash2', 'Systems connected to Custody Cash2', false, false, false, '2021-10-05 21:12:40.901925 +00:00', true, 0, 0);
INSERT INTO public.testdomains (id, guid, name, description, ready_for_use, activated, deleted, update_timestamp, replaced_by_new_version, domain_id, domain_version) VALUES (21, '828de370-2795-11ec-a0bd-9f11439f340c', 'nameb', 'descriptionb', true, true, false, '2021-10-08 14:45:19.977438 +00:00', true, 0, 0);
INSERT INTO public.testdomains (id, guid, name, description, ready_for_use, activated, deleted, update_timestamp, replaced_by_new_version, domain_id, domain_version) VALUES (14, 'ab34afa0-6456-4163-81f5-f85c2fbea3c8', 'Custody Cash2b', 'Systems connected to Custody Cash2b', false, false, false, '2021-10-07 18:30:48.494231 +00:00', true, 0, 0);
INSERT INTO public.testdomains (id, guid, name, description, ready_for_use, activated, deleted, update_timestamp, replaced_by_new_version, domain_id, domain_version) VALUES (24, '828de370-2795-11ec-a0bd-9f11439f340c', 'nameb_mmm_ab', 'descriptionb_llll_a', true, true, false, '2021-10-10 08:39:06.527082 +00:00', true, 0, 0);
INSERT INTO public.testdomains (id, guid, name, description, ready_for_use, activated, deleted, update_timestamp, replaced_by_new_version, domain_id, domain_version) VALUES (25, '22aa1e30-2941-11ec-8a51-58e9dd9e4d31', 'testnamn', 'beskrivning', false, true, true, '2021-10-09 20:54:33.536667 +00:00', false, 0, 0);
INSERT INTO public.testdomains (id, guid, name, description, ready_for_use, activated, deleted, update_timestamp, replaced_by_new_version, domain_id, domain_version) VALUES (27, 'b4077c40-2998-11ec-ba5b-ba233a60981c', 'Min nya tidzones-test 09:07 +0200', 'test', false, false, true, '2021-10-10 08:13:10.361872 +00:00', false, 0, 0);
INSERT INTO public.testdomains (id, guid, name, description, ready_for_use, activated, deleted, update_timestamp, replaced_by_new_version, domain_id, domain_version) VALUES (28, '828de370-2795-11ec-a0bd-9f11439f340c', 'nameb_mmm_abc', 'descriptionb_llll_a', true, true, true, '2021-10-11 04:50:22.356288 +00:00', false, 0, 0);
INSERT INTO public.testdomains (id, guid, name, description, ready_for_use, activated, deleted, update_timestamp, replaced_by_new_version, domain_id, domain_version) VALUES (16, 'ab34afa0-6456-4163-81f5-f85c2fbea3c8', 'Custody Cash2b', 'Systems connected to Custody Cash2b', false, false, false, '2021-10-07 18:30:48.494231 +00:00', true, 0, 0);
INSERT INTO public.testdomains (id, guid, name, description, ready_for_use, activated, deleted, update_timestamp, replaced_by_new_version, domain_id, domain_version) VALUES (33, 'bf4b8780-2a93-11ec-a969-850b9a58e9c2', 'NNNNNNNNNNNnnnnnn', 'DDDDdddd', true, true, true, '2021-10-11 14:24:00.122134 +00:00', false, 2, 1);
INSERT INTO public.testdomains (id, guid, name, description, ready_for_use, activated, deleted, update_timestamp, replaced_by_new_version, domain_id, domain_version) VALUES (26, '7b0105e1-1285-4c90-9480-cade1acec31c', 'Custody Arrangement2 0905', '2Custody Arrangement - Cobol2', false, false, false, '2021-10-11 14:24:25.341733 +00:00', true, 1, 1);
INSERT INTO public.testdomains (id, guid, name, description, ready_for_use, activated, deleted, update_timestamp, replaced_by_new_version, domain_id, domain_version) VALUES (34, '7b0105e1-1285-4c90-9480-cade1acec31c', 'Custody Arrangement 1624', 'Custody Arrangement - Cobol', false, false, false, '2021-10-11 14:24:47.365205 +00:00', true, 1, 2);
INSERT INTO public.testdomains (id, guid, name, description, ready_for_use, activated, deleted, update_timestamp, replaced_by_new_version, domain_id, domain_version) VALUES (35, '7b0105e1-1285-4c90-9480-cade1acec31c', 'Custody Arrangement', 'Custody Arrangement - Cobol', false, true, false, '2021-10-11 14:24:47.365205 +00:00', false, 1, 3);