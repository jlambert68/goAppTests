create table public.tabletoedit
(
	id integer default nextval('"tableToEdit_id_seq"'::regclass) not null,
	guid varchar not null,
	table_name varchar not null,
	grpc_api_identifier integer not null,
	valid_for_use boolean default true not null,
	gpc_api_identifier_name varchar not null
);

alter table public.tabletoedit owner to testuser;

create unique index tabletoedit_guid_uindex
	on public.tabletoedit (guid);

create unique index tabletoedit_id_uindex
	on public.tabletoedit (id);

create unique index tabletoedit_table_name_uindex
	on public.tabletoedit (table_name);

create unique index tabletoedit_grpc_api_identifier_uindex
	on public.tabletoedit (grpc_api_identifier);

