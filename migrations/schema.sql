--
-- PostgreSQL database dump
--

-- Dumped from database version 14.8
-- Dumped by pg_dump version 14.8

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

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: performance; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.performance (
    id bigint NOT NULL,
    nano_id character varying(21),
    notes_count integer,
    notes_encoding smallint DEFAULT '0'::smallint,
    notes bytea,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    song_id bigint NOT NULL
);


ALTER TABLE public.performance OWNER TO postgres;

--
-- Name: performance_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.performance_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.performance_id_seq OWNER TO postgres;

--
-- Name: performance_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.performance_id_seq OWNED BY public.performance.id;


--
-- Name: schema_migration; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.schema_migration (
    version character varying(14) NOT NULL
);


ALTER TABLE public.schema_migration OWNER TO postgres;

--
-- Name: song; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.song (
    id bigint NOT NULL,
    nano_id character varying(21),
    title character varying(255),
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    img_url character varying(255) NOT NULL,
    music_path character varying(255) NOT NULL,
    url character varying(255) NOT NULL
);


ALTER TABLE public.song OWNER TO postgres;

--
-- Name: song_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.song_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.song_id_seq OWNER TO postgres;

--
-- Name: song_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.song_id_seq OWNED BY public.song.id;


--
-- Name: vp_user; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.vp_user (
    id bigint NOT NULL,
    nano_id character varying(21),
    username character varying(255),
    email character varying(255),
    email_confirmed boolean DEFAULT false,
    password_hash character varying(255),
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.vp_user OWNER TO postgres;

--
-- Name: vp_user_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.vp_user_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.vp_user_id_seq OWNER TO postgres;

--
-- Name: vp_user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.vp_user_id_seq OWNED BY public.vp_user.id;


--
-- Name: performance id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.performance ALTER COLUMN id SET DEFAULT nextval('public.performance_id_seq'::regclass);


--
-- Name: song id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.song ALTER COLUMN id SET DEFAULT nextval('public.song_id_seq'::regclass);


--
-- Name: vp_user id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.vp_user ALTER COLUMN id SET DEFAULT nextval('public.vp_user_id_seq'::regclass);


--
-- Name: performance performance_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.performance
    ADD CONSTRAINT performance_pkey PRIMARY KEY (id);


--
-- Name: schema_migration schema_migration_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.schema_migration
    ADD CONSTRAINT schema_migration_pkey PRIMARY KEY (version);


--
-- Name: song song_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.song
    ADD CONSTRAINT song_pkey PRIMARY KEY (id);


--
-- Name: vp_user vp_user_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.vp_user
    ADD CONSTRAINT vp_user_pkey PRIMARY KEY (id);


--
-- Name: schema_migration_version_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX schema_migration_version_idx ON public.schema_migration USING btree (version);


--
-- Name: performance performance_song_song_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.performance
    ADD CONSTRAINT performance_song_song_id_fk FOREIGN KEY (song_id) REFERENCES public.song(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

