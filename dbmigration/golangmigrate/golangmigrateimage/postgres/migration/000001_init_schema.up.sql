--
-- Name: grids_states; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.grids_states (
    id integer NOT NULL,
    columns json,
    updated_at timestamp with time zone,
    created_at timestamp with time zone,
    updated_by integer,
    created_by integer,
    model_name character varying NOT NULL,
    grid_type character varying NOT NULL,
    groups json
);


--
-- Name: assets_crud_grids_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.assets_crud_grids_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: assets_crud_grids_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.assets_crud_grids_id_seq OWNED BY public.grids_states.id;


--
-- Name: dict_attack_classes; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.dict_attack_classes (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    updated_at timestamp with time zone,
    created_at timestamp with time zone,
    updated_by integer,
    created_by integer
);


--
-- Name: dict_attack_levels; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.dict_attack_levels (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    updated_at timestamp with time zone,
    created_at timestamp with time zone,
    updated_by integer,
    created_by integer
);


--
-- Name: dict_attack_methods; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.dict_attack_methods (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    recomendation_id integer NOT NULL,
    attack_class_id integer NOT NULL,
    attack_level_id integer NOT NULL,
    updated_at timestamp with time zone,
    created_at timestamp with time zone,
    updated_by integer,
    created_by integer
);


--
-- Name: dict_attack_methods_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.dict_attack_methods_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: dict_attack_methods_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.dict_attack_methods_id_seq OWNED BY public.dict_attack_methods.id;


--
-- Name: dict_attack_protocols; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.dict_attack_protocols (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    attack_level_id integer NOT NULL,
    default_t_protocol integer NOT NULL,
    default_port integer NOT NULL,
    updated_at timestamp with time zone,
    created_at timestamp with time zone,
    updated_by integer,
    created_by integer
);


--
-- Name: socbox_geo_cities; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.socbox_geo_cities (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    region_id integer,
    updated_by integer,
    created_by integer,
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);


--
-- Name: dict_cities_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.dict_cities_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: dict_cities_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.dict_cities_id_seq OWNED BY public.socbox_geo_cities.id;


--
-- Name: socbox_geo_districts; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.socbox_geo_districts (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    updated_by integer,
    created_by integer,
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);


--
-- Name: dict_districts_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.dict_districts_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: dict_districts_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.dict_districts_id_seq OWNED BY public.socbox_geo_districts.id;


--
-- Name: dict_event_statuses; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.dict_event_statuses (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    description text NOT NULL,
    updated_at timestamp with time zone,
    created_at timestamp with time zone,
    updated_by integer,
    created_by integer
);


--
-- Name: dict_event_statuses_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.dict_event_statuses_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: dict_event_statuses_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.dict_event_statuses_id_seq OWNED BY public.dict_event_statuses.id;


--
-- Name: dict_event_types; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.dict_event_types (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    updated_at timestamp with time zone,
    created_at timestamp with time zone,
    updated_by integer,
    created_by integer
);


--
-- Name: dict_misp_attribute_categories; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.dict_misp_attribute_categories (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    description character varying(255),
    updated_at timestamp with time zone,
    created_at timestamp with time zone,
    updated_by integer,
    created_by integer
);


--
-- Name: dict_misp_attribute_categories_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.dict_misp_attribute_categories_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: dict_misp_attribute_categories_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.dict_misp_attribute_categories_id_seq OWNED BY public.dict_misp_attribute_categories.id;


--
-- Name: dict_misp_attribute_types; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.dict_misp_attribute_types (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    description text,
    misp_attribute_category_id integer NOT NULL,
    updated_at timestamp with time zone,
    created_at timestamp with time zone,
    updated_by integer,
    created_by integer
);


--
-- Name: dict_misp_attribute_types_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.dict_misp_attribute_types_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: dict_misp_attribute_types_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.dict_misp_attribute_types_id_seq OWNED BY public.dict_misp_attribute_types.id;


--
-- Name: socbox_detection_mitre_tactics; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.socbox_detection_mitre_tactics (
    id integer NOT NULL,
    mitre_id character varying(255) NOT NULL,
    name character varying(255) NOT NULL,
    description text NOT NULL,
    reference text,
    updated_by integer,
    created_by integer,
    custom_fields_data jsonb DEFAULT '{}'::jsonb NOT NULL,
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);


--
-- Name: dict_mitre_tactics_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.dict_mitre_tactics_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: dict_mitre_tactics_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.dict_mitre_tactics_id_seq OWNED BY public.socbox_detection_mitre_tactics.id;


--
-- Name: socbox_detection_mitre_techniques; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.socbox_detection_mitre_techniques (
    id integer NOT NULL,
    mitre_id character varying(255) NOT NULL,
    name character varying(255) NOT NULL,
    description text NOT NULL,
    reference text NOT NULL,
    mitre_tactic_id integer NOT NULL,
    updated_by integer,
    created_by integer,
    custom_fields_data jsonb DEFAULT '{}'::jsonb NOT NULL,
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);


--
-- Name: dict_mitre_techniques_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.dict_mitre_techniques_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: dict_mitre_techniques_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.dict_mitre_techniques_id_seq OWNED BY public.socbox_detection_mitre_techniques.id;


--
-- Name: dict_person_modes; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.dict_person_modes (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    description text NOT NULL,
    updated_at timestamp with time zone,
    created_at timestamp with time zone,
    updated_by integer,
    created_by integer
);


--
-- Name: dict_person_orders; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.dict_person_orders (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    description text NOT NULL,
    updated_at timestamp with time zone,
    created_at timestamp with time zone,
    updated_by integer,
    created_by integer
);


--
-- Name: dict_person_roles; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.dict_person_roles (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    description text NOT NULL,
    updated_at timestamp with time zone,
    created_at timestamp with time zone,
    updated_by integer,
    created_by integer
);


--
-- Name: socbox_detection_recomendations; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.socbox_detection_recomendations (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    created_at timestamp without time zone,
    created_by integer,
    updated_at timestamp without time zone,
    updated_by integer,
    custom_fields_data jsonb DEFAULT '{}'::jsonb NOT NULL
);


--
-- Name: dict_recomendations_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.dict_recomendations_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: dict_recomendations_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.dict_recomendations_id_seq OWNED BY public.socbox_detection_recomendations.id;


--
-- Name: socbox_geo_regions; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.socbox_geo_regions (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    district_id integer,
    updated_by integer,
    created_by integer,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    geocode character varying
);


--
-- Name: dict_regions_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.dict_regions_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: dict_regions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.dict_regions_id_seq OWNED BY public.socbox_geo_regions.id;


--
-- Name: dict_soft_types; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.dict_soft_types (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    description text NOT NULL,
    updated_at timestamp with time zone,
    created_at timestamp with time zone,
    updated_by integer,
    created_by integer
);


--
-- Name: dict_softs; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.dict_softs (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    soft_type_id integer,
    updated_at timestamp with time zone,
    created_at timestamp with time zone,
    updated_by integer,
    created_by integer
);


--
-- Name: dict_softs_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.dict_softs_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: dict_softs_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.dict_softs_id_seq OWNED BY public.dict_softs.id;


--
-- Name: dict_vulnerabilities; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.dict_vulnerabilities (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    soft_id integer,
    attack_class_id integer NOT NULL,
    description character varying(255),
    updated_at timestamp with time zone,
    created_at timestamp with time zone,
    updated_by integer,
    created_by integer
);


--
-- Name: dict_vulnerabilities_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.dict_vulnerabilities_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: dict_vulnerabilities_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.dict_vulnerabilities_id_seq OWNED BY public.dict_vulnerabilities.id;


--
-- Name: event_attachments; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.event_attachments (
    id integer NOT NULL,
    event_id integer NOT NULL,
    name character varying(255) NOT NULL,
    filepath character varying(255) NOT NULL,
    updated_at timestamp with time zone,
    created_at timestamp with time zone,
    updated_by integer,
    created_by integer
);


--
-- Name: event_attachments_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.event_attachments_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: event_attachments_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.event_attachments_id_seq OWNED BY public.event_attachments.id;


--
-- Name: event_attributes; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.event_attributes (
    id integer NOT NULL,
    event_id integer NOT NULL,
    misp_attribute_type_id integer NOT NULL,
    value character varying(255) NOT NULL,
    updated_at timestamp with time zone,
    created_at timestamp with time zone,
    updated_by integer,
    created_by integer
);


--
-- Name: event_attributes_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.event_attributes_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: event_attributes_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.event_attributes_id_seq OWNED BY public.event_attributes.id;


--
-- Name: event_ext_ipaddresses; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.event_ext_ipaddresses (
    id integer NOT NULL,
    event_id integer NOT NULL,
    ip_address inet NOT NULL,
    country_alpha2 character varying(255) NOT NULL,
    city_id integer,
    updated_at timestamp with time zone,
    created_at timestamp with time zone,
    updated_by integer,
    created_by integer
);


--
-- Name: event_ext_ipaddresses_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.event_ext_ipaddresses_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: event_ext_ipaddresses_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.event_ext_ipaddresses_id_seq OWNED BY public.event_ext_ipaddresses.id;


--
-- Name: event_home_ipaddresses; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.event_home_ipaddresses (
    id integer NOT NULL,
    event_id integer NOT NULL,
    ip_address inet NOT NULL,
    domain character varying(255),
    public_ip_address inet,
    updated_at timestamp with time zone,
    created_at timestamp with time zone,
    updated_by integer,
    created_by integer
);


--
-- Name: event_home_ipaddresses_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.event_home_ipaddresses_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: event_home_ipaddresses_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.event_home_ipaddresses_id_seq OWNED BY public.event_home_ipaddresses.id;


--
-- Name: event_rules; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.event_rules (
    id integer NOT NULL,
    event_id integer NOT NULL,
    signatures_sid_id integer NOT NULL,
    signatures_gid_id integer NOT NULL,
    updated_at timestamp with time zone,
    created_at timestamp with time zone,
    updated_by integer,
    created_by integer,
    alert_count integer DEFAULT 1 NOT NULL
);


--
-- Name: event_rules_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.event_rules_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: event_rules_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.event_rules_id_seq OWNED BY public.event_rules.id;


--
-- Name: event_sources; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.event_sources (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    description text NOT NULL,
    updated_at timestamp with time zone,
    created_at timestamp with time zone,
    updated_by integer,
    created_by integer
);


--
-- Name: event_ttps; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.event_ttps (
    id integer NOT NULL,
    event_id integer NOT NULL,
    mitre_techniques_id integer NOT NULL,
    updated_at timestamp with time zone,
    created_at timestamp with time zone,
    updated_by integer,
    created_by integer
);


--
-- Name: event_ttps_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.event_ttps_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: event_ttps_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.event_ttps_id_seq OWNED BY public.event_ttps.id;


--
-- Name: events; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.events (
    id integer NOT NULL,
    updated_at timestamp with time zone,
    created_at timestamp with time zone,
    updated_by integer,
    created_by integer,
    e_start timestamp with time zone NOT NULL,
    e_end timestamp with time zone NOT NULL,
    description character varying(255) NOT NULL,
    nccci_classtype_id integer NOT NULL,
    recomendation_id integer NOT NULL,
    event_status_id integer NOT NULL,
    is_incident boolean NOT NULL,
    tags character varying(255),
    object_id integer NOT NULL,
    assigned_to integer,
    exported_to integer,
    external_system_id character varying(255)
);


--
-- Name: events_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.events_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: events_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.events_id_seq OWNED BY public.events.id;


--
-- Name: geoip; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.geoip (
    inetnum inet NOT NULL,
    netname character varying(255),
    description text,
    alpha2 character varying(255),
    city character varying(255),
    region character varying(255),
    asn integer,
    latitude numeric,
    longitude numeric,
    updated_at timestamp with time zone,
    created_at timestamp with time zone,
    updated_by integer,
    created_by integer,
    startip integer,
    endip integer
);


--
-- Name: iam_orgs; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.iam_orgs (
    created_at timestamp without time zone,
    created_by integer,
    updated_at timestamp without time zone,
    updated_by integer,
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    description text
);


--
-- Name: iam_orgs_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.iam_orgs_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: iam_orgs_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.iam_orgs_id_seq OWNED BY public.iam_orgs.id;


--
-- Name: iam_permissions; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.iam_permissions (
    id integer NOT NULL,
    action character varying NOT NULL,
    subject character varying NOT NULL,
    conditions jsonb,
    description character varying,
    inverted boolean DEFAULT false NOT NULL
);


--
-- Name: iam_permissions_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.iam_permissions_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: iam_permissions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.iam_permissions_id_seq OWNED BY public.iam_permissions.id;


--
-- Name: iam_roles; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.iam_roles (
    created_at timestamp without time zone,
    created_by integer,
    updated_at timestamp without time zone,
    updated_by integer,
    id integer NOT NULL,
    name character varying(50) NOT NULL,
    description character varying(255)
);


--
-- Name: iam_roles_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.iam_roles_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: iam_roles_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.iam_roles_id_seq OWNED BY public.iam_roles.id;


--
-- Name: iam_roles_permissions; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.iam_roles_permissions (
    role_id integer NOT NULL,
    permission_id integer NOT NULL
);


