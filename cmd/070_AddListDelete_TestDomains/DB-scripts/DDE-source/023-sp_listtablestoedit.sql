create function public.sp_listtablestoedit() returns TABLE(id integer, guid character varying, table_name character varying)
	language plpgsql
as $$
begin
    return query
        SELECT tabletoedit.id, tabletoedit.guid, tabletoedit.table_name
        FROM tabletoedit
        WHERE tabletoedit.valid_for_use = true
        ORDER BY tabletoedit.id;
end;
$$;

alter function public.sp_listtablestoedit() owner to caxdbuser;

