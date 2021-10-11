create sequence public.supported_metadata_tables_table_id_seq
	as integer;

alter sequence public.supported_metadata_tables_table_id_seq owner to testuser;

alter sequence public.supported_metadata_tables_table_id_seq owned by public.supported_metadata_tables."TableId";