--
-- Name: iam_users; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.iam_users (
    created_at timestamp without time zone,
    created_by integer,
    updated_at timestamp without time zone,
    updated_by integer,
    id integer NOT NULL,
    login character varying(50) NOT NULL,
    password text,
    name1 character varying(255),
    name2 character varying(255),
    name3 character varying(255),
    settings jsonb DEFAULT '{"theme": {"darkMode": false}, "preferences": {}, "notifications": {}}'::jsonb NOT NULL,
    has_password boolean DEFAULT false NOT NULL,
    has_api_key boolean DEFAULT false NOT NULL,
    blocked boolean DEFAULT false NOT NULL,
    org_id integer
);


--
-- Name: iam_users_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.iam_users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: iam_users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.iam_users_id_seq OWNED BY public.iam_users.id;


--
-- Name: iam_users_roles; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.iam_users_roles (
    user_id integer NOT NULL,
    role_id integer NOT NULL
);


--
-- Name: iocs; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.iocs (
    created_at timestamp without time zone,
    created_by integer,
    updated_at timestamp without time zone,
    updated_by integer,
    custom_fields_data jsonb DEFAULT '{}'::jsonb NOT NULL,
    id integer NOT NULL,
    value character varying(255) NOT NULL,
    active boolean DEFAULT true NOT NULL,
    description character varying(255),
    ioc_type_id integer,
    ioc_source_id integer,
    threat_level_id integer,
    reference character varying(255),
    score integer,
    rating_up integer,
    rating_down integer,
    author character varying(255),
    first_seen date,
    tags character varying[] DEFAULT '{}'::character varying[],
    meta jsonb DEFAULT '{}'::jsonb NOT NULL
);


--
-- Name: iocs_sources; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.iocs_sources (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    created_at timestamp without time zone,
    created_by integer,
    updated_at timestamp without time zone,
    updated_by integer,
    custom_fields_data jsonb DEFAULT '{}'::jsonb NOT NULL,
    weigth integer,
    ioc_type_id integer,
    config jsonb,
    is_active boolean DEFAULT true NOT NULL,
    description text
);


--
-- Name: iocs_sources_statuses; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.iocs_sources_statuses (
    id integer NOT NULL,
    source_id integer NOT NULL,
    start_time timestamp with time zone NOT NULL,
    end_time timestamp with time zone,
    duration numeric(10,3),
    processed integer DEFAULT 0,
    added integer DEFAULT 0,
    updated integer DEFAULT 0,
    skipped integer DEFAULT 0,
    errors integer DEFAULT 0,
    success boolean DEFAULT false NOT NULL,
    error_details jsonb DEFAULT '[]'::jsonb,
    run_metadata jsonb DEFAULT '{}'::jsonb,
    created_at timestamp with time zone DEFAULT now()
);


--
-- Name: TABLE iocs_sources_statuses; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON TABLE public.iocs_sources_statuses IS 'История выполнения коллектора IOC';


--
-- Name: COLUMN iocs_sources_statuses.source_id; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.iocs_sources_statuses.source_id IS 'ID источника из socbox_detection_iocs_sources';


--
-- Name: COLUMN iocs_sources_statuses.start_time; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.iocs_sources_statuses.start_time IS 'Время начала выполнения';


--
-- Name: COLUMN iocs_sources_statuses.end_time; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.iocs_sources_statuses.end_time IS 'Время окончания выполнения';


--
-- Name: COLUMN iocs_sources_statuses.duration; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.iocs_sources_statuses.duration IS 'Длительность выполнения в секундах';


--
-- Name: COLUMN iocs_sources_statuses.processed; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.iocs_sources_statuses.processed IS 'Количество обработанных IOC';


--
-- Name: COLUMN iocs_sources_statuses.added; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.iocs_sources_statuses.added IS 'Количество добавленных IOC';


--
-- Name: COLUMN iocs_sources_statuses.updated; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.iocs_sources_statuses.updated IS 'Количество обновленных IOC';


--
-- Name: COLUMN iocs_sources_statuses.skipped; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.iocs_sources_statuses.skipped IS 'Количество пропущенных IOC';


--
-- Name: COLUMN iocs_sources_statuses.errors; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.iocs_sources_statuses.errors IS 'Количество ошибок';


--
-- Name: COLUMN iocs_sources_statuses.success; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.iocs_sources_statuses.success IS 'Успешность выполнения';


--
-- Name: COLUMN iocs_sources_statuses.error_details; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.iocs_sources_statuses.error_details IS 'Детали ошибок в формате JSON';


--
-- Name: COLUMN iocs_sources_statuses.run_metadata; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.iocs_sources_statuses.run_metadata IS 'Дополнительные метаданные выполнения';


--
-- Name: iocs_types; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.iocs_types (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    created_at timestamp without time zone,
    created_by integer,
    updated_at timestamp without time zone,
    updated_by integer,
    custom_fields_data jsonb DEFAULT '{}'::jsonb NOT NULL,
    gid_id integer,
    weigth integer,
    description character varying(255),
    mappers jsonb
);


--
-- Name: iocs_whitelist; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.iocs_whitelist (
    id integer NOT NULL,
    value character varying(255) NOT NULL,
    "createdAt" timestamp without time zone,
    created_by integer,
    "updatedAt" timestamp without time zone,
    updated_by integer,
    description character varying(255),
    "objectId" integer,
    srcip character varying(255),
    dstip character varying(255),
    "ioctypeId" integer
);


--
-- Name: socbox_actors_objects; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.socbox_actors_objects (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    updated_by integer,
    created_by integer,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    custom_fields_data jsonb DEFAULT '{}'::jsonb NOT NULL,
    full_name character varying,
    latitude character varying,
    longitude character varying,
    disabled_at integer,
    city_id integer,
    subject_id integer,
    scope_id integer,
    is_kii boolean
);


--
-- Name: objects_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.objects_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: objects_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.objects_id_seq OWNED BY public.socbox_actors_objects.id;


--
-- Name: objects_networks; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.objects_networks (
    id integer NOT NULL,
    object_id integer,
    ip_net integer,
    start_to_ip_on_server integer,
    fin_to_ip_on_server integer,
    start_to_ip integer,
    fin_to_ip integer,
    start_to_int_on_server integer,
    fin_to_int_on_server integer,
    start_to_int integer,
    fin_to_int integer,
    family_type integer,
    description character varying(255),
    updated_at timestamp with time zone,
    created_at timestamp with time zone,
    updated_by integer,
    created_by integer
);


--
-- Name: objects_networks_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.objects_networks_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: objects_networks_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.objects_networks_id_seq OWNED BY public.objects_networks.id;


--
-- Name: objects_persons; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.objects_persons (
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    object_id integer NOT NULL,
    person_id integer NOT NULL
);


--
-- Name: socbox_assets_resources_types; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.socbox_assets_resources_types (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    description character varying(255) NOT NULL,
    updated_at timestamp with time zone,
    created_at timestamp with time zone,
    updated_by integer,
    created_by integer
);


--
-- Name: objects_resource_types_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.objects_resource_types_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: objects_resource_types_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.objects_resource_types_id_seq OWNED BY public.socbox_assets_resources_types.id;


--
-- Name: socbox_assets_resources; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.socbox_assets_resources (
    id integer NOT NULL,
    domain character varying(255) NOT NULL,
    url character varying(255),
    object_resource_type_id integer NOT NULL,
    description text NOT NULL,
    object_homenets_id integer NOT NULL,
    ip_address inet NOT NULL,
    local_ip_address inet NOT NULL,
    updated_at timestamp with time zone,
    created_at timestamp with time zone,
    updated_by integer,
    created_by integer
);


--
-- Name: objects_resources_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.objects_resources_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: objects_resources_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.objects_resources_id_seq OWNED BY public.socbox_assets_resources.id;


--
-- Name: orgs; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.orgs (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    description text,
    updated_at timestamp with time zone,
    created_at timestamp with time zone,
    updated_by integer,
    created_by integer
);


--
-- Name: orgs_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.orgs_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: orgs_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.orgs_id_seq OWNED BY public.orgs.id;


--
-- Name: persons; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.persons (
    id integer NOT NULL,
    name character varying(255),
    surname character varying(255),
    patronymic character varying(255),
    "position" character varying(255) NOT NULL,
    work_phone character varying(255),
    mobile_phone character varying(255),
    email character varying(255),
    order_id integer,
    mode_id integer,
    object_id integer,
    updated_at timestamp with time zone,
    created_at timestamp with time zone,
    updated_by integer,
    created_by integer
);


--
-- Name: persons_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.persons_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: persons_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.persons_id_seq OWNED BY public.persons.id;



--
-- Name: socbox_detection_sensors; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.socbox_detection_sensors (
    id integer NOT NULL,
    created_at timestamp without time zone,
    created_by integer,
    updated_at timestamp without time zone,
    updated_by integer,
    name character varying(255) NOT NULL,
    custom_fields_data jsonb DEFAULT '{}'::jsonb NOT NULL,
    hardware_id integer,
    group_id integer,
    is_org_property boolean,
    serial_number character varying,
    inventory_number character varying
);


--
-- Name: sensors_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.sensors_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: sensors_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.sensors_id_seq OWNED BY public.socbox_detection_sensors.id;


--
-- Name: sensors_settings; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.sensors_settings (
    id integer NOT NULL,
    alert jsonb,
    capture jsonb,
    check_sums jsonb,
    disk jsonb,
    dnsflow jsonb,
    netflow jsonb,
    moth jsonb,
    sensor jsonb,
    server jsonb,
    snort jsonb,
    status jsonb,
    storcap jsonb,
    suricata jsonb,
    storage jsonb,
    storindex jsonb,
    hash jsonb,
    updated_at timestamp with time zone,
    created_at timestamp with time zone,
    updated_by integer,
    created_by integer
);


--
-- Name: siem_actions; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.siem_actions (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    description character varying(255) NOT NULL,
    siem_viewed_id integer NOT NULL,
    roles jsonb NOT NULL,
    updated_at timestamp with time zone,
    created_at timestamp with time zone,
    updated_by integer,
    created_by integer
);


--
-- Name: siem_errors; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.siem_errors (
    id integer NOT NULL,
    description character varying(255) NOT NULL,
    cause text NOT NULL,
    updated_at timestamp with time zone,
    created_at timestamp with time zone,
    updated_by integer,
    created_by integer
);


--
-- Name: siem_external_systems; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.siem_external_systems (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    description text NOT NULL,
    settings jsonb NOT NULL,
    active boolean NOT NULL,
    updated_at timestamp with time zone,
    created_at timestamp with time zone,
    updated_by integer,
    created_by integer
);


--
-- Name: siem_external_systems_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.siem_external_systems_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: siem_external_systems_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.siem_external_systems_id_seq OWNED BY public.siem_external_systems.id;


--
-- Name: siem_filter_types; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.siem_filter_types (
    id integer NOT NULL,
    type character varying(255) NOT NULL,
    description text,
    updated_at timestamp with time zone,
    created_at timestamp with time zone,
    updated_by integer,
    created_by integer
);


--
-- Name: siem_filter_types_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.siem_filter_types_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: siem_filter_types_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.siem_filter_types_id_seq OWNED BY public.siem_filter_types.id;


--
-- Name: siem_filters; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.siem_filters (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    description text,
    siem_filter_type_id integer NOT NULL,
    updated_at timestamp with time zone,
    created_at timestamp with time zone,
    updated_by integer,
    created_by integer
);


--
-- Name: siem_filters_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.siem_filters_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: siem_filters_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.siem_filters_id_seq OWNED BY public.siem_filters.id;


--
-- Name: siem_task_statuses; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.siem_task_statuses (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    description text NOT NULL,
    updated_at timestamp with time zone,
    created_at timestamp with time zone,
    updated_by integer,
    created_by integer
);


--
-- Name: siem_task_statuses_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.siem_task_statuses_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: siem_task_statuses_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.siem_task_statuses_id_seq OWNED BY public.siem_task_statuses.id;


--
-- Name: siem_task_types; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.siem_task_types (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    description text NOT NULL,
    updated_at timestamp with time zone,
    created_at timestamp with time zone,
    updated_by integer,
    created_by integer
);


--
-- Name: siem_task_types_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.siem_task_types_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: siem_task_types_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.siem_task_types_id_seq OWNED BY public.siem_task_types.id;


--
-- Name: siem_tasks; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.siem_tasks (
    id integer NOT NULL,
    siem_task_type_id integer NOT NULL,
    description character varying(255) NOT NULL,
    siem_task_status_id integer NOT NULL,
    raw_data jsonb,
    updated_at timestamp with time zone,
    created_at timestamp with time zone,
    updated_by integer,
    created_by integer,
    assigned_to integer,
    event_id integer
);


--
-- Name: siem_tasks_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.siem_tasks_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: siem_tasks_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.siem_tasks_id_seq OWNED BY public.siem_tasks.id;


--
-- Name: siem_views; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.siem_views (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    description character varying(255) NOT NULL,
    roles jsonb NOT NULL,
    updated_at timestamp with time zone,
    created_at timestamp with time zone,
    updated_by integer,
    created_by integer
);


