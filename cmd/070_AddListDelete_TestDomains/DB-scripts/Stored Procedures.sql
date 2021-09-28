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

