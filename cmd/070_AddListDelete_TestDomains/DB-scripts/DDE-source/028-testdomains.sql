create table public.testdomains
(
	id serial
		constraint testdomains_pk
			primary key,
	guid varchar not null,
	name varchar not null,
	description varchar,
	ready_for_use boolean default false not null,
	activated boolean default false not null,
	deleted boolean default false not null,
	update_timestamp timestamp with time zone not null,
	replaced_by_new_version boolean default false not null,
	domain_id integer not null,
	domain_version integer not null
);

alter table public.testdomains owner to testuser;

create unique index testdomains_id_uindex
	on public.testdomains (id);