--
-- Name: signatures_overlay; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.signatures_overlay (
    id integer NOT NULL,
    sid integer NOT NULL,
    rev integer NOT NULL,
    use boolean NOT NULL,
    comment text NOT NULL,
    sensor_group_id integer NOT NULL,
    updated_at timestamp with time zone,
    created_at timestamp with time zone,
    updated_by integer,
    created_by integer
);


--
-- Name: signatures_overlay_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.signatures_overlay_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: signatures_overlay_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.signatures_overlay_id_seq OWNED BY public.signatures_overlay.id;


--
-- Name: signatures_test; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.signatures_test (
    id integer NOT NULL,
    i_sid integer NOT NULL,
    i_rev integer,
    i_b_rec integer,
    i_b_com integer,
    i_add_date integer,
    i_created_at integer,
    s_classtype text,
    i_priority integer,
    s_rule_type text,
    s_source_name text,
    s_msg text,
    s_content jsonb,
    s_flow text,
    s_threshold text,
    s_pcre text,
    s_cve text,
    s_ref text,
    s_bugtraq text,
    s_content_hash text,
    s_hash_light text,
    i_commented_date integer,
    i_exclusion_date integer,
    i_exclusion_date_from_rec integer,
    s_created_at_date text,
    s_updated_at_date text,
    i_b_was_in_last_package integer,
    i_b_get_traffic integer,
    i_b_ex_from_source integer,
    b_flowbit_noalert integer,
    s_flowbit_set text,
    s_flowbit_isset text,
    s_flowbit_isnotset text,
    s_flowbit_unset text,
    i_sid_percent integer,
    i_sid_true integer,
    i_sid_false integer,
    i_sid_unknown integer,
    s_rules_which_replace_this text,
    s_rules_which_replaced_by_this text,
    s_comment text,
    s_rule_body text,
    s_add_date integer,
    i_b_commented integer,
    i_exclusion_date_from_cur text,
    s_replacing_package_name text,
    s_users text,
    s_commented_date text,
    s_exclusion_date text,
    s_exclusion_date_from_cur text,
    rating_up integer,
    rating_down integer,
    local_classtype_id integer,
    nccci_classtype_id integer,
    classtype_id integer,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    created_by integer,
    updated_by integer,
    i_gid integer
);


--
-- Name: signatures_test_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.signatures_test_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: signatures_test_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.signatures_test_id_seq OWNED BY public.signatures_test.id;


--
-- Name: socbox_actors_objects_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.socbox_actors_objects_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: socbox_actors_objects_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.socbox_actors_objects_id_seq OWNED BY public.socbox_actors_objects.id;


--
-- Name: socbox_actors_objects_sensors; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.socbox_actors_objects_sensors (
    sensor_id integer NOT NULL,
    object_id integer NOT NULL
);


--
-- Name: socbox_actors_persons_orders; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.socbox_actors_persons_orders (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    created_at timestamp without time zone,
    created_by integer,
    updated_at timestamp without time zone,
    updated_by integer,
    custom_fields_data jsonb DEFAULT '{}'::jsonb NOT NULL
);


--
-- Name: socbox_actors_persons_orders_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.socbox_actors_persons_orders_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: socbox_actors_persons_orders_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.socbox_actors_persons_orders_id_seq OWNED BY public.socbox_actors_persons_orders.id;


--
-- Name: socbox_actors_scopes; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.socbox_actors_scopes (
    created_at timestamp without time zone,
    created_by integer,
    updated_at timestamp without time zone,
    updated_by integer,
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    custom_fields_data jsonb DEFAULT '{}'::jsonb NOT NULL
);


--
-- Name: socbox_actors_scopes_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.socbox_actors_scopes_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: socbox_actors_scopes_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.socbox_actors_scopes_id_seq OWNED BY public.socbox_actors_scopes.id;


--
-- Name: socbox_actors_subjects; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.socbox_actors_subjects (
    id serial,
    name character varying(255) NOT NULL,
    city_id integer,
    updated_by integer,
    created_by integer,
    custom_fields_data jsonb DEFAULT '{}'::jsonb NOT NULL,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    full_name character varying,
    address character varying,
    inn character varying,
    email character varying
);


--
-- Name: socbox_actors_subjects_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

-- CREATE SEQUENCE public.socbox_actors_subjects_id_seq
--     START WITH 1
--     INCREMENT BY 1
--     NO MINVALUE
--     NO MAXVALUE
--     CACHE 1;


--
-- Name: socbox_actors_subjects_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

-- ALTER SEQUENCE public.socbox_actors_subjects_id_seq OWNED BY public.socbox_actors_subjects.id;


--
-- Name: socbox_assets_homenets; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.socbox_assets_homenets (
    id integer NOT NULL,
    network inet NOT NULL,
    bandwidth integer,
    object_id integer,
    comment character varying(255),
    updated_at timestamp with time zone,
    created_at timestamp with time zone,
    updated_by integer,
    created_by integer
);


--
-- Name: socbox_assets_homenets_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.socbox_assets_homenets_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: socbox_assets_homenets_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.socbox_assets_homenets_seq OWNED BY public.socbox_assets_homenets.id;


--
-- Name: socbox_content_type; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.socbox_content_type (
    id integer NOT NULL,
    app_label character varying(100) NOT NULL,
    model character varying(100) NOT NULL,
    name character varying(100) NOT NULL,
    name_plural character varying(100) NOT NULL,
    name_accusative character varying(100)
);


--
-- Name: socbox_content_type_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.socbox_content_type_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: socbox_content_type_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.socbox_content_type_id_seq OWNED BY public.socbox_content_type.id;


--
-- Name: socbox_core_content_type; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.socbox_core_content_type (
    id integer NOT NULL,
    app_label character varying(100) NOT NULL,
    model character varying(100) NOT NULL,
    name character varying(100) NOT NULL,
    name_plural character varying(100) NOT NULL,
    name_accusative character varying(100)
);


--
-- Name: socbox_core_content_type_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.socbox_core_content_type_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: socbox_core_content_type_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.socbox_core_content_type_id_seq OWNED BY public.socbox_core_content_type.id;


--
-- Name: socbox_detection_attacks_priorities; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.socbox_detection_attacks_priorities (
    created_at timestamp without time zone,
    created_by integer,
    updated_at timestamp without time zone,
    updated_by integer,
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    custom_fields_data jsonb DEFAULT '{}'::jsonb NOT NULL,
    description text
);


--
-- Name: socbox_detection_classtypes; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.socbox_detection_classtypes (
    created_at timestamp without time zone,
    created_by integer,
    updated_at timestamp without time zone,
    updated_by integer,
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    custom_fields_data jsonb DEFAULT '{}'::jsonb NOT NULL,
    description character varying(255)
);


--
-- Name: socbox_detection_classtypes_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.socbox_detection_classtypes_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: socbox_detection_classtypes_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.socbox_detection_classtypes_id_seq OWNED BY public.socbox_detection_classtypes.id;


--
-- Name: socbox_detection_directions; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.socbox_detection_directions (
    created_at timestamp without time zone,
    created_by integer,
    updated_at timestamp without time zone,
    updated_by integer,
    custom_fields_data jsonb DEFAULT '{}'::jsonb NOT NULL,
    id integer NOT NULL,
    name character varying(255) NOT NULL
);


--
-- Name: socbox_detection_iocs_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.socbox_detection_iocs_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: socbox_detection_iocs_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.socbox_detection_iocs_id_seq OWNED BY public.iocs.id;


--
-- Name: socbox_detection_iocs_sources_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.socbox_detection_iocs_sources_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: socbox_detection_iocs_sources_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.socbox_detection_iocs_sources_id_seq OWNED BY public.iocs_sources.id;


--
-- Name: socbox_detection_iocs_sources_statuses_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.socbox_detection_iocs_sources_statuses_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: socbox_detection_iocs_sources_statuses_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.socbox_detection_iocs_sources_statuses_id_seq OWNED BY public.iocs_sources_statuses.id;


--
-- Name: socbox_detection_iocs_types_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.socbox_detection_iocs_types_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: socbox_detection_iocs_types_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.socbox_detection_iocs_types_id_seq OWNED BY public.iocs_types.id;


--
-- Name: socbox_detection_iocs_whitelist_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.socbox_detection_iocs_whitelist_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: socbox_detection_iocs_whitelist_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.socbox_detection_iocs_whitelist_id_seq OWNED BY public.iocs_whitelist.id;


--
-- Name: socbox_detection_local_classtypes; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.socbox_detection_local_classtypes (
    created_at timestamp without time zone,
    created_by integer,
    updated_at timestamp without time zone,
    updated_by integer,
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    custom_fields_data jsonb DEFAULT '{}'::jsonb NOT NULL,
    description text
);


--
-- Name: socbox_detection_local_classtypes_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.socbox_detection_local_classtypes_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: socbox_detection_local_classtypes_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.socbox_detection_local_classtypes_id_seq OWNED BY public.socbox_detection_local_classtypes.id;


--
-- Name: socbox_detection_mitre_tactics_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.socbox_detection_mitre_tactics_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: socbox_detection_mitre_tactics_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.socbox_detection_mitre_tactics_id_seq OWNED BY public.socbox_detection_mitre_tactics.id;


--
-- Name: socbox_detection_mitre_techniques_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.socbox_detection_mitre_techniques_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: socbox_detection_mitre_techniques_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.socbox_detection_mitre_techniques_id_seq OWNED BY public.socbox_detection_mitre_techniques.id;


--
-- Name: socbox_detection_nccci_classes; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.socbox_detection_nccci_classes (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    created_at timestamp without time zone,
    created_by integer,
    updated_at timestamp without time zone,
    updated_by integer,
    custom_fields_data jsonb DEFAULT '{}'::jsonb NOT NULL,
    description text,
    attack_method_id integer,
    is_incident boolean DEFAULT false NOT NULL
);


--
-- Name: socbox_detection_nccci_classes_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.socbox_detection_nccci_classes_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: socbox_detection_nccci_classes_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.socbox_detection_nccci_classes_id_seq OWNED BY public.socbox_detection_nccci_classes.id;


--
-- Name: socbox_detection_recomendations_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.socbox_detection_recomendations_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: socbox_detection_recomendations_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.socbox_detection_recomendations_id_seq OWNED BY public.socbox_detection_recomendations.id;


--
-- Name: socbox_detection_sensors_groups; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.socbox_detection_sensors_groups (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    created_at timestamp without time zone,
    created_by integer,
    updated_at timestamp without time zone,
    updated_by integer,
    custom_fields_data jsonb DEFAULT '{}'::jsonb NOT NULL
);


--
-- Name: socbox_detection_sensors_groups_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.socbox_detection_sensors_groups_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: socbox_detection_sensors_groups_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.socbox_detection_sensors_groups_id_seq OWNED BY public.socbox_detection_sensors_groups.id;


--
-- Name: socbox_detection_sensors_hardwares; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.socbox_detection_sensors_hardwares (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    created_at timestamp without time zone,
    created_by integer,
    updated_at timestamp without time zone,
    updated_by integer,
    custom_fields_data jsonb DEFAULT '{}'::jsonb NOT NULL
);


--
-- Name: socbox_detection_sensors_hardwares_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.socbox_detection_sensors_hardwares_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: socbox_detection_sensors_hardwares_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.socbox_detection_sensors_hardwares_id_seq OWNED BY public.socbox_detection_sensors_hardwares.id;


--
-- Name: socbox_detection_signatures; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.socbox_detection_signatures (
    created_at timestamp without time zone,
    created_by integer,
    updated_at timestamp without time zone,
    updated_by integer,
    custom_fields_data jsonb DEFAULT '{}'::jsonb NOT NULL,
    sid bigint NOT NULL,
    gid integer NOT NULL,
    rev integer NOT NULL,
    priority integer NOT NULL,
    classtype character varying(100) NOT NULL,
    msg text NOT NULL,
    rule text NOT NULL,
    description text,
    rec boolean DEFAULT false NOT NULL,
    source_id integer NOT NULL,
    content_hash character varying(64) NOT NULL,
    metadata jsonb,
    is_active boolean DEFAULT true NOT NULL,
    rating_up integer DEFAULT 0 NOT NULL,
    rating_down integer DEFAULT 0 NOT NULL
);


--
-- Name: socbox_detection_signatures_classtypes; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.socbox_detection_signatures_classtypes (
    created_at timestamp without time zone,
    created_by integer,
    updated_at timestamp without time zone,
    updated_by integer,
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    custom_fields_data jsonb DEFAULT '{}'::jsonb NOT NULL,
    description text
);


--
-- Name: socbox_detection_signatures_classtypes_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.socbox_detection_signatures_classtypes_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: socbox_detection_signatures_classtypes_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.socbox_detection_signatures_classtypes_id_seq OWNED BY public.socbox_detection_signatures_classtypes.id;


--
-- Name: socbox_detection_threat_levels; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.socbox_detection_threat_levels (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    description text NOT NULL,
    updated_at timestamp with time zone,
    created_at timestamp with time zone,
    updated_by integer,
    created_by integer
);


