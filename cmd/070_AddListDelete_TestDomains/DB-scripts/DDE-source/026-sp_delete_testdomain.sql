create function public.sp_delete_testdomain(in_guid character varying) returns TABLE(id integer, guid character varying, name character varying, description character varying, ready_for_use boolean, activated boolean, deleted boolean, update_timestamp timestamp with time zone)
	language plpgsql
as $$
DECLARE
    _currenttime timestamp;
begin

    SET TIMEZONE='CET';

    -- Get current timestamp
    _currenttime = CURRENT_TIMESTAMP;

    -- The old TestDomain is set to be 'old'
    UPDATE testdomains
    SET deleted = true,
        update_timestamp = _currenttime
    WHERE  testdomains.guid = in_guid AND
          testdomains.deleted = false AND
          testdomains.replaced_by_new_version = false;


    -- Retrieve the newly deletedTestDomain
    return query
        SELECT  td.id,
                td.guid,
                td.name,
                td.description,
                td.ready_for_use,
                td.activated,
                td.deleted,
                td.update_timestamp

        FROM testdomains td
        WHERE td.guid = in_guid AND
              td.deleted = true AND
              td.replaced_by_new_version = false
        ORDER BY td.id DESC
        LIMIT 1;


end
$$;

alter function public.sp_delete_testdomain(varchar) owner to testuser;

