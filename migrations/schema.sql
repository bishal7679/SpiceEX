--
-- PostgreSQL database dump
--

-- Dumped from database version 15.4
-- Dumped by pg_dump version 15.4

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
-- Name: bookingsrestriction; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.bookingsrestriction (
    id integer NOT NULL,
    travelway character varying(10) NOT NULL,
    flying_from character varying(255) NOT NULL,
    flying_to character varying(255) NOT NULL,
    depart character varying NOT NULL,
    return character varying NOT NULL,
    country_code character varying(5) NOT NULL,
    mobile_no character varying(15) NOT NULL,
    booking_id integer,
    user_id integer NOT NULL,
    restriction_id integer NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.bookingsrestriction OWNER TO postgres;

--
-- Name: bookingsrestriction_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.bookingsrestriction_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.bookingsrestriction_id_seq OWNER TO postgres;

--
-- Name: bookingsrestriction_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.bookingsrestriction_id_seq OWNED BY public.bookingsrestriction.id;


--
-- Name: flightbookings; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.flightbookings (
    id integer NOT NULL,
    travelway character varying(10) NOT NULL,
    flying_from character varying(255) NOT NULL,
    flying_to character varying(255) NOT NULL,
    depart character varying NOT NULL,
    return character varying NOT NULL,
    travel_class character varying NOT NULL,
    full_name character varying(50) NOT NULL,
    address character varying(100) NOT NULL,
    email character varying(255) NOT NULL,
    country_code character varying(5) NOT NULL,
    mobile_no character varying(15) NOT NULL,
    pincode character varying(255) NOT NULL,
    city character varying(255) NOT NULL,
    state character varying(255) NOT NULL,
    govtidentity bytea NOT NULL,
    user_id integer NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.flightbookings OWNER TO postgres;

--
-- Name: flightbookings_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.flightbookings_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.flightbookings_id_seq OWNER TO postgres;

--
-- Name: flightbookings_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.flightbookings_id_seq OWNED BY public.flightbookings.id;


--
-- Name: payments; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.payments (
    id integer NOT NULL,
    email character varying(255) NOT NULL,
    card_number character varying(16) NOT NULL,
    validity character varying(255) NOT NULL,
    cvv integer NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.payments OWNER TO postgres;

--
-- Name: payments_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.payments_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.payments_id_seq OWNER TO postgres;

--
-- Name: payments_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.payments_id_seq OWNED BY public.payments.id;


--
-- Name: restrictions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.restrictions (
    id integer NOT NULL,
    restriction_name character varying(255) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.restrictions OWNER TO postgres;

--
-- Name: restrictions_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.restrictions_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.restrictions_id_seq OWNER TO postgres;

--
-- Name: restrictions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.restrictions_id_seq OWNED BY public.restrictions.id;


--
-- Name: schema_migration; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.schema_migration (
    version character varying(14) NOT NULL
);


ALTER TABLE public.schema_migration OWNER TO postgres;

--
-- Name: usersignup; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.usersignup (
    id integer NOT NULL,
    username character varying(50) NOT NULL,
    email character varying(255) NOT NULL,
    password character varying(60) NOT NULL,
    access_level integer DEFAULT 1 NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.usersignup OWNER TO postgres;

--
-- Name: usersignup_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.usersignup_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.usersignup_id_seq OWNER TO postgres;

--
-- Name: usersignup_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.usersignup_id_seq OWNED BY public.usersignup.id;


--
-- Name: bookingsrestriction id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bookingsrestriction ALTER COLUMN id SET DEFAULT nextval('public.bookingsrestriction_id_seq'::regclass);


--
-- Name: flightbookings id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.flightbookings ALTER COLUMN id SET DEFAULT nextval('public.flightbookings_id_seq'::regclass);


--
-- Name: payments id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.payments ALTER COLUMN id SET DEFAULT nextval('public.payments_id_seq'::regclass);


--
-- Name: restrictions id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.restrictions ALTER COLUMN id SET DEFAULT nextval('public.restrictions_id_seq'::regclass);


--
-- Name: usersignup id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.usersignup ALTER COLUMN id SET DEFAULT nextval('public.usersignup_id_seq'::regclass);


--
-- Name: bookingsrestriction bookingsrestriction_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bookingsrestriction
    ADD CONSTRAINT bookingsrestriction_pkey PRIMARY KEY (id);


--
-- Name: flightbookings flightbookings_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.flightbookings
    ADD CONSTRAINT flightbookings_pkey PRIMARY KEY (id);


--
-- Name: payments payments_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.payments
    ADD CONSTRAINT payments_pkey PRIMARY KEY (id);


--
-- Name: restrictions restrictions_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.restrictions
    ADD CONSTRAINT restrictions_pkey PRIMARY KEY (id);


--
-- Name: schema_migration schema_migration_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.schema_migration
    ADD CONSTRAINT schema_migration_pkey PRIMARY KEY (version);


--
-- Name: usersignup usersignup_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.usersignup
    ADD CONSTRAINT usersignup_pkey PRIMARY KEY (id);


--
-- Name: bookingsrestriction_booking_id_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX bookingsrestriction_booking_id_idx ON public.bookingsrestriction USING btree (booking_id);


--
-- Name: bookingsrestriction_depart_return_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX bookingsrestriction_depart_return_idx ON public.bookingsrestriction USING btree (depart, return);


--
-- Name: bookingsrestriction_flying_from_flying_to_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX bookingsrestriction_flying_from_flying_to_idx ON public.bookingsrestriction USING btree (flying_from, flying_to);


--
-- Name: flightbookings_email_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX flightbookings_email_idx ON public.flightbookings USING btree (email);


--
-- Name: flightbookings_full_name_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX flightbookings_full_name_idx ON public.flightbookings USING btree (full_name);


--
-- Name: schema_migration_version_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX schema_migration_version_idx ON public.schema_migration USING btree (version);


--
-- Name: usersignup_email_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX usersignup_email_idx ON public.usersignup USING btree (email);


--
-- Name: bookingsrestriction bookingsrestriction_flightbookings_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bookingsrestriction
    ADD CONSTRAINT bookingsrestriction_flightbookings_id_fk FOREIGN KEY (booking_id) REFERENCES public.flightbookings(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: bookingsrestriction bookingsrestriction_restrictions_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bookingsrestriction
    ADD CONSTRAINT bookingsrestriction_restrictions_id_fk FOREIGN KEY (restriction_id) REFERENCES public.restrictions(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: bookingsrestriction bookingsrestriction_usersignup_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bookingsrestriction
    ADD CONSTRAINT bookingsrestriction_usersignup_id_fk FOREIGN KEY (user_id) REFERENCES public.usersignup(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: flightbookings flightbookings_usersignup_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.flightbookings
    ADD CONSTRAINT flightbookings_usersignup_id_fk FOREIGN KEY (user_id) REFERENCES public.usersignup(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

