--
-- PostgreSQL database dump
--

-- Dumped from database version 13.1 (Debian 13.1-1.pgdg100+1)
-- Dumped by pg_dump version 13.1 (Debian 13.1-1.pgdg100+1)

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
-- Name: my_career_dev; Type: DATABASE; Schema: -; Owner: developer
--

ALTER DATABASE my_career_dev OWNER TO developer;

\connect my_career_dev

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
-- Name: department; Type: TABLE; Schema: public; Owner: developer
--

CREATE TABLE public.department (
    code integer NOT NULL,
    name text NOT NULL
);


ALTER TABLE public.department OWNER TO developer;

--
-- Name: TABLE department; Type: COMMENT; Schema: public; Owner: developer
--

COMMENT ON TABLE public.department IS 'DANE departments';


--
-- Name: document_type; Type: TABLE; Schema: public; Owner: developer
--

CREATE TABLE public.document_type (
    id character(5) NOT NULL,
    value text NOT NULL
);


ALTER TABLE public.document_type OWNER TO developer;

--
-- Name: TABLE document_type; Type: COMMENT; Schema: public; Owner: developer
--

COMMENT ON TABLE public.document_type IS 'Types of documents id';


--
-- Name: institution_type; Type: TABLE; Schema: public; Owner: developer
--

CREATE TABLE public.institution_type (
    id integer NOT NULL,
    value text NOT NULL
);


ALTER TABLE public.institution_type OWNER TO developer;

--
-- Name: TABLE institution_type; Type: COMMENT; Schema: public; Owner: developer
--

COMMENT ON TABLE public.institution_type IS 'If a public, private or semi public entity';


--
-- Name: institution_type_id_seq; Type: SEQUENCE; Schema: public; Owner: developer
--

CREATE SEQUENCE public.institution_type_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.institution_type_id_seq OWNER TO developer;

--
-- Name: institution_type_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: developer
--

ALTER SEQUENCE public.institution_type_id_seq OWNED BY public.institution_type.id;


--
-- Name: municipality; Type: TABLE; Schema: public; Owner: developer
--

CREATE TABLE public.municipality (
    code integer NOT NULL,
    department_code integer NOT NULL,
    name text NOT NULL
);


ALTER TABLE public.municipality OWNER TO developer;

--
-- Name: TABLE municipality; Type: COMMENT; Schema: public; Owner: developer
--

COMMENT ON TABLE public.municipality IS 'DANE municipality';


--
-- Name: study_level; Type: TABLE; Schema: public; Owner: developer
--

CREATE TABLE public.study_level (
    id integer NOT NULL,
    value text NOT NULL
);


ALTER TABLE public.study_level OWNER TO developer;

--
-- Name: TABLE study_level; Type: COMMENT; Schema: public; Owner: developer
--

COMMENT ON TABLE public.study_level IS 'Level of study';


--
-- Name: study_level_id_seq; Type: SEQUENCE; Schema: public; Owner: developer
--

CREATE SEQUENCE public.study_level_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.study_level_id_seq OWNER TO developer;

--
-- Name: study_level_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: developer
--

ALTER SEQUENCE public.study_level_id_seq OWNED BY public.study_level.id;


--
-- Name: user; Type: TABLE; Schema: public; Owner: developer
--

CREATE TABLE public."user" (
    first_name text NOT NULL,
    second_name text,
    first_surname text NOT NULL,
    second_surname text NOT NULL,
    email text NOT NULL,
    password text NOT NULL,
    document_type character(5) NOT NULL,
    institution_name text NOT NULL,
    study_level integer NOT NULL,
    institution_type integer NOT NULL,
    registry_confirmed boolean DEFAULT false NOT NULL,
    department_code integer NOT NULL,
    municipality_code integer NOT NULL
);


ALTER TABLE public."user" OWNER TO developer;

--
-- Name: TABLE "user"; Type: COMMENT; Schema: public; Owner: developer
--

COMMENT ON TABLE public."user" IS 'User information';


--
-- Name: institution_type id; Type: DEFAULT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.institution_type ALTER COLUMN id SET DEFAULT nextval('public.institution_type_id_seq'::regclass);


--
-- Name: study_level id; Type: DEFAULT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.study_level ALTER COLUMN id SET DEFAULT nextval('public.study_level_id_seq'::regclass);


--
-- Data for Name: department; Type: TABLE DATA; Schema: public; Owner: developer
--

