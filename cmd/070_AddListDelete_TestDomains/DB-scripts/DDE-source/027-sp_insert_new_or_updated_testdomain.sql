create function public.sp_insert_new_or_updated_testdomain(in_guid character varying, in_name character varying, in_description character varying, in_ready_for_use boolean, in_activated boolean) returns TABLE(id integer, guid character varying, name character varying, description character varying, ready_for_use boolean, activated boolean)
	language plpgsql
as $$
DECLARE
    _currenttime timestamp;
    _domain_id integer;
    _domain_version integer;

begin

    SET TIMEZONE='CET';
    
    -- Get current timestamp
    _currenttime = CURRENT_TIMESTAMP;

    -- Get domain_id & domain_version from current DB-data
    -- If exists then add 1 to version
    -- Otherwise set 1 to both id and version
    IF EXISTS(SELECT domain_id
              FROM testdomains td
              WHERE td.guid = in_guid AND
                    td.replaced_by_new_version = false AND
                    td.deleted = false)
    THEN
        -- Existing, just add 1 to version number
        SELECT domain_id, domain_version into _domain_id, _domain_version
        FROM testdomains td
        WHERE td.guid = in_guid AND
        td.replaced_by_new_version = false AND
        td.deleted = false;

        _domain_version = _domain_version + 1;
    ELSE
        -- New, find highest domain_id and add 1 and set it to be version 1
        SELECT domain_id into _domain_id
        FROM testdomains td
        WHERE td.replaced_by_new_version = false AND
            td.deleted = false
        ORDER BY td.id DESC
        LIMIT 1;

        _domain_id = _domain_id + 1;
        _domain_version = 1;
    END IF;

    -- The old TestDomain is set to be 'old' 
    UPDATE testdomains
    SET replaced_by_new_version = true,
        update_timestamp =_currenttime
    WHERE testdomains.guid = in_guid AND
          testdomains.replaced_by_new_version = false AND
          testdomains.deleted = false;
--commit;
    -- Insert New or Updated TestDomain as new row
    insert into testdomains(guid,
                            name,
                            description,
                            ready_for_use,
                            activated,
                            deleted,
                            update_timestamp,
                            replaced_by_new_version,
                            domain_id,
                            domain_version)
    values (in_guid,
            in_name,
            in_description,
            in_ready_for_use,
            in_activated,
            false,
            _currenttime,
            false,
            _domain_id,
            _domain_version);
    --commit;

    -- Retrieve the newly created or updated TestDomain
    return query
    SELECT td.id,
           td.guid,
           td.name,
           td.description,
           td.ready_for_use,
           td.activated
           --td.deleted
           --td.update_timestamp,
           --td.domain_id,
           --.domain_version

    FROM testdomains td
    WHERE td.guid = in_guid AND
            td.deleted = false AND
            td.replaced_by_new_version = false
    ORDER BY td.id DESC
    LIMIT 1;


end
$$;

alter function public.sp_insert_new_or_updated_testdomain(varchar, varchar, varchar, boolean, boolean) owner to testuser;

