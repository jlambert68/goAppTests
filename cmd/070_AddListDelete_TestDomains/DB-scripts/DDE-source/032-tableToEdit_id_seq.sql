create sequence public."tableToEdit_id_seq"
	as integer;

alter sequence public."tableToEdit_id_seq" owner to testuser;

alter sequence public."tableToEdit_id_seq" owned by public.tabletoedit.id;

