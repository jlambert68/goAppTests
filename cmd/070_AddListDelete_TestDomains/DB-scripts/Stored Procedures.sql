-- ********************************************************************
-- List all TestDomains
create function sp_listtestdomains()
    returns TABLE(id integer, guid character varying, name character varying, description character varying, ready_for_use boolean, activated boolean, deleted boolean, update_timestamp timestamp with time zone)
    language plpgsql
as
$$
begin
    return query
        SELECT *
        FROM testdomains
        WHERE testdomains.deleted = false AND
                testdomains.replaced_by_new_version = false
        ORDER BY testdomains.name;


end;
$$;

alter function sp_listtestdomains() owner to testuser;



-- ********************************************************************


-- ********************************************************************
-- List all Tables that the user can edit
create function sp_listtablestoedit()
    returns TABLE(id integer, guid character varying, table_name character varying)
    language plpgsql

as
$$
begin
    return query
        SELECT tabletoedit.id, tabletoedit.guid, tabletoedit.table_name
        FROM tabletoedit
        WHERE tabletoedit.valid_for_use = true
        ORDER BY tabletoedit.id;
end;
$$;

alter function sp_listtablestoedit() owner to testuser;


-- ********************************************************************


-- ********************************************************************
-- List all Metadata for a certain Table, that can be edit
create function sp_list_magictable_metadata(inGuid varchar)
    returns TABLE(
                     ColumnHeaderName character varying, ColumnDataName character varying, ColumnDataType integer,
                     Sortable boolean, FormatPresentationType integer, ShouldBeVisible boolean,
                     PresentationOrder integer, UpdateIsEditable boolean, NewIsEditable boolean)
    language plpgsql

as
$$
begin
    return query
        SELECT mtmd."ColumnHeaderName", mtmd."ColumnDataName", mtmd."ColumnDataType",
               mtmd."Sortable", mtmd."FormatPresentationType", mtmd."ShouldBeVisible",
               mtmd."PresentationOrder", mtmd."UpdateIsEditable", mtmd."NewIsEditable"
        FROM magictable_metadata mtmd, tabletoedit tte
        WHERE mtmd."TableId" = tte.id AND
                tte.guid = inGuid
        ORDER BY mtmd."PresentationOrder";
end;
$$;

alter function sp_list_magictable_metadata() owner to testuser;

-- ********************************************************************


-- ********************************************************************
-- Add a New TestDomain or Update an Existing. Either way, a new row is created.

create function sp_insert_new_or_updated_testdomain(in_guid character varying, in_name character varying, in_description character varying, in_ready_for_use boolean, in_activated boolean, in_deleted boolean, in_update_timestamp timestamp with time zone)
    returns TABLE(id integer, guid character varying, name character varying, description character varying, ready_for_use boolean, activated boolean)
    language plpgsql
as
$$
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
            in_deleted,
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
               td.activated,
               td.deleted,
               td.update_timestamp,
               td.domain_id,
               td.domain_version

        FROM testdomains td
        WHERE td.guid = in_guid AND
                td.deleted = false AND
                td.replaced_by_new_version = false
        ORDER BY td.id DESC
        LIMIT 1;


end
$$;

alter function sp_insert_new_or_updated_testdomain(varchar, varchar, varchar, boolean, boolean, boolean, timestamp with time zone) owner to testuser;




-- ********************************************************************


-- ********************************************************************
-- Delete a TestDomain by setting the deleted -flag.

create function sp_delete_testdomain(in_guid character varying)
    returns TABLE(id integer, guid character varying, name character varying, description character varying, ready_for_use boolean, activated boolean, deleted boolean, update_timestamp timestamp with time zone)
    language plpgsql
as
$$
begin

    -- The old TestDomain is set to be 'old'
    UPDATE testdomains
    SET deleted = true
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

alter function sp_delete_testdomain(varchar) owner to testuser;