--
-- Name: socbox_extras_changelog; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.socbox_extras_changelog (
    id integer NOT NULL,
    "time" timestamp without time zone NOT NULL,
    user_id integer NOT NULL,
    request_id uuid NOT NULL,
    action character varying(50) NOT NULL,
    object_name character varying NOT NULL,
    prechange_data jsonb,
    postchange_data jsonb,
    object_id integer,
    object_type_id integer
);


--
-- Name: socbox_extras_changelog_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.socbox_extras_changelog_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: socbox_extras_changelog_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.socbox_extras_changelog_id_seq OWNED BY public.socbox_extras_changelog.id;


--
-- Name: socbox_extras_custom_fields; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.socbox_extras_custom_fields (
    created_at timestamp without time zone,
    created_by integer,
    updated_at timestamp without time zone,
    updated_by integer,
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    label character varying(50),
    type character varying(50) NOT NULL,
    description character varying(200),
    required boolean,
    filter_logic character varying(50),
    "default" jsonb,
    weight integer NOT NULL,
    validation_minimum integer,
    validation_maximum integer,
    validation_regex character varying(500),
    "groupName" character varying(50),
    search_weight integer,
    is_cloneable boolean,
    ui_editable character varying(50) NOT NULL,
    ui_visible character varying(50) NOT NULL,
    comments text,
    related_object_type_id integer,
    choiceset_id integer,
    custom_fields_data jsonb DEFAULT '{}'::jsonb NOT NULL
);


--
-- Name: socbox_extras_custom_fields_choicesets; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.socbox_extras_custom_fields_choicesets (
    created_at timestamp without time zone,
    created_by integer,
    updated_at timestamp without time zone,
    updated_by integer,
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    custom_fields_data jsonb DEFAULT '{}'::jsonb NOT NULL,
    description character varying(200),
    choices character varying[] NOT NULL,
    order_alphabetically boolean NOT NULL
);


--
-- Name: socbox_extras_custom_fields_choicesets_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.socbox_extras_custom_fields_choicesets_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: socbox_extras_custom_fields_choicesets_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.socbox_extras_custom_fields_choicesets_id_seq OWNED BY public.socbox_extras_custom_fields_choicesets.id;


--
-- Name: socbox_extras_custom_fields_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.socbox_extras_custom_fields_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: socbox_extras_custom_fields_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.socbox_extras_custom_fields_id_seq OWNED BY public.socbox_extras_custom_fields.id;


--
-- Name: socbox_extras_custom_fields_objects_types; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.socbox_extras_custom_fields_objects_types (
    id integer NOT NULL,
    custom_field_id integer NOT NULL,
    object_type_id integer NOT NULL
);


--
-- Name: socbox_extras_custom_fields_objects_types_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.socbox_extras_custom_fields_objects_types_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: socbox_extras_custom_fields_objects_types_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.socbox_extras_custom_fields_objects_types_id_seq OWNED BY public.socbox_extras_custom_fields_objects_types.id;


--
-- Name: socbox_extras_menu; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.socbox_extras_menu (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    created_at timestamp without time zone,
    created_by integer,
    updated_at timestamp without time zone,
    updated_by integer,
    data jsonb NOT NULL
);


--
-- Name: socbox_extras_menu_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.socbox_extras_menu_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: socbox_extras_menu_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.socbox_extras_menu_id_seq OWNED BY public.socbox_extras_menu.id;


--
-- Name: socbox_extras_tag_object_types; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.socbox_extras_tag_object_types (
    id bigint NOT NULL,
    tag_id integer NOT NULL,
    object_type_id integer NOT NULL
);


--
-- Name: socbox_extras_tag_object_types_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.socbox_extras_tag_object_types_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: socbox_extras_tag_object_types_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.socbox_extras_tag_object_types_id_seq OWNED BY public.socbox_extras_tag_object_types.id;


--
-- Name: socbox_extras_tagged_item; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.socbox_extras_tagged_item (
    id bigint NOT NULL,
    object_type_id integer NOT NULL,
    object_id integer NOT NULL,
    tag_id integer NOT NULL
);


--
-- Name: socbox_extras_tagged_item_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.socbox_extras_tagged_item_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: socbox_extras_tagged_item_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.socbox_extras_tagged_item_id_seq OWNED BY public.socbox_extras_tagged_item.id;


--
-- Name: socbox_extras_tags; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.socbox_extras_tags (
    created_at timestamp without time zone,
    created_by integer,
    updated_at timestamp without time zone,
    updated_by integer,
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    color character varying(6),
    description character varying(200)
);


--
-- Name: socbox_extras_tags_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.socbox_extras_tags_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: socbox_extras_tags_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.socbox_extras_tags_id_seq OWNED BY public.socbox_extras_tags.id;


--
-- Name: socbox_geo_cities_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.socbox_geo_cities_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: socbox_geo_cities_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.socbox_geo_cities_id_seq OWNED BY public.socbox_geo_cities.id;


--
-- Name: socbox_geo_countries; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.socbox_geo_countries (
    created_at timestamp without time zone,
    created_by integer,
    updated_at timestamp without time zone,
    updated_by integer,
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    custom_fields_data jsonb DEFAULT '{}'::jsonb NOT NULL,
    alpha2 character varying(2) NOT NULL,
    alpha3 character varying(3) NOT NULL,
    iso integer,
    full_name character varying(255),
    location_precise character varying(255)
);


--
-- Name: socbox_geo_countries_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.socbox_geo_countries_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: socbox_geo_countries_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.socbox_geo_countries_id_seq OWNED BY public.socbox_geo_countries.id;


--
-- Name: socbox_geo_districts_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.socbox_geo_districts_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: socbox_geo_districts_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.socbox_geo_districts_id_seq OWNED BY public.socbox_geo_districts.id;


--
-- Name: socbox_geo_regions_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.socbox_geo_regions_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: socbox_geo_regions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.socbox_geo_regions_id_seq OWNED BY public.socbox_geo_regions.id;


--
-- Name: statistics; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.statistics (
    sid integer NOT NULL,
    involvement_in_mal_infr integer DEFAULT 0,
    slowdown_d_do_s integer DEFAULT 0,
    infection_malware integer DEFAULT 0,
    network_traffic_capture integer DEFAULT 0,
    using_for_phishing integer DEFAULT 0,
    account_compromise integer DEFAULT 0,
    unauthorized_change_info integer DEFAULT 0,
    unauthorized_disclosure_info integer DEFAULT 0,
    publication_prohibited_info integer DEFAULT 0,
    sending_s_p_a_m integer DEFAULT 0,
    successful_exploitation integer DEFAULT 0,
    d_do_s_attack integer DEFAULT 0,
    failed_auth_attemp integer DEFAULT 0,
    attempt_introduce_malware integer DEFAULT 0,
    attempt_exploitation integer DEFAULT 0,
    publication_fraudulent_info integer DEFAULT 0,
    network_scan integer DEFAULT 0,
    social_engineering integer DEFAULT 0,
    is_incident integer DEFAULT 0,
    total_report integer DEFAULT 0,
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);


--
-- Name: subjects_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.subjects_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: subjects_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.subjects_id_seq OWNED BY public.socbox_actors_subjects.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.users (
    id integer NOT NULL,
    login character varying(255) NOT NULL,
    password character varying(255),
    name1 character varying(255),
    name2 character varying(255),
    name3 character varying(255),
    roles character varying[] NOT NULL,
    org_id integer NOT NULL,
    has_password boolean,
    has_api_key boolean,
    is_duty boolean,
    blocked boolean,
    settings jsonb,
    updated_at timestamp with time zone,
    created_at timestamp with time zone,
    updated_by integer,
    created_by integer,
    email character varying
);


--
-- Name: users_actions; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.users_actions (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    updated_at timestamp with time zone,
    created_at timestamp with time zone,
    updated_by integer,
    created_by integer
);


--
-- Name: users_actions_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.users_actions_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: users_actions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.users_actions_id_seq OWNED BY public.users_actions.id;


--
-- Name: users_histories; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.users_histories (
    id integer NOT NULL,
    user_id integer NOT NULL,
    expression character varying(255) NOT NULL,
    updated_at timestamp with time zone,
    created_at timestamp with time zone,
    updated_by integer,
    created_by integer
);


--
-- Name: users_histories_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.users_histories_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: users_histories_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.users_histories_id_seq OWNED BY public.users_histories.id;


--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: users_metrics; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.users_metrics (
    id integer NOT NULL,
    user_action_id integer NOT NULL,
    user_id integer NOT NULL,
    object_id integer,
    rule_id integer,
    updated_at timestamp with time zone,
    created_at timestamp with time zone,
    updated_by integer,
    created_by integer
);


--
-- Name: users_metrics_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.users_metrics_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: users_metrics_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.users_metrics_id_seq OWNED BY public.users_metrics.id;


--
-- Name: users_roles; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.users_roles (
    id integer NOT NULL,
    name character varying(255),
    description text,
    default_siem_action jsonb,
    updated_at timestamp with time zone,
    created_at timestamp with time zone,
    updated_by integer,
    created_by integer
);


--
-- Name: users_roles_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.users_roles_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: users_roles_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.users_roles_id_seq OWNED BY public.users_roles.id;


--
-- Name: dict_attack_methods id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.dict_attack_methods ALTER COLUMN id SET DEFAULT nextval('public.dict_attack_methods_id_seq'::regclass);


--
-- Name: dict_event_statuses id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.dict_event_statuses ALTER COLUMN id SET DEFAULT nextval('public.dict_event_statuses_id_seq'::regclass);


--
-- Name: dict_misp_attribute_categories id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.dict_misp_attribute_categories ALTER COLUMN id SET DEFAULT nextval('public.dict_misp_attribute_categories_id_seq'::regclass);


--
-- Name: dict_misp_attribute_types id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.dict_misp_attribute_types ALTER COLUMN id SET DEFAULT nextval('public.dict_misp_attribute_types_id_seq'::regclass);


--
-- Name: dict_softs id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.dict_softs ALTER COLUMN id SET DEFAULT nextval('public.dict_softs_id_seq'::regclass);


--
-- Name: dict_vulnerabilities id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.dict_vulnerabilities ALTER COLUMN id SET DEFAULT nextval('public.dict_vulnerabilities_id_seq'::regclass);


--
-- Name: event_attachments id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.event_attachments ALTER COLUMN id SET DEFAULT nextval('public.event_attachments_id_seq'::regclass);


--
-- Name: event_attributes id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.event_attributes ALTER COLUMN id SET DEFAULT nextval('public.event_attributes_id_seq'::regclass);


--
-- Name: event_ext_ipaddresses id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.event_ext_ipaddresses ALTER COLUMN id SET DEFAULT nextval('public.event_ext_ipaddresses_id_seq'::regclass);


--
-- Name: event_home_ipaddresses id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.event_home_ipaddresses ALTER COLUMN id SET DEFAULT nextval('public.event_home_ipaddresses_id_seq'::regclass);


--
-- Name: event_rules id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.event_rules ALTER COLUMN id SET DEFAULT nextval('public.event_rules_id_seq'::regclass);


--
-- Name: event_ttps id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.event_ttps ALTER COLUMN id SET DEFAULT nextval('public.event_ttps_id_seq'::regclass);


--
-- Name: events id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.events ALTER COLUMN id SET DEFAULT nextval('public.events_id_seq'::regclass);


--
-- Name: grids_states id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.grids_states ALTER COLUMN id SET DEFAULT nextval('public.assets_crud_grids_id_seq'::regclass);


--
-- Name: iam_orgs id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.iam_orgs ALTER COLUMN id SET DEFAULT nextval('public.iam_orgs_id_seq'::regclass);


--
-- Name: iam_permissions id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.iam_permissions ALTER COLUMN id SET DEFAULT nextval('public.iam_permissions_id_seq'::regclass);


--
-- Name: iam_roles id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.iam_roles ALTER COLUMN id SET DEFAULT nextval('public.iam_roles_id_seq'::regclass);


--
-- Name: iam_users id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.iam_users ALTER COLUMN id SET DEFAULT nextval('public.iam_users_id_seq'::regclass);


--
-- Name: iocs id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.iocs ALTER COLUMN id SET DEFAULT nextval('public.socbox_detection_iocs_id_seq'::regclass);


--
-- Name: iocs_sources id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.iocs_sources ALTER COLUMN id SET DEFAULT nextval('public.socbox_detection_iocs_sources_id_seq'::regclass);


--
-- Name: iocs_sources_statuses id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.iocs_sources_statuses ALTER COLUMN id SET DEFAULT nextval('public.socbox_detection_iocs_sources_statuses_id_seq'::regclass);


--
-- Name: iocs_types id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.iocs_types ALTER COLUMN id SET DEFAULT nextval('public.socbox_detection_iocs_types_id_seq'::regclass);


--
-- Name: iocs_whitelist id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.iocs_whitelist ALTER COLUMN id SET DEFAULT nextval('public.socbox_detection_iocs_whitelist_id_seq'::regclass);


