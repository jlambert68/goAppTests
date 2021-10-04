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
        WHERE testdomains.deleted = false
        ORDER BY testdomains.id;
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

create function sp_insert_new_or_updated_testdomain(
    in_guid character varying,
    in_name character varying,
    in_description character varying,
    in_ready_for_use boolean,
    in_activated boolean,
    in_deleted boolean,
    in_update_timestamp timestamptz)

-- Returns the newly created or updated TestDomain
    RETURNS TABLE
            (
                id               integer,
                guid             character varying,
                name             character varying,
                description      character varying,
                ready_for_use    boolean,
                activated        boolean,
                deleted          boolean,
                update_timestamp timestamptz
            )


    language plpgsql
as
$$


begin
    -- Insert New or Updated TestDomain as new row
    insert into testdomains(guid,
                            name,
                            description,
                            ready_for_use,
                            activated,
                            deleted,
                            update_timestamp)
    values (in_guid,
            in_name,
            in_description,
            in_ready_for_use,
            in_activated,
            in_deleted,
            in_update_timestamp);
    commit;

    -- Retrieve the newly created or updated TestDomain
    return query
        SELECT  (td.id,
                 td.guid,
                 td.name,
                 td.description,
                 td.ready_for_use,
                 td.activated,
                 td.deleted,
                 td.update_timestamp)

        FROM testdomains td
        WHERE td.guid = in_guid
        ORDER BY td.id DESC
        LIMIT 1;


end;
$$;

alter function sp_insert_new_or_updated_testdomain(
    varchar,
    varchar,
    varchar,
    boolean,
    boolean,
    boolean,
    timestamptz) owner to testuser;



-- ********************************************************************


-- ********************************************************************
-- Lxxxx