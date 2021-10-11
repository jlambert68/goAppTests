create function public.sp_listtestdomains() returns TABLE(id integer, guid character varying, name character varying, description character varying, ready_for_use boolean, activated boolean, deleted boolean, update_timestamp timestamp with time zone, domain_id integer, domain_version integer)
	language plpgsql
as $$
DECLARE
    _currenttime timestamp;
begin

    SET TIMEZONE='CET';

    return query
    SELECT td.id, td.guid, td.name, td.description, td.ready_for_use, td.activated, td.deleted, td.update_timestamp, td.domain_id, td.domain_version
    FROM testdomains td
    WHERE td.deleted = false AND
          td.replaced_by_new_version = false
    ORDER BY td.name;


end;
$$;

alter function public.sp_listtestdomains() owner to testuser;