--
-- Name: objects_networks id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.objects_networks ALTER COLUMN id SET DEFAULT nextval('public.objects_networks_id_seq'::regclass);


--
-- Name: orgs id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.orgs ALTER COLUMN id SET DEFAULT nextval('public.orgs_id_seq'::regclass);


--
-- Name: persons id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.persons ALTER COLUMN id SET DEFAULT nextval('public.persons_id_seq'::regclass);


--
-- Name: siem_external_systems id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.siem_external_systems ALTER COLUMN id SET DEFAULT nextval('public.siem_external_systems_id_seq'::regclass);


--
-- Name: siem_filter_types id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.siem_filter_types ALTER COLUMN id SET DEFAULT nextval('public.siem_filter_types_id_seq'::regclass);


--
-- Name: siem_filters id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.siem_filters ALTER COLUMN id SET DEFAULT nextval('public.siem_filters_id_seq'::regclass);


--
-- Name: siem_task_statuses id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.siem_task_statuses ALTER COLUMN id SET DEFAULT nextval('public.siem_task_statuses_id_seq'::regclass);


--
-- Name: siem_task_types id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.siem_task_types ALTER COLUMN id SET DEFAULT nextval('public.siem_task_types_id_seq'::regclass);


--
-- Name: siem_tasks id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.siem_tasks ALTER COLUMN id SET DEFAULT nextval('public.siem_tasks_id_seq'::regclass);


--
-- Name: signatures_overlay id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.signatures_overlay ALTER COLUMN id SET DEFAULT nextval('public.signatures_overlay_id_seq'::regclass);


--
-- Name: signatures_test id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.signatures_test ALTER COLUMN id SET DEFAULT nextval('public.signatures_test_id_seq'::regclass);


--
-- Name: socbox_actors_objects id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_actors_objects ALTER COLUMN id SET DEFAULT nextval('public.socbox_actors_objects_id_seq'::regclass);


--
-- Name: socbox_actors_persons_orders id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_actors_persons_orders ALTER COLUMN id SET DEFAULT nextval('public.socbox_actors_persons_orders_id_seq'::regclass);


--
-- Name: socbox_actors_scopes id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_actors_scopes ALTER COLUMN id SET DEFAULT nextval('public.socbox_actors_scopes_id_seq'::regclass);


--
-- Name: socbox_assets_homenets id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_assets_homenets ALTER COLUMN id SET DEFAULT nextval('public.socbox_assets_homenets_seq'::regclass);


--
-- Name: socbox_assets_resources id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_assets_resources ALTER COLUMN id SET DEFAULT nextval('public.objects_resources_id_seq'::regclass);


--
-- Name: socbox_assets_resources_types id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_assets_resources_types ALTER COLUMN id SET DEFAULT nextval('public.objects_resource_types_id_seq'::regclass);


--
-- Name: socbox_content_type id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_content_type ALTER COLUMN id SET DEFAULT nextval('public.socbox_content_type_id_seq'::regclass);


--
-- Name: socbox_core_content_type id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_core_content_type ALTER COLUMN id SET DEFAULT nextval('public.socbox_core_content_type_id_seq'::regclass);


--
-- Name: socbox_detection_classtypes id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_detection_classtypes ALTER COLUMN id SET DEFAULT nextval('public.socbox_detection_classtypes_id_seq'::regclass);


--
-- Name: socbox_detection_local_classtypes id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_detection_local_classtypes ALTER COLUMN id SET DEFAULT nextval('public.socbox_detection_local_classtypes_id_seq'::regclass);


--
-- Name: socbox_detection_mitre_tactics id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_detection_mitre_tactics ALTER COLUMN id SET DEFAULT nextval('public.socbox_detection_mitre_tactics_id_seq'::regclass);


--
-- Name: socbox_detection_mitre_techniques id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_detection_mitre_techniques ALTER COLUMN id SET DEFAULT nextval('public.socbox_detection_mitre_techniques_id_seq'::regclass);


--
-- Name: socbox_detection_nccci_classes id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_detection_nccci_classes ALTER COLUMN id SET DEFAULT nextval('public.socbox_detection_nccci_classes_id_seq'::regclass);


--
-- Name: socbox_detection_sensors_groups id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_detection_sensors_groups ALTER COLUMN id SET DEFAULT nextval('public.socbox_detection_sensors_groups_id_seq'::regclass);


--
-- Name: socbox_detection_sensors_hardwares id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_detection_sensors_hardwares ALTER COLUMN id SET DEFAULT nextval('public.socbox_detection_sensors_hardwares_id_seq'::regclass);


--
-- Name: socbox_detection_signatures_classtypes id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_detection_signatures_classtypes ALTER COLUMN id SET DEFAULT nextval('public.socbox_detection_signatures_classtypes_id_seq'::regclass);


--
-- Name: socbox_extras_changelog id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_extras_changelog ALTER COLUMN id SET DEFAULT nextval('public.socbox_extras_changelog_id_seq'::regclass);


--
-- Name: socbox_extras_custom_fields id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_extras_custom_fields ALTER COLUMN id SET DEFAULT nextval('public.socbox_extras_custom_fields_id_seq'::regclass);


--
-- Name: socbox_extras_custom_fields_choicesets id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_extras_custom_fields_choicesets ALTER COLUMN id SET DEFAULT nextval('public.socbox_extras_custom_fields_choicesets_id_seq'::regclass);


--
-- Name: socbox_extras_custom_fields_objects_types id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_extras_custom_fields_objects_types ALTER COLUMN id SET DEFAULT nextval('public.socbox_extras_custom_fields_objects_types_id_seq'::regclass);


--
-- Name: socbox_extras_menu id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_extras_menu ALTER COLUMN id SET DEFAULT nextval('public.socbox_extras_menu_id_seq'::regclass);


--
-- Name: socbox_extras_tag_object_types id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_extras_tag_object_types ALTER COLUMN id SET DEFAULT nextval('public.socbox_extras_tag_object_types_id_seq'::regclass);


--
-- Name: socbox_extras_tagged_item id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_extras_tagged_item ALTER COLUMN id SET DEFAULT nextval('public.socbox_extras_tagged_item_id_seq'::regclass);


--
-- Name: socbox_extras_tags id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_extras_tags ALTER COLUMN id SET DEFAULT nextval('public.socbox_extras_tags_id_seq'::regclass);


--
-- Name: socbox_geo_cities id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_geo_cities ALTER COLUMN id SET DEFAULT nextval('public.socbox_geo_cities_id_seq'::regclass);


--
-- Name: socbox_geo_countries id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_geo_countries ALTER COLUMN id SET DEFAULT nextval('public.socbox_geo_countries_id_seq'::regclass);


--
-- Name: socbox_geo_districts id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_geo_districts ALTER COLUMN id SET DEFAULT nextval('public.socbox_geo_districts_id_seq'::regclass);


