create function public.sp_list_magictable_metadata(inguid character varying) returns TABLE(columnheadername character varying, columndataname character varying, columndatatype integer, sortable boolean, formatpresentationtype integer, shouldbevisible boolean, presentationorder integer, updateiseditable boolean, newiseditable boolean)
	language plpgsql
as $$
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

alter function public.sp_list_magictable_metadata(varchar) owner to caxdbuser;