COPY public.department (code, name) FROM stdin;
5	Antioquia
70	Sucre
\.


--
-- Data for Name: document_type; Type: TABLE DATA; Schema: public; Owner: developer
--

COPY public.document_type (id, value) FROM stdin;
CC   	Cédula de Ciudadanía
TI   	Tarjeta de Identidad
CE   	Cédula de Extranjería
\.


--
-- Data for Name: institution_type; Type: TABLE DATA; Schema: public; Owner: developer
--

COPY public.institution_type (id, value) FROM stdin;
1	Pública
2	Privada
3	Semi-privada
\.


--
-- Data for Name: municipality; Type: TABLE DATA; Schema: public; Owner: developer
--

COPY public.municipality (code, department_code, name) FROM stdin;
1	70	Sincelejo
221	70	Coveñas
1	5	Medellín
266	5	Envigado
\.


--
-- Data for Name: study_level; Type: TABLE DATA; Schema: public; Owner: developer
--

COPY public.study_level (id, value) FROM stdin;
1	Bachiller
2	Técnico
3	Tecnólogo
4	Universitario
5	Primaria
6	Posgrado
\.


--
-- Data for Name: user; Type: TABLE DATA; Schema: public; Owner: developer
--

COPY public."user" (first_name, second_name, first_surname, second_surname, email, password, document_type, institution_name, study_level, institution_type, registry_confirmed, department_code, municipality_code) FROM stdin;
\.


--
-- Name: institution_type_id_seq; Type: SEQUENCE SET; Schema: public; Owner: developer
--

SELECT pg_catalog.setval('public.institution_type_id_seq', 3, true);


--
-- Name: study_level_id_seq; Type: SEQUENCE SET; Schema: public; Owner: developer
--

SELECT pg_catalog.setval('public.study_level_id_seq', 6, true);


--
-- Name: department department_pk; Type: CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.department
    ADD CONSTRAINT department_pk PRIMARY KEY (code);


--
-- Name: document_type document_type_pk; Type: CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.document_type
    ADD CONSTRAINT document_type_pk PRIMARY KEY (id);


--
-- Name: institution_type institution_type_pk; Type: CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.institution_type
    ADD CONSTRAINT institution_type_pk PRIMARY KEY (id);


--
-- Name: municipality municipality_pk; Type: CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.municipality
    ADD CONSTRAINT municipality_pk PRIMARY KEY (department_code, code);


--
-- Name: study_level study_level_pk; Type: CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.study_level
    ADD CONSTRAINT study_level_pk PRIMARY KEY (id);


--
-- Name: user user_pk; Type: CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public."user"
    ADD CONSTRAINT user_pk PRIMARY KEY (email);


--
-- Name: department_code_uindex; Type: INDEX; Schema: public; Owner: developer
--

CREATE UNIQUE INDEX department_code_uindex ON public.department USING btree (code);


--
-- Name: document_type_id_uindex; Type: INDEX; Schema: public; Owner: developer
--

CREATE UNIQUE INDEX document_type_id_uindex ON public.document_type USING btree (id);


--
-- Name: user_email_uindex; Type: INDEX; Schema: public; Owner: developer
--

CREATE UNIQUE INDEX user_email_uindex ON public."user" USING btree (email);


--
-- Name: municipality municipality_department_code_fk; Type: FK CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.municipality
    ADD CONSTRAINT municipality_department_code_fk FOREIGN KEY (department_code) REFERENCES public.department(code);


--
-- Name: user user_document_type_fk; Type: FK CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public."user"
    ADD CONSTRAINT user_document_type_fk FOREIGN KEY (document_type) REFERENCES public.document_type(id);


--
-- Name: user user_institution_type_fk; Type: FK CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public."user"
    ADD CONSTRAINT user_institution_type_fk FOREIGN KEY (institution_type) REFERENCES public.institution_type(id);


--
-- Name: user user_municipality_code_department_code_fk; Type: FK CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public."user"
    ADD CONSTRAINT user_municipality_code_department_code_fk FOREIGN KEY (municipality_code, department_code) REFERENCES public.municipality(code, department_code);


--
-- Name: user user_study_level_fk; Type: FK CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public."user"
    ADD CONSTRAINT user_study_level_fk FOREIGN KEY (study_level) REFERENCES public.study_level(id);


--
-- PostgreSQL database dump complete
--

