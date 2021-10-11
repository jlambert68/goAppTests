create table public.magictable_metadata
(
	"ColumnHeaderName" varchar not null,
	"ColumnDataName" varchar not null,
	"ColumnDataType" integer not null,
	"Sortable" boolean not null,
	"FormatPresentationType" integer not null,
	"ShouldBeVisible" boolean not null,
	"TableId" integer not null,
	"Id" serial,
	"PresentationOrder" integer not null,
	"UpdateIsEditable" boolean default true,
	"NewIsEditable" boolean default true not null
);

alter table public.magictable_metadata owner to testuser;

create unique index magictable_metadata_id_uindex
	on public.magictable_metadata ("Id");