--
-- Name: socbox_geo_regions id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_geo_regions ALTER COLUMN id SET DEFAULT nextval('public.socbox_geo_regions_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Name: users_actions id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users_actions ALTER COLUMN id SET DEFAULT nextval('public.users_actions_id_seq'::regclass);


--
-- Name: users_histories id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users_histories ALTER COLUMN id SET DEFAULT nextval('public.users_histories_id_seq'::regclass);


--
-- Name: users_metrics id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users_metrics ALTER COLUMN id SET DEFAULT nextval('public.users_metrics_id_seq'::regclass);


--
-- Name: users_roles id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users_roles ALTER COLUMN id SET DEFAULT nextval('public.users_roles_id_seq'::regclass);


--
-- Name: socbox_core_content_type PK_001c982499ea7b72379090b7c8a; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_core_content_type
    ADD CONSTRAINT "PK_001c982499ea7b72379090b7c8a" PRIMARY KEY (id);


--
-- Name: socbox_detection_sensors_hardwares PK_0179a43e84e23619f0b8dd4c92c; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_detection_sensors_hardwares
    ADD CONSTRAINT "PK_0179a43e84e23619f0b8dd4c92c" PRIMARY KEY (id);


--
-- Name: iam_users PK_02086c69f80fed8ae319ec498ec; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.iam_users
    ADD CONSTRAINT "PK_02086c69f80fed8ae319ec498ec" PRIMARY KEY (id);


--
-- Name: socbox_extras_custom_fields_choicesets PK_0a92876f6fe5eb96b4223519bcf; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_extras_custom_fields_choicesets
    ADD CONSTRAINT "PK_0a92876f6fe5eb96b4223519bcf" PRIMARY KEY (id);


--
-- Name: socbox_detection_attacks_priorities PK_13fb2df0863ce19910ace8a2e8a; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_detection_attacks_priorities
    ADD CONSTRAINT "PK_13fb2df0863ce19910ace8a2e8a" PRIMARY KEY (id);


--
-- Name: socbox_extras_changelog PK_1b84da1435787c5ab74dd65bf6e; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_extras_changelog
    ADD CONSTRAINT "PK_1b84da1435787c5ab74dd65bf6e" PRIMARY KEY (id);


--
-- Name: socbox_detection_local_classtypes PK_1dd705ac2a2bdb3406edd0abfc1; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_detection_local_classtypes
    ADD CONSTRAINT "PK_1dd705ac2a2bdb3406edd0abfc1" PRIMARY KEY (id);


--
-- Name: socbox_detection_directions PK_1f7f0c959a8e7cba74a49b66558; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_detection_directions
    ADD CONSTRAINT "PK_1f7f0c959a8e7cba74a49b66558" PRIMARY KEY (id);


--
-- Name: socbox_detection_signatures_classtypes PK_237250286fff972a61b3f17168a; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_detection_signatures_classtypes
    ADD CONSTRAINT "PK_237250286fff972a61b3f17168a" PRIMARY KEY (id);


--
-- Name: socbox_extras_tagged_item PK_27ee9cc446fdbc3e87c1d9b5012; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_extras_tagged_item
    ADD CONSTRAINT "PK_27ee9cc446fdbc3e87c1d9b5012" PRIMARY KEY (id);


--
-- Name: socbox_detection_sensors_groups PK_29d365db7e7d2ceaea2a888c249; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_detection_sensors_groups
    ADD CONSTRAINT "PK_29d365db7e7d2ceaea2a888c249" PRIMARY KEY (id);


--
-- Name: iam_orgs PK_31a8aee85ab6636e04864853dac; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.iam_orgs
    ADD CONSTRAINT "PK_31a8aee85ab6636e04864853dac" PRIMARY KEY (id);


--
-- Name: socbox_content_type PK_3d7d173856cead46b551aac18c0; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_content_type
    ADD CONSTRAINT "PK_3d7d173856cead46b551aac18c0" PRIMARY KEY (id);


--
-- Name: iocs_types PK_4ad15ad7a5a5ff822f32af8f5bd; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.iocs_types
    ADD CONSTRAINT "PK_4ad15ad7a5a5ff822f32af8f5bd" PRIMARY KEY (id);


--
-- Name: socbox_actors_scopes PK_5aee33fe6e1f9e8b834104140a9; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_actors_scopes
    ADD CONSTRAINT "PK_5aee33fe6e1f9e8b834104140a9" PRIMARY KEY (id);


--
-- Name: socbox_detection_nccci_classes PK_5b5a98d0172dca94fc89852241d; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_detection_nccci_classes
    ADD CONSTRAINT "PK_5b5a98d0172dca94fc89852241d" PRIMARY KEY (id);


--
-- Name: socbox_detection_recomendations PK_6203ea57bef3b948911bcdbf596; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_detection_recomendations
    ADD CONSTRAINT "PK_6203ea57bef3b948911bcdbf596" PRIMARY KEY (id);


--
-- Name: iam_users_roles PK_62d4f96c63153de1812ad362469; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.iam_users_roles
    ADD CONSTRAINT "PK_62d4f96c63153de1812ad362469" PRIMARY KEY (user_id, role_id);


--
-- Name: iam_permissions PK_647c72677d99c172d9ed329f39c; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.iam_permissions
    ADD CONSTRAINT "PK_647c72677d99c172d9ed329f39c" PRIMARY KEY (id);


--
-- Name: socbox_extras_tag_object_types PK_6c37d679af351dac2781864cc18; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_extras_tag_object_types
    ADD CONSTRAINT "PK_6c37d679af351dac2781864cc18" PRIMARY KEY (id);


--
-- Name: iocs_whitelist PK_763562b0f4db5685a6022a200bf; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.iocs_whitelist
    ADD CONSTRAINT "PK_763562b0f4db5685a6022a200bf" PRIMARY KEY (id);


--
-- Name: iocs_sources PK_76af7e246d194726aaee4887277; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.iocs_sources
    ADD CONSTRAINT "PK_76af7e246d194726aaee4887277" PRIMARY KEY (id);


--
-- Name: socbox_extras_custom_fields_objects_types PK_7c23440207b0809f9c89c752d3e; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_extras_custom_fields_objects_types
    ADD CONSTRAINT "PK_7c23440207b0809f9c89c752d3e" PRIMARY KEY (id);


--
-- Name: socbox_extras_custom_fields PK_82af90c219a9d327bbca5ca5e5c; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_extras_custom_fields
    ADD CONSTRAINT "PK_82af90c219a9d327bbca5ca5e5c" PRIMARY KEY (id);


--
-- Name: socbox_detection_signatures PK_96583e7d37c0f45879470683569; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_detection_signatures
    ADD CONSTRAINT "PK_96583e7d37c0f45879470683569" PRIMARY KEY (sid, gid, rev);


--
-- Name: socbox_detection_classtypes PK_979f6dcbae4537549aedfbaa3bd; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_detection_classtypes
    ADD CONSTRAINT "PK_979f6dcbae4537549aedfbaa3bd" PRIMARY KEY (id);


--
-- Name: socbox_actors_objects_sensors PK_a0b6d0728f7ad1d1a7e3c5e5c4f; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_actors_objects_sensors
    ADD CONSTRAINT "PK_a0b6d0728f7ad1d1a7e3c5e5c4f" PRIMARY KEY (sensor_id, object_id);


--
-- Name: iam_roles PK_aa79b0099c20eed09191e9d4159; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.iam_roles
    ADD CONSTRAINT "PK_aa79b0099c20eed09191e9d4159" PRIMARY KEY (id);


--
-- Name: socbox_detection_sensors PK_b8bd5fcfd700e39e96bcd9ba6b7; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_detection_sensors
    ADD CONSTRAINT "PK_b8bd5fcfd700e39e96bcd9ba6b7" PRIMARY KEY (id);


--
-- Name: socbox_actors_persons_orders PK_b9011ffe622250ca8970b65631e; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_actors_persons_orders
    ADD CONSTRAINT "PK_b9011ffe622250ca8970b65631e" PRIMARY KEY (id);


--
-- Name: socbox_extras_tags PK_bf5d32dd3e916c1c91046996e98; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_extras_tags
    ADD CONSTRAINT "PK_bf5d32dd3e916c1c91046996e98" PRIMARY KEY (id);


--
-- Name: socbox_extras_menu PK_d5de65b7472e4f057da8615e364; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_extras_menu
    ADD CONSTRAINT "PK_d5de65b7472e4f057da8615e364" PRIMARY KEY (id);


--
-- Name: iam_roles_permissions PK_eecaa181a1b6be67a39b3cf0138; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.iam_roles_permissions
    ADD CONSTRAINT "PK_eecaa181a1b6be67a39b3cf0138" PRIMARY KEY (role_id, permission_id);


--
-- Name: iocs PK_efac1d43a82ae1513d4de9f434c; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.iocs
    ADD CONSTRAINT "PK_efac1d43a82ae1513d4de9f434c" PRIMARY KEY (id);


--
-- Name: socbox_geo_countries PK_f50279873cbdb6d0f7739f645ab; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_geo_countries
    ADD CONSTRAINT "PK_f50279873cbdb6d0f7739f645ab" PRIMARY KEY (id);


--
-- Name: socbox_actors_persons_orders UQ_037982de1e6a690a6afc309d1a3; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_actors_persons_orders
    ADD CONSTRAINT "UQ_037982de1e6a690a6afc309d1a3" UNIQUE (name);


--
-- Name: socbox_extras_custom_fields UQ_0b5cf9c9098191d772885232e48; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_extras_custom_fields
    ADD CONSTRAINT "UQ_0b5cf9c9098191d772885232e48" UNIQUE (name);


--
-- Name: socbox_detection_local_classtypes UQ_1f7fc16c23c17586f471c6a96ab; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_detection_local_classtypes
    ADD CONSTRAINT "UQ_1f7fc16c23c17586f471c6a96ab" UNIQUE (name);


--
-- Name: socbox_detection_nccci_classes UQ_264e71dc806964928a82dff1eb2; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_detection_nccci_classes
    ADD CONSTRAINT "UQ_264e71dc806964928a82dff1eb2" UNIQUE (name);


--
-- Name: socbox_actors_scopes UQ_42c4014dc9321d9be83a5dc9993; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_actors_scopes
    ADD CONSTRAINT "UQ_42c4014dc9321d9be83a5dc9993" UNIQUE (name);


--
-- Name: socbox_detection_mitre_techniques UQ_4458ef85bd35cb2602637f2e2f0; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_detection_mitre_techniques
    ADD CONSTRAINT "UQ_4458ef85bd35cb2602637f2e2f0" UNIQUE (name);


--
-- Name: socbox_actors_subjects UQ_47a287fe64bd0e1027e603c335c; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_actors_subjects
    ADD CONSTRAINT "UQ_47a287fe64bd0e1027e603c335c" UNIQUE (name);


--
-- Name: iocs_whitelist UQ_4b1c4d4bdfe79281686f2e7d768; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.iocs_whitelist
    ADD CONSTRAINT "UQ_4b1c4d4bdfe79281686f2e7d768" UNIQUE (value);


--
-- Name: socbox_geo_countries UQ_5b2432df180a340a8392b58f9bf; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_geo_countries
    ADD CONSTRAINT "UQ_5b2432df180a340a8392b58f9bf" UNIQUE (alpha2);


--
-- Name: socbox_detection_sensors_hardwares UQ_605d1ca9913b3b08b842e54edbe; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_detection_sensors_hardwares
    ADD CONSTRAINT "UQ_605d1ca9913b3b08b842e54edbe" UNIQUE (name);


--
-- Name: socbox_detection_classtypes UQ_60c372b82180897f2fe3deb0592; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_detection_classtypes
    ADD CONSTRAINT "UQ_60c372b82180897f2fe3deb0592" UNIQUE (name);


--
-- Name: iam_orgs UQ_635361b445187a0f3f9f4354c5d; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.iam_orgs
    ADD CONSTRAINT "UQ_635361b445187a0f3f9f4354c5d" UNIQUE (name);


--
-- Name: iam_users UQ_6cd7bf7792bf699075afff5d3b0; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.iam_users
    ADD CONSTRAINT "UQ_6cd7bf7792bf699075afff5d3b0" UNIQUE (login);


--
-- Name: iocs_types UQ_6e1b3c2f06afe7ed0cd78c8685e; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.iocs_types
    ADD CONSTRAINT "UQ_6e1b3c2f06afe7ed0cd78c8685e" UNIQUE (name);


--
-- Name: socbox_detection_sensors_groups UQ_775496deecba05f29cd2a066eec; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_detection_sensors_groups
    ADD CONSTRAINT "UQ_775496deecba05f29cd2a066eec" UNIQUE (name);


--
-- Name: socbox_extras_custom_fields_choicesets UQ_911fc537bcf7be824b8485fb50f; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_extras_custom_fields_choicesets
    ADD CONSTRAINT "UQ_911fc537bcf7be824b8485fb50f" UNIQUE (name);


--
-- Name: socbox_extras_menu UQ_9b254c0070b47b8381cf0f3de9f; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_extras_menu
    ADD CONSTRAINT "UQ_9b254c0070b47b8381cf0f3de9f" UNIQUE (name);


--
-- Name: socbox_actors_objects UQ_a7ac54147928203fae0cb502e17; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_actors_objects
    ADD CONSTRAINT "UQ_a7ac54147928203fae0cb502e17" UNIQUE (name);


--
-- Name: socbox_detection_signatures_classtypes UQ_b2557634f1089eee187a90993d6; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_detection_signatures_classtypes
    ADD CONSTRAINT "UQ_b2557634f1089eee187a90993d6" UNIQUE (name);


--
-- Name: socbox_detection_recomendations UQ_bc0384d75368110048c3e98cfa2; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_detection_recomendations
    ADD CONSTRAINT "UQ_bc0384d75368110048c3e98cfa2" UNIQUE (name);


--
-- Name: socbox_detection_mitre_tactics UQ_dbfba85a8fd6b9b8e8d7e68b342; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_detection_mitre_tactics
    ADD CONSTRAINT "UQ_dbfba85a8fd6b9b8e8d7e68b342" UNIQUE (name);


--
-- Name: socbox_detection_sensors UQ_e35e9a861b337546d6a7f749fc5; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_detection_sensors
    ADD CONSTRAINT "UQ_e35e9a861b337546d6a7f749fc5" UNIQUE (name);


--
-- Name: socbox_geo_cities UQ_e76735a2df755e5608b592d0c0a; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_geo_cities
    ADD CONSTRAINT "UQ_e76735a2df755e5608b592d0c0a" UNIQUE (name, region_id);


--
-- Name: iocs_sources UQ_efc30f100024d2838f402deb127; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.iocs_sources
    ADD CONSTRAINT "UQ_efc30f100024d2838f402deb127" UNIQUE (name);


--
-- Name: iam_roles UQ_f0201a16c79111ad493ba62ae0f; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.iam_roles
    ADD CONSTRAINT "UQ_f0201a16c79111ad493ba62ae0f" UNIQUE (name);


--
-- Name: grids_states assets_crud_grids_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.grids_states
    ADD CONSTRAINT assets_crud_grids_pkey PRIMARY KEY (id);


--
-- Name: dict_attack_classes dict_attack_classes_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.dict_attack_classes
    ADD CONSTRAINT dict_attack_classes_pkey PRIMARY KEY (id);


--
-- Name: dict_attack_levels dict_attack_levels_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.dict_attack_levels
    ADD CONSTRAINT dict_attack_levels_pkey PRIMARY KEY (id);


--
-- Name: dict_attack_methods dict_attack_methods_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.dict_attack_methods
    ADD CONSTRAINT dict_attack_methods_pkey PRIMARY KEY (id);


--
-- Name: dict_attack_protocols dict_attack_protocols_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.dict_attack_protocols
    ADD CONSTRAINT dict_attack_protocols_pkey PRIMARY KEY (id);


--
-- Name: socbox_geo_cities dict_cities_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_geo_cities
    ADD CONSTRAINT dict_cities_pkey PRIMARY KEY (id);


--
-- Name: socbox_geo_districts dict_districts_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_geo_districts
    ADD CONSTRAINT dict_districts_pkey PRIMARY KEY (id);


--
-- Name: event_sources dict_event_sources_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.event_sources
    ADD CONSTRAINT dict_event_sources_pkey PRIMARY KEY (id);


--
-- Name: dict_event_statuses dict_event_statuses_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.dict_event_statuses
    ADD CONSTRAINT dict_event_statuses_pkey PRIMARY KEY (id);


--
-- Name: dict_event_types dict_event_types_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.dict_event_types
    ADD CONSTRAINT dict_event_types_pkey PRIMARY KEY (id);


--
-- Name: dict_misp_attribute_categories dict_misp_attribute_categories_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.dict_misp_attribute_categories
    ADD CONSTRAINT dict_misp_attribute_categories_pkey PRIMARY KEY (id);


--
-- Name: dict_misp_attribute_types dict_misp_attribute_types_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.dict_misp_attribute_types
    ADD CONSTRAINT dict_misp_attribute_types_pkey PRIMARY KEY (id);


--
-- Name: socbox_detection_mitre_tactics dict_mitre_tactics_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_detection_mitre_tactics
    ADD CONSTRAINT dict_mitre_tactics_pkey PRIMARY KEY (id);


--
-- Name: socbox_detection_mitre_techniques dict_mitre_techniques_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_detection_mitre_techniques
    ADD CONSTRAINT dict_mitre_techniques_pkey PRIMARY KEY (id);


--
-- Name: dict_person_modes dict_modes_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.dict_person_modes
    ADD CONSTRAINT dict_modes_pkey PRIMARY KEY (id);


--
-- Name: dict_person_orders dict_orders_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.dict_person_orders
    ADD CONSTRAINT dict_orders_pkey PRIMARY KEY (id);


--
-- Name: socbox_geo_regions dict_regions_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_geo_regions
    ADD CONSTRAINT dict_regions_pkey PRIMARY KEY (id);


--
-- Name: dict_person_roles dict_roles_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.dict_person_roles
    ADD CONSTRAINT dict_roles_pkey PRIMARY KEY (id);


--
-- Name: dict_soft_types dict_soft_types_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.dict_soft_types
    ADD CONSTRAINT dict_soft_types_pkey PRIMARY KEY (id);


--
-- Name: dict_softs dict_softs_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.dict_softs
    ADD CONSTRAINT dict_softs_pkey PRIMARY KEY (id);


--
-- Name: dict_vulnerabilities dict_vulnerabilities_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.dict_vulnerabilities
    ADD CONSTRAINT dict_vulnerabilities_pkey PRIMARY KEY (id);


--
-- Name: event_attachments event_attachments_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.event_attachments
    ADD CONSTRAINT event_attachments_pkey PRIMARY KEY (id);


--
-- Name: event_attributes event_attributes_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.event_attributes
    ADD CONSTRAINT event_attributes_pkey PRIMARY KEY (id);


--
-- Name: event_ext_ipaddresses event_ext_ipaddresses_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.event_ext_ipaddresses
    ADD CONSTRAINT event_ext_ipaddresses_pkey PRIMARY KEY (id);


--
-- Name: event_home_ipaddresses event_home_ipaddresses_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.event_home_ipaddresses
    ADD CONSTRAINT event_home_ipaddresses_pkey PRIMARY KEY (id);


--
-- Name: event_rules event_rules_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.event_rules
    ADD CONSTRAINT event_rules_pkey PRIMARY KEY (id);


--
-- Name: event_ttps event_ttps_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.event_ttps
    ADD CONSTRAINT event_ttps_pkey PRIMARY KEY (id);


--
-- Name: events events_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.events
    ADD CONSTRAINT events_pkey PRIMARY KEY (id);


--
-- Name: geoip geoip_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.geoip
    ADD CONSTRAINT geoip_pkey PRIMARY KEY (inetnum);


--
-- Name: socbox_assets_homenets objects_homenets_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_assets_homenets
    ADD CONSTRAINT objects_homenets_pkey PRIMARY KEY (id);


--
-- Name: objects_networks objects_networks_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.objects_networks
    ADD CONSTRAINT objects_networks_pkey PRIMARY KEY (id);


--
-- Name: objects_persons objects_persons_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.objects_persons
    ADD CONSTRAINT objects_persons_pkey PRIMARY KEY (object_id, person_id);


--
-- Name: socbox_actors_objects objects_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_actors_objects
    ADD CONSTRAINT objects_pkey PRIMARY KEY (id);


--
-- Name: socbox_assets_resources_types objects_resource_types_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_assets_resources_types
    ADD CONSTRAINT objects_resource_types_pkey PRIMARY KEY (id);


--
-- Name: socbox_assets_resources objects_resources_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_assets_resources
    ADD CONSTRAINT objects_resources_pkey PRIMARY KEY (id);


--
-- Name: orgs orgs_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.orgs
    ADD CONSTRAINT orgs_pkey PRIMARY KEY (id);


--
-- Name: orgs orgs_unique_name; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.orgs
    ADD CONSTRAINT orgs_unique_name UNIQUE (name);


--
-- Name: persons persons_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.persons
    ADD CONSTRAINT persons_pkey PRIMARY KEY (id);

--
-- Name: sensors_settings sensors_settings_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.sensors_settings
    ADD CONSTRAINT sensors_settings_pkey PRIMARY KEY (id);


--
-- Name: siem_actions siem_actions_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.siem_actions
    ADD CONSTRAINT siem_actions_pkey PRIMARY KEY (id);


--
-- Name: siem_errors siem_errors_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.siem_errors
    ADD CONSTRAINT siem_errors_pkey PRIMARY KEY (id);


--
-- Name: siem_external_systems siem_external_systems_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.siem_external_systems
    ADD CONSTRAINT siem_external_systems_pkey PRIMARY KEY (id);


--
-- Name: siem_filter_types siem_filter_types_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.siem_filter_types
    ADD CONSTRAINT siem_filter_types_pkey PRIMARY KEY (id);


--
-- Name: siem_filters siem_filters_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.siem_filters
    ADD CONSTRAINT siem_filters_pkey PRIMARY KEY (id);


--
-- Name: siem_task_statuses siem_task_statuses_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.siem_task_statuses
    ADD CONSTRAINT siem_task_statuses_pkey PRIMARY KEY (id);


--
-- Name: siem_task_types siem_task_types_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.siem_task_types
    ADD CONSTRAINT siem_task_types_pkey PRIMARY KEY (id);


--
-- Name: siem_tasks siem_tasks_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.siem_tasks
    ADD CONSTRAINT siem_tasks_pkey PRIMARY KEY (id);


--
-- Name: siem_views siem_views_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.siem_views
    ADD CONSTRAINT siem_views_pkey PRIMARY KEY (id);


--
-- Name: signatures_overlay signatures_overlay_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.signatures_overlay
    ADD CONSTRAINT signatures_overlay_pkey PRIMARY KEY (id);


--
-- Name: iocs_sources_statuses socbox_detection_iocs_sources_statuses_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.iocs_sources_statuses
    ADD CONSTRAINT socbox_detection_iocs_sources_statuses_pkey PRIMARY KEY (id);


--
-- Name: statistics statistics_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.statistics
    ADD CONSTRAINT statistics_pkey PRIMARY KEY (sid);


--
-- Name: socbox_actors_subjects subjects_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_actors_subjects
    ADD CONSTRAINT subjects_pkey PRIMARY KEY (id);


--
-- Name: socbox_detection_threat_levels threat_levels_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_detection_threat_levels
    ADD CONSTRAINT threat_levels_pkey PRIMARY KEY (id);


--
-- Name: users_actions users_actions_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users_actions
    ADD CONSTRAINT users_actions_pkey PRIMARY KEY (id);


--
-- Name: users_histories users_histories_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users_histories
    ADD CONSTRAINT users_histories_pkey PRIMARY KEY (id);


--
-- Name: users_metrics users_metrics_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users_metrics
    ADD CONSTRAINT users_metrics_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: users_roles users_roles_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users_roles
    ADD CONSTRAINT users_roles_pkey PRIMARY KEY (id);


--
-- Name: users users_unique_login; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_unique_login UNIQUE (login);


--
-- Name: IDX_0e95f4b15aefe7b621f5ca8e73; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX "IDX_0e95f4b15aefe7b621f5ca8e73" ON public.socbox_actors_objects_sensors USING btree (object_id);


--
-- Name: IDX_29c1e511b10f71a409c0942351; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX "IDX_29c1e511b10f71a409c0942351" ON public.socbox_extras_tagged_item USING btree (object_type_id, object_id);


--
-- Name: IDX_3928031ee22160235ae1479ddc; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX "IDX_3928031ee22160235ae1479ddc" ON public.iam_users_roles USING btree (user_id);


--
-- Name: IDX_4930739bbc6ad5ebf0593bcc01; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX "IDX_4930739bbc6ad5ebf0593bcc01" ON public.iam_roles_permissions USING btree (role_id);


--
-- Name: IDX_7dd37c6a4ab567b4050e543909; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX "IDX_7dd37c6a4ab567b4050e543909" ON public.socbox_content_type USING btree (app_label, model);


--
-- Name: IDX_a50c16349273b8189a861953b7; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX "IDX_a50c16349273b8189a861953b7" ON public.socbox_extras_tag_object_types USING btree (tag_id, object_type_id);


--
-- Name: IDX_a5422c181d5e355977b60b13e7; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX "IDX_a5422c181d5e355977b60b13e7" ON public.socbox_actors_objects_sensors USING btree (sensor_id);


--
-- Name: IDX_ae8a14ee577e52a38e0cb77757; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX "IDX_ae8a14ee577e52a38e0cb77757" ON public.socbox_core_content_type USING btree (app_label, model);


--
-- Name: IDX_b1a501895c3376b9a40cbda500; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX "IDX_b1a501895c3376b9a40cbda500" ON public.iam_users_roles USING btree (role_id);


--
-- Name: IDX_e22da5effbda40cac2bcabdded; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX "IDX_e22da5effbda40cac2bcabdded" ON public.socbox_extras_tags USING btree (name);


--
-- Name: IDX_fb13ed1d19e4a9af783c07563a; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX "IDX_fb13ed1d19e4a9af783c07563a" ON public.iam_roles_permissions USING btree (permission_id);


--
-- Name: UQ_iocs_value_type_source; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX "UQ_iocs_value_type_source" ON public.iocs USING btree (value, ioc_type_id, ioc_source_id);


--
-- Name: assets_grids_states_grid_type_idx; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX assets_grids_states_grid_type_idx ON public.grids_states USING btree (grid_type, model_name);


--
-- Name: idx_classtype; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_classtype ON public.socbox_detection_signatures USING btree (classtype);


--
-- Name: idx_content_hash; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_content_hash ON public.socbox_detection_signatures USING btree (content_hash);


--
-- Name: idx_is_active; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_is_active ON public.socbox_detection_signatures USING btree (is_active);


--
-- Name: idx_metadata; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_metadata ON public.socbox_detection_signatures USING btree (metadata);


--
-- Name: idx_priority; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_priority ON public.socbox_detection_signatures USING btree (priority);


--
-- Name: idx_source_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_source_id ON public.socbox_detection_signatures USING btree (source_id);


--
-- Name: idx_source_statuses_created; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_source_statuses_created ON public.iocs_sources_statuses USING btree (created_at DESC);


--
-- Name: idx_source_statuses_source_time; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_source_statuses_source_time ON public.iocs_sources_statuses USING btree (source_id, start_time DESC);


--
-- Name: idx_source_statuses_success; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_source_statuses_success ON public.iocs_sources_statuses USING btree (success, start_time DESC);


--
-- Name: objects_homenets_network_idx; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX objects_homenets_network_idx ON public.socbox_assets_homenets USING btree (network, object_id);


--
-- Name: signatures_test_i_sid_idx; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX signatures_test_i_sid_idx ON public.signatures_test USING btree (i_sid, i_rev, i_gid);


--
-- Name: socbox_extras_customfield_choice_set_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX socbox_extras_customfield_choice_set_id ON public.socbox_extras_custom_fields USING btree (choiceset_id);


--
-- Name: socbox_extras_customfield_content_customfield_id_contentty_uniq; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX socbox_extras_customfield_content_customfield_id_contentty_uniq ON public.socbox_extras_custom_fields_objects_types USING btree (custom_field_id, object_type_id);


--
-- Name: socbox_extras_customfield_content_types_content_type_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX socbox_extras_customfield_content_types_content_type_id ON public.socbox_extras_custom_fields_objects_types USING btree (object_type_id);


--
-- Name: socbox_extras_customfield_content_types_customfield_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX socbox_extras_customfield_content_types_customfield_id ON public.socbox_extras_custom_fields_objects_types USING btree (custom_field_id);


--
-- Name: socbox_extras_customfield_name; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX socbox_extras_customfield_name ON public.socbox_extras_custom_fields USING btree (name);


--
-- Name: socbox_extras_customfield_object_type_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX socbox_extras_customfield_object_type_id ON public.socbox_extras_custom_fields USING btree (related_object_type_id);


--
-- Name: socbox_actors_objects FK_0387693a98d574367bfd4dbf872; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_actors_objects
    ADD CONSTRAINT "FK_0387693a98d574367bfd4dbf872" FOREIGN KEY (city_id) REFERENCES public.socbox_geo_cities(id);


--
-- Name: iocs FK_0e33ef3f03c6b74a5958b5e547b; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.iocs
    ADD CONSTRAINT "FK_0e33ef3f03c6b74a5958b5e547b" FOREIGN KEY (ioc_source_id) REFERENCES public.iocs_sources(id);


--
-- Name: socbox_actors_objects_sensors FK_0e95f4b15aefe7b621f5ca8e73a; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_actors_objects_sensors
    ADD CONSTRAINT "FK_0e95f4b15aefe7b621f5ca8e73a" FOREIGN KEY (object_id) REFERENCES public.socbox_actors_objects(id);


--
-- Name: socbox_actors_objects FK_120a4181c387c35cf4175f753f2; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_actors_objects
    ADD CONSTRAINT "FK_120a4181c387c35cf4175f753f2" FOREIGN KEY (subject_id) REFERENCES public.socbox_actors_subjects(id);


--
-- Name: iam_users FK_24e1c02589bf70287daf0a678de; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.iam_users
    ADD CONSTRAINT "FK_24e1c02589bf70287daf0a678de" FOREIGN KEY (org_id) REFERENCES public.iam_orgs(id);


--
-- Name: socbox_geo_cities FK_260091947a0bdd88a4a5d68f500; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_geo_cities
    ADD CONSTRAINT "FK_260091947a0bdd88a4a5d68f500" FOREIGN KEY (region_id) REFERENCES public.socbox_geo_regions(id);


--
-- Name: socbox_detection_sensors FK_36a77cbe7ca9c7dd04ef4889634; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_detection_sensors
    ADD CONSTRAINT "FK_36a77cbe7ca9c7dd04ef4889634" FOREIGN KEY (hardware_id) REFERENCES public.socbox_detection_sensors_hardwares(id);


--
-- Name: iam_users_roles FK_3928031ee22160235ae1479ddcb; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.iam_users_roles
    ADD CONSTRAINT "FK_3928031ee22160235ae1479ddcb" FOREIGN KEY (user_id) REFERENCES public.iam_users(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: socbox_extras_custom_fields FK_3baf322671751b3fa10326175e3; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_extras_custom_fields
    ADD CONSTRAINT "FK_3baf322671751b3fa10326175e3" FOREIGN KEY (related_object_type_id) REFERENCES public.socbox_core_content_type(id) ON DELETE SET NULL;


--
-- Name: iocs_whitelist FK_45ad16fb296e2db00ce82d07a5d; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.iocs_whitelist
    ADD CONSTRAINT "FK_45ad16fb296e2db00ce82d07a5d" FOREIGN KEY ("ioctypeId") REFERENCES public.iocs_types(id);


--
-- Name: iam_roles_permissions FK_4930739bbc6ad5ebf0593bcc013; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.iam_roles_permissions
    ADD CONSTRAINT "FK_4930739bbc6ad5ebf0593bcc013" FOREIGN KEY (role_id) REFERENCES public.iam_roles(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: socbox_detection_sensors FK_4aacedae2096e4c6454d60adafd; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_detection_sensors
    ADD CONSTRAINT "FK_4aacedae2096e4c6454d60adafd" FOREIGN KEY (group_id) REFERENCES public.socbox_detection_sensors_groups(id);


--
-- Name: iocs_sources FK_62c372fe60491b58cfdfdc1bb15; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.iocs_sources
    ADD CONSTRAINT "FK_62c372fe60491b58cfdfdc1bb15" FOREIGN KEY (ioc_type_id) REFERENCES public.iocs_types(id);


--
-- Name: iocs FK_6efb2f31c1f409d1344dd61d3cd; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.iocs
    ADD CONSTRAINT "FK_6efb2f31c1f409d1344dd61d3cd" FOREIGN KEY (ioc_type_id) REFERENCES public.iocs_types(id);


--
-- Name: socbox_extras_changelog FK_80eeb9863efbdf09c0e2f772779; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_extras_changelog
    ADD CONSTRAINT "FK_80eeb9863efbdf09c0e2f772779" FOREIGN KEY (object_type_id) REFERENCES public.socbox_core_content_type(id) ON DELETE CASCADE;


--
-- Name: socbox_detection_mitre_techniques FK_82f86aa83c59b7d027b54c18fd2; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_detection_mitre_techniques
    ADD CONSTRAINT "FK_82f86aa83c59b7d027b54c18fd2" FOREIGN KEY (mitre_tactic_id) REFERENCES public.socbox_detection_mitre_tactics(id);


--
-- Name: socbox_extras_tagged_item FK_89569b40fb9910441c0d84c4fad; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_extras_tagged_item
    ADD CONSTRAINT "FK_89569b40fb9910441c0d84c4fad" FOREIGN KEY (object_type_id) REFERENCES public.socbox_core_content_type(id) ON DELETE CASCADE;


--
-- Name: socbox_extras_tag_object_types FK_9c97426508f7884a66a21a1e484; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_extras_tag_object_types
    ADD CONSTRAINT "FK_9c97426508f7884a66a21a1e484" FOREIGN KEY (tag_id) REFERENCES public.socbox_extras_tags(id) ON DELETE CASCADE;


--
-- Name: socbox_actors_objects_sensors FK_a5422c181d5e355977b60b13e70; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_actors_objects_sensors
    ADD CONSTRAINT "FK_a5422c181d5e355977b60b13e70" FOREIGN KEY (sensor_id) REFERENCES public.socbox_detection_sensors(id) ON UPDATE CASCADE ON DELETE SET NULL;


--
-- Name: socbox_extras_custom_fields_objects_types FK_a6a060f916314202e944d64cd36; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_extras_custom_fields_objects_types
    ADD CONSTRAINT "FK_a6a060f916314202e944d64cd36" FOREIGN KEY (custom_field_id) REFERENCES public.socbox_extras_custom_fields(id) ON DELETE CASCADE;


--
-- Name: socbox_extras_tag_object_types FK_a6c189d091a282a7f6bd7765197; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_extras_tag_object_types
    ADD CONSTRAINT "FK_a6c189d091a282a7f6bd7765197" FOREIGN KEY (object_type_id) REFERENCES public.socbox_core_content_type(id) ON DELETE CASCADE;


--
-- Name: iam_users_roles FK_b1a501895c3376b9a40cbda5006; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.iam_users_roles
    ADD CONSTRAINT "FK_b1a501895c3376b9a40cbda5006" FOREIGN KEY (role_id) REFERENCES public.iam_roles(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: socbox_actors_subjects FK_c68f2c7657166cd86628adf16a6; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_actors_subjects
    ADD CONSTRAINT "FK_c68f2c7657166cd86628adf16a6" FOREIGN KEY (city_id) REFERENCES public.socbox_geo_cities(id);


--
-- Name: socbox_geo_regions FK_d6007cc5afd6a051a1028ce9503; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_geo_regions
    ADD CONSTRAINT "FK_d6007cc5afd6a051a1028ce9503" FOREIGN KEY (district_id) REFERENCES public.socbox_geo_districts(id);


--
-- Name: socbox_extras_tagged_item FK_ef6541f3cf9694f5661233b8948; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_extras_tagged_item
    ADD CONSTRAINT "FK_ef6541f3cf9694f5661233b8948" FOREIGN KEY (tag_id) REFERENCES public.socbox_extras_tags(id) ON DELETE CASCADE;


--
-- Name: socbox_extras_custom_fields FK_f5c55f76821667997c512b59949; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_extras_custom_fields
    ADD CONSTRAINT "FK_f5c55f76821667997c512b59949" FOREIGN KEY (choiceset_id) REFERENCES public.socbox_extras_custom_fields_choicesets(id) ON DELETE SET NULL;


--
-- Name: iam_roles_permissions FK_fb13ed1d19e4a9af783c07563a9; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.iam_roles_permissions
    ADD CONSTRAINT "FK_fb13ed1d19e4a9af783c07563a9" FOREIGN KEY (permission_id) REFERENCES public.iam_permissions(id) ON DELETE CASCADE;


--
-- Name: socbox_extras_custom_fields_objects_types FK_fbb18c90eadaf19593f93cc547f; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_extras_custom_fields_objects_types
    ADD CONSTRAINT "FK_fbb18c90eadaf19593f93cc547f" FOREIGN KEY (object_type_id) REFERENCES public.socbox_core_content_type(id) ON DELETE CASCADE;


--
-- Name: dict_attack_methods dict_attack_methods_attack_class_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.dict_attack_methods
    ADD CONSTRAINT dict_attack_methods_attack_class_id_fkey FOREIGN KEY (attack_class_id) REFERENCES public.dict_attack_classes(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: dict_attack_methods dict_attack_methods_attack_level_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.dict_attack_methods
    ADD CONSTRAINT dict_attack_methods_attack_level_id_fkey FOREIGN KEY (attack_level_id) REFERENCES public.dict_attack_levels(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: dict_attack_protocols dict_attack_protocols_attack_level_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.dict_attack_protocols
    ADD CONSTRAINT dict_attack_protocols_attack_level_id_fkey FOREIGN KEY (attack_level_id) REFERENCES public.dict_attack_levels(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: dict_misp_attribute_types dict_misp_attribute_types_misp_attribute_category_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.dict_misp_attribute_types
    ADD CONSTRAINT dict_misp_attribute_types_misp_attribute_category_id_fkey FOREIGN KEY (misp_attribute_category_id) REFERENCES public.dict_misp_attribute_categories(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: dict_softs dict_softs_soft_type_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.dict_softs
    ADD CONSTRAINT dict_softs_soft_type_id_fkey FOREIGN KEY (soft_type_id) REFERENCES public.dict_soft_types(id) ON UPDATE CASCADE ON DELETE SET NULL;


--
-- Name: dict_vulnerabilities dict_vulnerabilities_attack_class_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.dict_vulnerabilities
    ADD CONSTRAINT dict_vulnerabilities_attack_class_id_fkey FOREIGN KEY (attack_class_id) REFERENCES public.dict_attack_classes(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: dict_vulnerabilities dict_vulnerabilities_soft_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.dict_vulnerabilities
    ADD CONSTRAINT dict_vulnerabilities_soft_id_fkey FOREIGN KEY (soft_id) REFERENCES public.dict_softs(id) ON UPDATE CASCADE ON DELETE SET NULL;


--
-- Name: event_attachments event_attachments_event_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.event_attachments
    ADD CONSTRAINT event_attachments_event_id_fkey FOREIGN KEY (event_id) REFERENCES public.events(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: event_attributes event_attributes_event_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.event_attributes
    ADD CONSTRAINT event_attributes_event_id_fkey FOREIGN KEY (event_id) REFERENCES public.events(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: event_attributes event_attributes_misp_attribute_type_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.event_attributes
    ADD CONSTRAINT event_attributes_misp_attribute_type_id_fkey FOREIGN KEY (misp_attribute_type_id) REFERENCES public.dict_misp_attribute_types(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: event_ext_ipaddresses event_ext_ipaddresses_event_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.event_ext_ipaddresses
    ADD CONSTRAINT event_ext_ipaddresses_event_id_fkey FOREIGN KEY (event_id) REFERENCES public.events(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: event_home_ipaddresses event_home_ipaddresses_event_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.event_home_ipaddresses
    ADD CONSTRAINT event_home_ipaddresses_event_id_fkey FOREIGN KEY (event_id) REFERENCES public.events(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: event_rules event_rules_event_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.event_rules
    ADD CONSTRAINT event_rules_event_id_fkey FOREIGN KEY (event_id) REFERENCES public.events(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: event_ttps event_ttps_event_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.event_ttps
    ADD CONSTRAINT event_ttps_event_id_fkey FOREIGN KEY (event_id) REFERENCES public.events(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: event_ttps event_ttps_mitre_techniques_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.event_ttps
    ADD CONSTRAINT event_ttps_mitre_techniques_id_fkey FOREIGN KEY (mitre_techniques_id) REFERENCES public.socbox_detection_mitre_techniques(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: events events_event_status_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.events
    ADD CONSTRAINT events_event_status_id_fkey FOREIGN KEY (event_status_id) REFERENCES public.dict_event_statuses(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: events events_object_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.events
    ADD CONSTRAINT events_object_id_fkey FOREIGN KEY (object_id) REFERENCES public.socbox_actors_objects(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: socbox_assets_homenets objects_homenets_object_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_assets_homenets
    ADD CONSTRAINT objects_homenets_object_id_fkey FOREIGN KEY (object_id) REFERENCES public.socbox_actors_objects(id) ON UPDATE CASCADE ON DELETE SET NULL;


--
-- Name: objects_networks objects_networks_object_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.objects_networks
    ADD CONSTRAINT objects_networks_object_id_fkey FOREIGN KEY (object_id) REFERENCES public.socbox_actors_objects(id) ON UPDATE CASCADE ON DELETE SET NULL;


--
-- Name: objects_persons objects_persons_object_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.objects_persons
    ADD CONSTRAINT objects_persons_object_id_fkey FOREIGN KEY (object_id) REFERENCES public.socbox_actors_objects(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: objects_persons objects_persons_person_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.objects_persons
    ADD CONSTRAINT objects_persons_person_id_fkey FOREIGN KEY (person_id) REFERENCES public.persons(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: socbox_assets_resources objects_resources_object_resource_type_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socbox_assets_resources
    ADD CONSTRAINT objects_resources_object_resource_type_id_fkey FOREIGN KEY (object_resource_type_id) REFERENCES public.socbox_assets_resources_types(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: persons persons_mode_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.persons
    ADD CONSTRAINT persons_mode_id_fkey FOREIGN KEY (mode_id) REFERENCES public.dict_person_modes(id) ON UPDATE CASCADE ON DELETE SET NULL;


--
-- Name: persons persons_order_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.persons
    ADD CONSTRAINT persons_order_id_fkey FOREIGN KEY (order_id) REFERENCES public.dict_person_orders(id) ON UPDATE CASCADE ON DELETE SET NULL;


--
-- Name: siem_actions siem_actions_siem_viewed_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.siem_actions
    ADD CONSTRAINT siem_actions_siem_viewed_id_fkey FOREIGN KEY (siem_viewed_id) REFERENCES public.siem_views(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: siem_filters siem_filters_siem_filter_type_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.siem_filters
    ADD CONSTRAINT siem_filters_siem_filter_type_id_fkey FOREIGN KEY (siem_filter_type_id) REFERENCES public.siem_filter_types(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: siem_tasks siem_tasks_assigned_to_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.siem_tasks
    ADD CONSTRAINT siem_tasks_assigned_to_fkey FOREIGN KEY (assigned_to) REFERENCES public.users(id) ON UPDATE CASCADE ON DELETE SET NULL;


--
-- Name: siem_tasks siem_tasks_siem_task_status_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.siem_tasks
    ADD CONSTRAINT siem_tasks_siem_task_status_id_fkey FOREIGN KEY (siem_task_status_id) REFERENCES public.siem_task_statuses(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: siem_tasks siem_tasks_siem_task_type_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.siem_tasks
    ADD CONSTRAINT siem_tasks_siem_task_type_id_fkey FOREIGN KEY (siem_task_type_id) REFERENCES public.siem_task_types(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: iocs_sources_statuses socbox_detection_iocs_sources_statuses_source_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.iocs_sources_statuses
    ADD CONSTRAINT socbox_detection_iocs_sources_statuses_source_id_fkey FOREIGN KEY (source_id) REFERENCES public.iocs_sources(id) ON DELETE CASCADE;


--
-- Name: users_histories users_histories_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users_histories
    ADD CONSTRAINT users_histories_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: users_metrics users_metrics_object_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users_metrics
    ADD CONSTRAINT users_metrics_object_id_fkey FOREIGN KEY (object_id) REFERENCES public.socbox_actors_objects(id) ON UPDATE CASCADE ON DELETE SET NULL;


--
-- Name: users_metrics users_metrics_user_action_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users_metrics
    ADD CONSTRAINT users_metrics_user_action_id_fkey FOREIGN KEY (user_action_id) REFERENCES public.users_actions(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: users users_org_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_org_id_fkey FOREIGN KEY (org_id) REFERENCES public.orgs(id) ON UPDATE CASCADE ON DELETE CASCADE;
