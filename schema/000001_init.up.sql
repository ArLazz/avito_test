CREATE TABLE segments
(
    segment_id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    segment_name character varying(100) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT segments_pkey PRIMARY KEY (segment_id),
    CONSTRAINT unique_column_constraint UNIQUE (segment_name)
);

CREATE TABLE users
(
    user_id integer NOT NULL,
    segment integer NOT NULL,
    CONSTRAINT fk_segments_good FOREIGN KEY (segment)
        REFERENCES public.segments (segment_id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE
);;

CREATE TABLE usershistory
(
    user_id integer NOT NULL,
    segment_name character varying(100) COLLATE pg_catalog."default" NOT NULL,
    operation character varying(10) COLLATE pg_catalog."default" NOT NULL,
    "time" timestamp without time zone NOT NULL
);