create table magictable_metadata
(
    "ColumnHeaderName"       varchar              not null,
    "ColumnDataName"         varchar              not null,
    "ColumnDataType"         integer              not null,
    "Sortable"               boolean              not null,
    "FormatPresentationType" integer              not null,
    "ShouldBeVisible"        boolean              not null,
    "TableId"                integer              not null,
    "Id"                     serial,
    "PresentationOrder"      integer              not null,
    "UpdateIsEditable"       boolean default true,
    "NewIsEditable"          boolean default true not null
);

alter table magictable_metadata
    owner to caxdbuser;

create unique index magictable_metadata_id_uindex
    on magictable_metadata ("Id");

INSERT INTO public.magictable_metadata ("ColumnHeaderName", "ColumnDataName", "ColumnDataType", "Sortable", "FormatPresentationType", "ShouldBeVisible", "TableId", "Id", "PresentationOrder", "UpdateIsEditable", "NewIsEditable") VALUES ('Network', 'Network', 0, true, 0, true, 1, 4, 5, true, true);
INSERT INTO public.magictable_metadata ("ColumnHeaderName", "ColumnDataName", "ColumnDataType", "Sortable", "FormatPresentationType", "ShouldBeVisible", "TableId", "Id", "PresentationOrder", "UpdateIsEditable", "NewIsEditable") VALUES ('Mem', 'Memory', 2, true, 0, true, 1, 3, 4, true, true);
INSERT INTO public.magictable_metadata ("ColumnHeaderName", "ColumnDataName", "ColumnDataType", "Sortable", "FormatPresentationType", "ShouldBeVisible", "TableId", "Id", "PresentationOrder", "UpdateIsEditable", "NewIsEditable") VALUES ('Name', 'Name', 0, true, 0, true, 1, 6, 2, true, true);
INSERT INTO public.magictable_metadata ("ColumnHeaderName", "ColumnDataName", "ColumnDataType", "Sortable", "FormatPresentationType", "ShouldBeVisible", "TableId", "Id", "PresentationOrder", "UpdateIsEditable", "NewIsEditable") VALUES ('Description', 'Description', 0, true, 0, true, 2, 14, 3, true, true);
INSERT INTO public.magictable_metadata ("ColumnHeaderName", "ColumnDataName", "ColumnDataType", "Sortable", "FormatPresentationType", "ShouldBeVisible", "TableId", "Id", "PresentationOrder", "UpdateIsEditable", "NewIsEditable") VALUES ('Id', 'Id', 1, true, 0, true, 2, 12, 1, false, false);
INSERT INTO public.magictable_metadata ("ColumnHeaderName", "ColumnDataName", "ColumnDataType", "Sortable", "FormatPresentationType", "ShouldBeVisible", "TableId", "Id", "PresentationOrder", "UpdateIsEditable", "NewIsEditable") VALUES ('Price', 'Price', 2, true, 0, true, 1, 5, 6, true, true);
INSERT INTO public.magictable_metadata ("ColumnHeaderName", "ColumnDataName", "ColumnDataType", "Sortable", "FormatPresentationType", "ShouldBeVisible", "TableId", "Id", "PresentationOrder", "UpdateIsEditable", "NewIsEditable") VALUES ('Deleted', 'Deleted', 3, true, 0, false, 2, 9, 6, true, true);
INSERT INTO public.magictable_metadata ("ColumnHeaderName", "ColumnDataName", "ColumnDataType", "Sortable", "FormatPresentationType", "ShouldBeVisible", "TableId", "Id", "PresentationOrder", "UpdateIsEditable", "NewIsEditable") VALUES ('ECU', 'Ecu', 2, true, 0, true, 1, 0, 7, true, true);
INSERT INTO public.magictable_metadata ("ColumnHeaderName", "ColumnDataName", "ColumnDataType", "Sortable", "FormatPresentationType", "ShouldBeVisible", "TableId", "Id", "PresentationOrder", "UpdateIsEditable", "NewIsEditable") VALUES ('Name', 'Name', 0, true, 0, true, 2, 10, 2, true, true);
INSERT INTO public.magictable_metadata ("ColumnHeaderName", "ColumnDataName", "ColumnDataType", "Sortable", "FormatPresentationType", "ShouldBeVisible", "TableId", "Id", "PresentationOrder", "UpdateIsEditable", "NewIsEditable") VALUES ('Ready for Use', 'ReadyForUse', 3, true, 0, true, 2, 7, 7, true, true);
INSERT INTO public.magictable_metadata ("ColumnHeaderName", "ColumnDataName", "ColumnDataType", "Sortable", "FormatPresentationType", "ShouldBeVisible", "TableId", "Id", "PresentationOrder", "UpdateIsEditable", "NewIsEditable") VALUES ('Domain Id', 'DomainId', 1, false, 0, true, 2, 15, 9, false, false);
INSERT INTO public.magictable_metadata ("ColumnHeaderName", "ColumnDataName", "ColumnDataType", "Sortable", "FormatPresentationType", "ShouldBeVisible", "TableId", "Id", "PresentationOrder", "UpdateIsEditable", "NewIsEditable") VALUES ('Activated', 'Activated', 3, true, 0, true, 2, 11, 5, true, true);
INSERT INTO public.magictable_metadata ("ColumnHeaderName", "ColumnDataName", "ColumnDataType", "Sortable", "FormatPresentationType", "ShouldBeVisible", "TableId", "Id", "PresentationOrder", "UpdateIsEditable", "NewIsEditable") VALUES ('Instance Type', 'InstanceType', 0, false, 0, true, 1, 2, 3, true, true);
INSERT INTO public.magictable_metadata ("ColumnHeaderName", "ColumnDataName", "ColumnDataType", "Sortable", "FormatPresentationType", "ShouldBeVisible", "TableId", "Id", "PresentationOrder", "UpdateIsEditable", "NewIsEditable") VALUES ('Unique Id', 'UniqueId', 1, true, 0, true, 1, 1, 1, false, false);
INSERT INTO public.magictable_metadata ("ColumnHeaderName", "ColumnDataName", "ColumnDataType", "Sortable", "FormatPresentationType", "ShouldBeVisible", "TableId", "Id", "PresentationOrder", "UpdateIsEditable", "NewIsEditable") VALUES ('Update TimeStamp', 'UpdateTimestamp', 0, true, 0, true, 2, 8, 8, false, false);
INSERT INTO public.magictable_metadata ("ColumnHeaderName", "ColumnDataName", "ColumnDataType", "Sortable", "FormatPresentationType", "ShouldBeVisible", "TableId", "Id", "PresentationOrder", "UpdateIsEditable", "NewIsEditable") VALUES ('Guid', 'Guid', 0, true, 0, true, 2, 13, 4, false, false);
INSERT INTO public.magictable_metadata ("ColumnHeaderName", "ColumnDataName", "ColumnDataType", "Sortable", "FormatPresentationType", "ShouldBeVisible", "TableId", "Id", "PresentationOrder", "UpdateIsEditable", "NewIsEditable") VALUES ('Domain Version', 'DomainVersion', 1, false, 0, true, 2, 16, 10, false, false);