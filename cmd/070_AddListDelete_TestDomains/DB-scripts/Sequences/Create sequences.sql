create sequence testdomains_id_seq
    as integer;

alter sequence testdomains_id_seq owner to testuser;

alter sequence testdomains_id_seq owned by testdomains.id;

create sequence testinstructions_id_seq
    as integer;

alter sequence testinstructions_id_seq owner to testuser;

alter sequence testinstructions_id_seq owned by testinstructions.id;

create sequence supported_metadata_tables_table_id_seq
    as integer;

alter sequence supported_metadata_tables_table_id_seq owner to testuser;

alter sequence supported_metadata_tables_table_id_seq owned by supported_metadata_tables."TableId";

create sequence "tableToEdit_id_seq"
    as integer;

alter sequence "tableToEdit_id_seq" owner to testuser;

alter sequence "tableToEdit_id_seq" owned by tabletoedit.id;

create sequence "magictable_metadata_Id_seq"
    as integer;

alter sequence "magictable_metadata_Id_seq" owner to testuser;

alter sequence "magictable_metadata_Id_seq" owned by magictable_metadata."Id";

