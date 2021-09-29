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
-- Lxxxx

