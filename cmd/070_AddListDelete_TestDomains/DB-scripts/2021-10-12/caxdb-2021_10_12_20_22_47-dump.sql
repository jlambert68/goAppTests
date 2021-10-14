--
-- PostgreSQL database dump
--

-- Dumped from database version 14.0 (Ubuntu 14.0-1.pgdg18.04+1)
-- Dumped by pg_dump version 14.0 (Ubuntu 14.0-1.pgdg18.04+1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

DROP DATABASE IF EXISTS caxdb;
--
-- Name: caxdb; Type: DATABASE; Schema: -; Owner: postgres
--

CREATE DATABASE caxdb WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE = 'en_US.UTF-8';


ALTER DATABASE caxdb OWNER TO postgres;

\connect caxdb

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: uuid-ossp; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;


--
-- Name: EXTENSION "uuid-ossp"; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION "uuid-ossp" IS 'generate universally unique identifiers (UUIDs)';


--
-- Name: sp_delete_testdomain(character varying); Type: FUNCTION; Schema: public; Owner: caxdbuser
--

CREATE FUNCTION public.sp_delete_testdomain(in_guid character varying) RETURNS TABLE(id integer, guid uuid, name character varying, description character varying, ready_for_use boolean, activated boolean, deleted boolean, update_timestamp timestamp with time zone)
    LANGUAGE plpgsql
    AS $$
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


ALTER FUNCTION public.sp_delete_testdomain(in_guid character varying) OWNER TO caxdbuser;

--
-- Name: sp_list_magictable_metadata(uuid); Type: FUNCTION; Schema: public; Owner: caxdbuser
--

CREATE FUNCTION public.sp_list_magictable_metadata(in_guid uuid) RETURNS TABLE(columnheadername character varying, columndataname character varying, columndatatype integer, sortable boolean, formatpresentationtype integer, shouldbevisible boolean, presentationorder integer, updateiseditable boolean, newiseditable boolean)
    LANGUAGE plpgsql
    AS $$
begin
    return query
        SELECT mtmd."ColumnHeaderName", mtmd."ColumnDataName", mtmd."ColumnDataType",
               mtmd."Sortable", mtmd."FormatPresentationType", mtmd."ShouldBeVisible",
               mtmd."PresentationOrder", mtmd."UpdateIsEditable", mtmd."NewIsEditable"
        FROM magictable_metadata mtmd, tabletoedit tte
        WHERE mtmd."TableId" = tte.id AND
              tte.guid = in_guid
        ORDER BY mtmd."PresentationOrder";
end;
$$;


ALTER FUNCTION public.sp_list_magictable_metadata(in_guid uuid) OWNER TO caxdbuser;

--
-- Name: sp_listtablestoedit(); Type: FUNCTION; Schema: public; Owner: caxdbuser
--

CREATE FUNCTION public.sp_listtablestoedit() RETURNS TABLE(id integer, guid uuid, table_name character varying)
    LANGUAGE plpgsql
    AS $$
begin
    return query
        SELECT tabletoedit.id, tabletoedit.guid, tabletoedit.table_name
        FROM tabletoedit
        WHERE tabletoedit.valid_for_use = true
        ORDER BY tabletoedit.id;
end;
$$;


ALTER FUNCTION public.sp_listtablestoedit() OWNER TO caxdbuser;

--
-- Name: sp_listtestdomains(); Type: FUNCTION; Schema: public; Owner: caxdbuser
--

CREATE FUNCTION public.sp_listtestdomains() RETURNS TABLE(id integer, guid uuid, name character varying, description character varying, ready_for_use boolean, activated boolean, deleted boolean, update_timestamp timestamp with time zone, domain_id integer, domain_version integer)
    LANGUAGE plpgsql
    AS $$
DECLARE
    _currenttime timestamp;
begin

    SET TIMEZONE='CET';

    return query
    SELECT td.id, td.guid, td.name, td.description, td.ready_for_use, td.activated, td.deleted, td.update_timestamp, td.domain_id, td.domain_version
    FROM testdomains td
    WHERE td.deleted = false AND
          td.replaced_by_new_version = false
    ORDER BY td.name;


end;
$$;


ALTER FUNCTION public.sp_listtestdomains() OWNER TO caxdbuser;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: magictable_metadata; Type: TABLE; Schema: public; Owner: caxdbuser
--

CREATE TABLE public.magictable_metadata (
    "ColumnHeaderName" character varying NOT NULL,
    "ColumnDataName" character varying NOT NULL,
    "ColumnDataType" integer NOT NULL,
    "Sortable" boolean NOT NULL,
    "FormatPresentationType" integer NOT NULL,
    "ShouldBeVisible" boolean NOT NULL,
    "TableId" integer NOT NULL,
    "Id" integer NOT NULL,
    "PresentationOrder" integer NOT NULL,
    "UpdateIsEditable" boolean DEFAULT true,
    "NewIsEditable" boolean DEFAULT true NOT NULL
);


ALTER TABLE public.magictable_metadata OWNER TO caxdbuser;

--
-- Name: magictable_metadata_Id_seq; Type: SEQUENCE; Schema: public; Owner: caxdbuser
--

CREATE SEQUENCE public."magictable_metadata_Id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."magictable_metadata_Id_seq" OWNER TO caxdbuser;

--
-- Name: magictable_metadata_Id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: caxdbuser
--

ALTER SEQUENCE public."magictable_metadata_Id_seq" OWNED BY public.magictable_metadata."Id";


--
-- Name: supported_metadata_tables; Type: TABLE; Schema: public; Owner: caxdbuser
--

CREATE TABLE public.supported_metadata_tables (
    "TableId" integer NOT NULL,
    "TableName" character varying NOT NULL
);


ALTER TABLE public.supported_metadata_tables OWNER TO caxdbuser;

--
-- Name: tabletoedit; Type: TABLE; Schema: public; Owner: caxdbuser
--

CREATE TABLE public.tabletoedit (
    id integer NOT NULL,
    guid uuid NOT NULL,
    table_name character varying NOT NULL,
    grpc_api_identifier integer NOT NULL,
    valid_for_use boolean DEFAULT true NOT NULL,
    gpc_api_identifier_name character varying NOT NULL
);


ALTER TABLE public.tabletoedit OWNER TO caxdbuser;

--
-- Name: tabletoedit_id_seq; Type: SEQUENCE; Schema: public; Owner: caxdbuser
--

CREATE SEQUENCE public.tabletoedit_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.tabletoedit_id_seq OWNER TO caxdbuser;

--
-- Name: tabletoedit_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: caxdbuser
--

ALTER SEQUENCE public.tabletoedit_id_seq OWNED BY public.tabletoedit.id;


--
-- Name: testdomains; Type: TABLE; Schema: public; Owner: caxdbuser
--

CREATE TABLE public.testdomains (
    id integer NOT NULL,
    guid uuid NOT NULL,
    name character varying NOT NULL,
    description character varying,
    ready_for_use boolean DEFAULT false NOT NULL,
    activated boolean DEFAULT false NOT NULL,
    deleted boolean DEFAULT false NOT NULL,
    update_timestamp timestamp with time zone NOT NULL,
    replaced_by_new_version boolean DEFAULT false NOT NULL,
    domain_id integer NOT NULL,
    domain_version integer NOT NULL
);


ALTER TABLE public.testdomains OWNER TO caxdbuser;

--
-- Name: testdomains_id_seq; Type: SEQUENCE; Schema: public; Owner: caxdbuser
--

CREATE SEQUENCE public.testdomains_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.testdomains_id_seq OWNER TO caxdbuser;

--
-- Name: testdomains_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: caxdbuser
--

ALTER SEQUENCE public.testdomains_id_seq OWNED BY public.testdomains.id;


--
-- Name: testinstructions; Type: TABLE; Schema: public; Owner: caxdbuser
--

CREATE TABLE public.testinstructions (
    id integer NOT NULL,
    guid uuid NOT NULL,
    name character varying NOT NULL,
    description character varying,
    ready_for_use boolean DEFAULT false NOT NULL,
    activated boolean DEFAULT false NOT NULL,
    deleted boolean DEFAULT false NOT NULL,
    update_timestamp timestamp with time zone NOT NULL
);


ALTER TABLE public.testinstructions OWNER TO caxdbuser;

--
-- Name: testinstructions_id_seq; Type: SEQUENCE; Schema: public; Owner: caxdbuser
--

CREATE SEQUENCE public.testinstructions_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.testinstructions_id_seq OWNER TO caxdbuser;

--
-- Name: testinstructions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: caxdbuser
--

ALTER SEQUENCE public.testinstructions_id_seq OWNED BY public.testinstructions.id;


--
-- Name: magictable_metadata Id; Type: DEFAULT; Schema: public; Owner: caxdbuser
--

ALTER TABLE ONLY public.magictable_metadata ALTER COLUMN "Id" SET DEFAULT nextval('public."magictable_metadata_Id_seq"'::regclass);


--
-- Name: tabletoedit id; Type: DEFAULT; Schema: public; Owner: caxdbuser
--

ALTER TABLE ONLY public.tabletoedit ALTER COLUMN id SET DEFAULT nextval('public.tabletoedit_id_seq'::regclass);


--
-- Name: testdomains id; Type: DEFAULT; Schema: public; Owner: caxdbuser
--

ALTER TABLE ONLY public.testdomains ALTER COLUMN id SET DEFAULT nextval('public.testdomains_id_seq'::regclass);


--
-- Name: testinstructions id; Type: DEFAULT; Schema: public; Owner: caxdbuser
--

ALTER TABLE ONLY public.testinstructions ALTER COLUMN id SET DEFAULT nextval('public.testinstructions_id_seq'::regclass);


--
-- Data for Name: magictable_metadata; Type: TABLE DATA; Schema: public; Owner: caxdbuser
--

COPY public.magictable_metadata ("ColumnHeaderName", "ColumnDataName", "ColumnDataType", "Sortable", "FormatPresentationType", "ShouldBeVisible", "TableId", "Id", "PresentationOrder", "UpdateIsEditable", "NewIsEditable") FROM stdin;
Network	Network	0	t	0	t	1	4	5	t	t
Mem	Memory	2	t	0	t	1	3	4	t	t
Name	Name	0	t	0	t	1	6	2	t	t
Description	Description	0	t	0	t	2	14	3	t	t
Id	Id	1	t	0	t	2	12	1	f	f
Price	Price	2	t	0	t	1	5	6	t	t
Deleted	Deleted	3	t	0	f	2	9	6	t	t
ECU	Ecu	2	t	0	t	1	0	7	t	t
Name	Name	0	t	0	t	2	10	2	t	t
Ready for Use	ReadyForUse	3	t	0	t	2	7	7	t	t
Domain Id	DomainId	1	f	0	t	2	15	9	f	f
Activated	Activated	3	t	0	t	2	11	5	t	t
Instance Type	InstanceType	0	f	0	t	1	2	3	t	t
Unique Id	UniqueId	1	t	0	t	1	1	1	f	f
Update TimeStamp	UpdateTimestamp	0	t	0	t	2	8	8	f	f
Guid	Guid	0	t	0	t	2	13	4	f	f
Domain Version	DomainVersion	1	f	0	t	2	16	10	f	f
\.


--
-- Data for Name: supported_metadata_tables; Type: TABLE DATA; Schema: public; Owner: caxdbuser
--

COPY public.supported_metadata_tables ("TableId", "TableName") FROM stdin;
0	Test Domains
1	Test Plugins
2	Original json file
\.


--
-- Data for Name: tabletoedit; Type: TABLE DATA; Schema: public; Owner: caxdbuser
--

COPY public.tabletoedit (id, guid, table_name, grpc_api_identifier, valid_for_use, gpc_api_identifier_name) FROM stdin;
2	8acacaaf-676e-4b36-abe6-c5310822ade1	TestDomains	1	t	DomainModel
3	81c5d008-a38a-4c47-936a-d6c3c258ae13	TestInstructions	2	f	InstructionModel
1	51253aba-41a9-42ef-b5f1-d8d1d7116b47	Orginal MagicTable	0	t	TestModel
\.


--
-- Data for Name: testdomains; Type: TABLE DATA; Schema: public; Owner: caxdbuser
--

COPY public.testdomains (id, guid, name, description, ready_for_use, activated, deleted, update_timestamp, replaced_by_new_version, domain_id, domain_version) FROM stdin;
18	ab34afa0-6456-4163-81f5-f85c2fbea3c8	Custody Cash2bb	Systems connected to Custody Cash2bb	f	f	f	2021-10-11 16:25:04.378844+02	t	1	1
36	ab34afa0-6456-4163-81f5-f85c2fbea3c8	Custody Cash	Systems connected to Custody Cash	f	f	f	2021-10-11 16:25:04.378844+02	f	1	2
17	ab34afa0-6456-4163-81f5-f85c2fbea3c8	Custody Cash2b	Systems connected to Custody Cash2b	f	f	f	2021-10-07 20:30:48.494231+02	t	0	0
19	7b0105e1-1285-4c90-9480-cade1acec31c	Custody Arrangement2	2Custody Arrangement - Cobol2	f	f	f	2021-10-10 09:05:36.670262+02	t	0	0
20	828de370-2795-11ec-a0bd-9f11439f340c	name	description	t	t	f	2021-10-07 19:39:35.413661+02	t	0	0
22	828de370-2795-11ec-a0bd-9f11439f340c	nameb_mmm	descriptionb_llll	t	t	f	2021-10-08 21:18:34.103993+02	t	0	0
23	828de370-2795-11ec-a0bd-9f11439f340c	nameb_mmm_a	descriptionb_llll_a	t	t	f	2021-10-08 21:19:20.630565+02	t	0	0
9	ab34afa0-6456-4163-81f5-f85c2fbea3c8	Custody Cash2b	Systems connected to Custody Cash2b	f	f	f	2021-10-07 20:30:48.494231+02	t	0	0
15	ab34afa0-6456-4163-81f5-f85c2fbea3c8	Custody Cash2b	Systems connected to Custody Cash2b	f	f	f	2021-10-07 20:30:48.494231+02	t	0	0
13	ab34afa0-6456-4163-81f5-f85c2fbea3c8	Custody Cash2b	Systems connected to Custody Cash2b	f	f	f	2021-10-07 20:30:48.494231+02	t	0	0
2	7b0105e1-1285-4c90-9480-cade1acec31c	Custody Arrangement	Custody Arrangement - Cobol	f	f	f	2021-09-08 20:44:42.289+02	t	0	0
1	ab34afa0-6456-4163-81f5-f85c2fbea3c8	Custody Cash	Systems connected to Custody Cash	f	f	f	2021-09-08 20:41:36.339+02	t	0	0
7	ab34afa0-6456-4163-81f5-f85c2fbea3c8	Custody Cash2	Systems connected to Custody Cash2	f	f	f	2021-10-05 23:12:40.901925+02	t	0	0
21	828de370-2795-11ec-a0bd-9f11439f340c	nameb	descriptionb	t	t	f	2021-10-08 16:45:19.977438+02	t	0	0
14	ab34afa0-6456-4163-81f5-f85c2fbea3c8	Custody Cash2b	Systems connected to Custody Cash2b	f	f	f	2021-10-07 20:30:48.494231+02	t	0	0
24	828de370-2795-11ec-a0bd-9f11439f340c	nameb_mmm_ab	descriptionb_llll_a	t	t	f	2021-10-10 10:39:06.527082+02	t	0	0
25	22aa1e30-2941-11ec-8a51-58e9dd9e4d31	testnamn	beskrivning	f	t	t	2021-10-09 22:54:33.536667+02	f	0	0
27	b4077c40-2998-11ec-ba5b-ba233a60981c	Min nya tidzones-test 09:07 +0200	test	f	f	t	2021-10-10 10:13:10.361872+02	f	0	0
28	828de370-2795-11ec-a0bd-9f11439f340c	nameb_mmm_abc	descriptionb_llll_a	t	t	t	2021-10-11 06:50:22.356288+02	f	0	0
16	ab34afa0-6456-4163-81f5-f85c2fbea3c8	Custody Cash2b	Systems connected to Custody Cash2b	f	f	f	2021-10-07 20:30:48.494231+02	t	0	0
33	bf4b8780-2a93-11ec-a969-850b9a58e9c2	NNNNNNNNNNNnnnnnn	DDDDdddd	t	t	t	2021-10-11 16:24:00.122134+02	f	2	1
26	7b0105e1-1285-4c90-9480-cade1acec31c	Custody Arrangement2 0905	2Custody Arrangement - Cobol2	f	f	f	2021-10-11 16:24:25.341733+02	t	1	1
34	7b0105e1-1285-4c90-9480-cade1acec31c	Custody Arrangement 1624	Custody Arrangement - Cobol	f	f	f	2021-10-11 16:24:47.365205+02	t	1	2
35	7b0105e1-1285-4c90-9480-cade1acec31c	Custody Arrangement	Custody Arrangement - Cobol	f	t	f	2021-10-11 16:24:47.365205+02	f	1	3
\.


--
-- Data for Name: testinstructions; Type: TABLE DATA; Schema: public; Owner: caxdbuser
--

COPY public.testinstructions (id, guid, name, description, ready_for_use, activated, deleted, update_timestamp) FROM stdin;
\.


--
-- Name: magictable_metadata_Id_seq; Type: SEQUENCE SET; Schema: public; Owner: caxdbuser
--

SELECT pg_catalog.setval('public."magictable_metadata_Id_seq"', 1, false);


--
-- Name: tabletoedit_id_seq; Type: SEQUENCE SET; Schema: public; Owner: caxdbuser
--

SELECT pg_catalog.setval('public.tabletoedit_id_seq', 2, true);


--
-- Name: testdomains_id_seq; Type: SEQUENCE SET; Schema: public; Owner: caxdbuser
--

SELECT pg_catalog.setval('public.testdomains_id_seq', 1, false);


--
-- Name: testinstructions_id_seq; Type: SEQUENCE SET; Schema: public; Owner: caxdbuser
--

SELECT pg_catalog.setval('public.testinstructions_id_seq', 1, false);


--
-- Name: supported_metadata_tables supported_metadata_tables_pk; Type: CONSTRAINT; Schema: public; Owner: caxdbuser
--

ALTER TABLE ONLY public.supported_metadata_tables
    ADD CONSTRAINT supported_metadata_tables_pk PRIMARY KEY ("TableId");


--
-- Name: testdomains testdomains_pk; Type: CONSTRAINT; Schema: public; Owner: caxdbuser
--

ALTER TABLE ONLY public.testdomains
    ADD CONSTRAINT testdomains_pk PRIMARY KEY (id);


--
-- Name: testinstructions testinstructions_pk; Type: CONSTRAINT; Schema: public; Owner: caxdbuser
--

ALTER TABLE ONLY public.testinstructions
    ADD CONSTRAINT testinstructions_pk PRIMARY KEY (id);


--
-- Name: magictable_metadata_id_uindex; Type: INDEX; Schema: public; Owner: caxdbuser
--

CREATE UNIQUE INDEX magictable_metadata_id_uindex ON public.magictable_metadata USING btree ("Id");


--
-- Name: supported_metadata_tables_table_id_uindex; Type: INDEX; Schema: public; Owner: caxdbuser
--

CREATE UNIQUE INDEX supported_metadata_tables_table_id_uindex ON public.supported_metadata_tables USING btree ("TableId");


--
-- Name: supported_metadata_tables_tablename_uindex; Type: INDEX; Schema: public; Owner: caxdbuser
--

CREATE UNIQUE INDEX supported_metadata_tables_tablename_uindex ON public.supported_metadata_tables USING btree ("TableName");


--
-- Name: tabletoedit_grpc_api_identifier_uindex; Type: INDEX; Schema: public; Owner: caxdbuser
--

CREATE UNIQUE INDEX tabletoedit_grpc_api_identifier_uindex ON public.tabletoedit USING btree (grpc_api_identifier);


--
-- Name: tabletoedit_guid_uindex; Type: INDEX; Schema: public; Owner: caxdbuser
--

CREATE UNIQUE INDEX tabletoedit_guid_uindex ON public.tabletoedit USING btree (guid);


--
-- Name: tabletoedit_id_uindex; Type: INDEX; Schema: public; Owner: caxdbuser
--

CREATE UNIQUE INDEX tabletoedit_id_uindex ON public.tabletoedit USING btree (id);


--
-- Name: tabletoedit_table_name_uindex; Type: INDEX; Schema: public; Owner: caxdbuser
--

CREATE UNIQUE INDEX tabletoedit_table_name_uindex ON public.tabletoedit USING btree (table_name);


--
-- Name: testdomains_id_uindex; Type: INDEX; Schema: public; Owner: caxdbuser
--

CREATE UNIQUE INDEX testdomains_id_uindex ON public.testdomains USING btree (id);


--
-- Name: testinstructions_guid_uindex; Type: INDEX; Schema: public; Owner: caxdbuser
--

CREATE UNIQUE INDEX testinstructions_guid_uindex ON public.testinstructions USING btree (guid);


--
-- Name: testinstructions_id_uindex; Type: INDEX; Schema: public; Owner: caxdbuser
--

CREATE UNIQUE INDEX testinstructions_id_uindex ON public.testinstructions USING btree (id);


--
-- Name: testinstructions_name_uindex; Type: INDEX; Schema: public; Owner: caxdbuser
--

CREATE UNIQUE INDEX testinstructions_name_uindex ON public.testinstructions USING btree (name);


--
-- PostgreSQL database dump complete
--

