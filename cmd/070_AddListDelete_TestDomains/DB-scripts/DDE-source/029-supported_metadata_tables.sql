create table public.supported_metadata_tables
(
	"TableId" integer default nextval('supported_metadata_tables_table_id_seq'::regclass) not null
		constraint supported_metadata_tables_pk
			primary key,
	"TableName" varchar not null
);

alter table public.supported_metadata_tables owner to testuser;

create unique index supported_metadata_tables_table_id_uindex
	on public.supported_metadata_tables ("TableId");

create unique index supported_metadata_tables_tablename_uindex
	on public.supported_metadata_tables ("TableName");

