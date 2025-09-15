CREATE TABLE public.users_actions (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    updated_at timestamp with time zone,
    created_at timestamp with time zone,
    updated_by integer,
    created_by integer
);
