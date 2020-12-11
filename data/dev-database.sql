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
-- Name: pgcrypto; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS pgcrypto WITH SCHEMA public;


--
-- Name: EXTENSION pgcrypto; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION pgcrypto IS 'cryptographic functions';


--
-- Name: authenticate_user(text, text); Type: FUNCTION; Schema: public; Owner: developer
--

CREATE FUNCTION public.authenticate_user(u_email text, u_pass text) RETURNS boolean
    LANGUAGE plpgsql
    AS $$
DECLARE
    match BOOLEAN DEFAULT FALSE;
BEGIN
    IF (check_user_existence(u_email)) THEN
        -- Probably could add registry confirmation validation.
        SELECT (password = crypt(u_pass, password))
        INTO match
        FROM "user"
        WHERE email = u_email;
    END IF;

    RETURN match;
END;
$$;


ALTER FUNCTION public.authenticate_user(u_email text, u_pass text) OWNER TO developer;

--
-- Name: check_user_existence(text); Type: FUNCTION; Schema: public; Owner: developer
--

CREATE FUNCTION public.check_user_existence(u_email text) RETURNS boolean
    LANGUAGE plpgsql
    AS $$
DECLARE
    exists BOOLEAN DEFAULT false;
BEGIN
    SELECT (count(email) > 0)
    INTO exists
    FROM "user"
    WHERE email = u_email;

    RETURN exists;
END;
$$;


ALTER FUNCTION public.check_user_existence(u_email text) OWNER TO developer;

--
-- Name: check_user_registry_confirmed(text); Type: FUNCTION; Schema: public; Owner: developer
--

CREATE FUNCTION public.check_user_registry_confirmed(u_email text) RETURNS boolean
    LANGUAGE plpgsql
    AS $$
DECLARE
    confirmed BOOLEAN DEFAULT false;
BEGIN
    SELECT COUNT(registry_confirmed) > 0
    INTO confirmed
    FROM "user"
    WHERE email = u_email
      AND registry_confirmed = true;

    RETURN confirmed;
END;
$$;


ALTER FUNCTION public.check_user_registry_confirmed(u_email text) OWNER TO developer;

--
-- Name: cypher_new_user_pass(); Type: FUNCTION; Schema: public; Owner: developer
--

CREATE FUNCTION public.cypher_new_user_pass() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    UPDATE "user" SET password = encrypt_user_password(NEW.password) WHERE email = NEW.email;
    RETURN NEW;
END;
$$;


ALTER FUNCTION public.cypher_new_user_pass() OWNER TO developer;

--
-- Name: encrypt_user_password(text); Type: FUNCTION; Schema: public; Owner: developer
--

CREATE FUNCTION public.encrypt_user_password(pass text) RETURNS text
    LANGUAGE sql
    AS $$
SELECT crypt(pass, gen_salt('bf'));
$$;


ALTER FUNCTION public.encrypt_user_password(pass text) OWNER TO developer;

--
-- Name: process_user_answer_audit(); Type: FUNCTION; Schema: public; Owner: developer
--

CREATE FUNCTION public.process_user_answer_audit() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF (TG_OP = 'DELETE') THEN
        INSERT INTO log_user_answer (operation, time_stamp, email, document_type, document, question, survey, answer)
        SELECT 'D', now(), OLD.email, OLD.document_type, OLD.document, OLD.question, OLD.survey, OLD.answer;
        RETURN OLD;
    ELSIF (TG_OP = 'UPDATE') THEN
        INSERT INTO log_user_answer (operation, time_stamp, email, document_type, document, question, survey, answer)
        SELECT 'U', now(), NEW.email, NEW.document_type, NEW.document, NEW.question, NEW.survey, NEW.answer;
        RETURN NEW;
    ELSIF (TG_OP = 'INSERT') THEN
        INSERT INTO log_user_answer (operation, time_stamp, email, document_type, document, question, survey, answer)
        SELECT 'I', now(), NEW.email, NEW.document_type, NEW.document, NEW.question, NEW.survey, NEW.answer;
        RETURN NEW;
    END IF;

    RETURN NULL; -- result is ignored since this is an AFTER trigger
END;
$$;


ALTER FUNCTION public.process_user_answer_audit() OWNER TO developer;

--
-- Name: process_user_audit(); Type: FUNCTION; Schema: public; Owner: developer
--

CREATE FUNCTION public.process_user_audit() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF (TG_OP = 'DELETE') THEN
        INSERT INTO log_user (operation, time_stamp, email, document_type, document, first_name, second_name,
                              first_surname, second_surname, password, institution_name, study_level, institution_type,
                              registry_confirmed, department_code, municipality_code, country_code)
        SELECT 'D',
               now(),
               OLD.email,
               OLD.document_type,
               OLD.document,
               OLD.first_name,
               OLD.second_name,
               OLD.first_surname,
               OLD.second_surname,
               encrypt_user_password(OLD.password),
               OLD.institution_name,
               OLD.study_level,
               OLD.institution_type,
               OLD.registry_confirmed,
               OLD.department_code,
               OLD.municipality_code,
               OLD.country_code;
        RETURN OLD;
    ELSIF (TG_OP = 'UPDATE') THEN
        INSERT INTO log_user (operation, time_stamp, email, document_type, document, first_name, second_name,
                              first_surname, second_surname, password, institution_name, study_level, institution_type,
                              registry_confirmed, department_code, municipality_code, country_code)
        SELECT 'U',
               now(),
               NEW.email,
               NEW.document_type,
               NEW.document,
               NEW.first_name,
               NEW.second_name,
               NEW.first_surname,
               NEW.second_surname,
               encrypt_user_password(NEW.password),
               NEW.institution_name,
               NEW.study_level,
               NEW.institution_type,
               NEW.registry_confirmed,
               NEW.department_code,
               NEW.municipality_code,
               NEW.country_code;
        RETURN NEW;
    ELSIF (TG_OP = 'INSERT') THEN
        INSERT INTO log_user (operation, time_stamp, email, document_type, document, first_name, second_name,
                              first_surname, second_surname, password, institution_name, study_level, institution_type,
                              registry_confirmed, department_code, municipality_code, country_code)
        SELECT 'I',
               now(),
               NEW.email,
               NEW.document_type,
               NEW.document,
               NEW.first_name,
               NEW.second_name,
               NEW.first_surname,
               NEW.second_surname,
               encrypt_user_password(NEW.password),
               NEW.institution_name,
               NEW.study_level,
               NEW.institution_type,
               NEW.registry_confirmed,
               NEW.department_code,
               NEW.municipality_code,
               NEW.country_code;
        RETURN NEW;
    END IF;

    RETURN NULL; -- result is ignored since this is an AFTER trigger
END;
$$;


ALTER FUNCTION public.process_user_audit() OWNER TO developer;

--
-- Name: retrieve_poll(integer); Type: FUNCTION; Schema: public; Owner: developer
--

CREATE FUNCTION public.retrieve_poll(poll_id integer) RETURNS TABLE(survey_id integer, survey_name text, question_id integer, question text, question_type text, option_id integer, option text)
    LANGUAGE sql
    AS $$
SELECT s.id AS surve_id,
       s.name,
       q.id AS question_id,
       q.question,
       qt.type,
       a.id AS option_id,
       a.option
FROM survey s
         INNER JOIN survey_question sq
                    ON s.id = sq.survey_id
         INNER JOIN question q ON q.id = sq.question_id
         INNER JOIN question_type qt ON q.question_type = qt.id
         LEFT JOIN answer_options ao ON q.answer_options = ao.code
         LEFT JOIN answer_option a ON ao.answer_option = a.id
WHERE s.id = poll_id AND s.active = true
ORDER BY q.id, a.option;
$$;


ALTER FUNCTION public.retrieve_poll(poll_id integer) OWNER TO developer;

--
-- Name: retrieve_survey_answers_by_user_survey(text, integer); Type: FUNCTION; Schema: public; Owner: developer
--

CREATE FUNCTION public.retrieve_survey_answers_by_user_survey(u_email text, survey_id integer) RETURNS TABLE(email text, question text, answer text)
    LANGUAGE sql
    AS $$
SELECT ua.email AS email,
       q.question,
       CASE
           WHEN ao.option is null THEN ua.answer
           ELSE ao.option
           END      answer
FROM user_answer ua
         INNER JOIN question q ON q.id = ua.question
         LEFT JOIN answer_option ao ON ua.answer = ao.id::text
WHERE ua.email = u_email AND ua.survey = survey_id;
$$;


ALTER FUNCTION public.retrieve_survey_answers_by_user_survey(u_email text, survey_id integer) OWNER TO developer;

SET default_tablespace = '';

SET default_table_access_method = heap;

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
    department_code text NOT NULL,
    municipality_code text NOT NULL,
    country_code text,
    document text NOT NULL
);


ALTER TABLE public."user" OWNER TO developer;

--
-- Name: TABLE "user"; Type: COMMENT; Schema: public; Owner: developer
--

COMMENT ON TABLE public."user" IS 'User information';


--
-- Name: update_user_pass(text, text, text); Type: FUNCTION; Schema: public; Owner: developer
--

CREATE FUNCTION public.update_user_pass(u_email text, old_pass text, new_pass text) RETURNS public."user"
    LANGUAGE plpgsql
    AS $$
DECLARE
    u_user "user";
BEGIN
    IF (authenticate_user(u_email, old_pass)) THEN
        UPDATE "user" SET password = encrypt_user_password(new_pass) WHERE email = u_email;
        SELECT * INTO u_user FROM "user" WHERE email = u_email;

        RETURN u_user;
    ELSE
        RAISE EXCEPTION '% must be auth', u_email;
    END IF;
END;
$$;


ALTER FUNCTION public.update_user_pass(u_email text, old_pass text, new_pass text) OWNER TO developer;

--
-- Name: answer_option; Type: TABLE; Schema: public; Owner: developer
--

CREATE TABLE public.answer_option (
    id integer NOT NULL,
    option text NOT NULL
);


ALTER TABLE public.answer_option OWNER TO developer;

--
-- Name: TABLE answer_option; Type: COMMENT; Schema: public; Owner: developer
--

COMMENT ON TABLE public.answer_option IS 'Option for the answer_options';


--
-- Name: answer_option_id_seq; Type: SEQUENCE; Schema: public; Owner: developer
--

CREATE SEQUENCE public.answer_option_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.answer_option_id_seq OWNER TO developer;

--
-- Name: answer_option_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: developer
--

ALTER SEQUENCE public.answer_option_id_seq OWNED BY public.answer_option.id;


--
-- Name: answer_options; Type: TABLE; Schema: public; Owner: developer
--

CREATE TABLE public.answer_options (
    id integer NOT NULL,
    code integer NOT NULL,
    answer_option integer
);


ALTER TABLE public.answer_options OWNER TO developer;

--
-- Name: TABLE answer_options; Type: COMMENT; Schema: public; Owner: developer
--

COMMENT ON TABLE public.answer_options IS 'Options for the question';


--
-- Name: answer_options_id_seq; Type: SEQUENCE; Schema: public; Owner: developer
--

CREATE SEQUENCE public.answer_options_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.answer_options_id_seq OWNER TO developer;

--
-- Name: answer_options_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: developer
--

ALTER SEQUENCE public.answer_options_id_seq OWNED BY public.answer_options.id;


--
-- Name: country; Type: TABLE; Schema: public; Owner: developer
--

CREATE TABLE public.country (
    iso_code text NOT NULL,
    name text NOT NULL
);


ALTER TABLE public.country OWNER TO developer;

--
-- Name: department; Type: TABLE; Schema: public; Owner: developer
--

CREATE TABLE public.department (
    code text NOT NULL,
    name text NOT NULL,
    country_code text NOT NULL
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
-- Name: log_user; Type: TABLE; Schema: public; Owner: developer
--

CREATE TABLE public.log_user (
    operation character(1) NOT NULL,
    time_stamp timestamp without time zone NOT NULL,
    email text,
    document_type character varying(5),
    document text,
    first_name text,
    second_name text,
    first_surname text,
    second_surname text,
    password text,
    institution_name text,
    study_level integer,
    institution_type integer,
    registry_confirmed boolean,
    department_code text,
    municipality_code text,
    country_code text
);


ALTER TABLE public.log_user OWNER TO developer;

--
-- Name: TABLE log_user; Type: COMMENT; Schema: public; Owner: developer
--

COMMENT ON TABLE public.log_user IS 'U after an I is the same, the U is the crypt password updated after a new register';


--
-- Name: log_user_answer; Type: TABLE; Schema: public; Owner: developer
--

CREATE TABLE public.log_user_answer (
    email text NOT NULL,
    survey integer,
    question integer,
    answer text,
    time_stamp timestamp without time zone NOT NULL,
    operation character(1) NOT NULL,
    document_type character varying(5),
    document text
);


ALTER TABLE public.log_user_answer OWNER TO developer;

--
-- Name: municipality; Type: TABLE; Schema: public; Owner: developer
--

CREATE TABLE public.municipality (
    code text NOT NULL,
    department_code text NOT NULL,
    name text NOT NULL,
    country_code text NOT NULL
);


ALTER TABLE public.municipality OWNER TO developer;

--
-- Name: TABLE municipality; Type: COMMENT; Schema: public; Owner: developer
--

COMMENT ON TABLE public.municipality IS 'DANE municipality';


--
-- Name: question; Type: TABLE; Schema: public; Owner: developer
--

CREATE TABLE public.question (
    id integer NOT NULL,
    question text NOT NULL,
    question_type integer NOT NULL,
    answer_options integer
);


ALTER TABLE public.question OWNER TO developer;

--
-- Name: TABLE question; Type: COMMENT; Schema: public; Owner: developer
--

COMMENT ON TABLE public.question IS 'Questions for polls';


--
-- Name: question_id_seq; Type: SEQUENCE; Schema: public; Owner: developer
--

CREATE SEQUENCE public.question_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.question_id_seq OWNER TO developer;

--
-- Name: question_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: developer
--

ALTER SEQUENCE public.question_id_seq OWNED BY public.question.id;


--
-- Name: question_type; Type: TABLE; Schema: public; Owner: developer
--

CREATE TABLE public.question_type (
    id integer NOT NULL,
    type text NOT NULL
);


ALTER TABLE public.question_type OWNER TO developer;

--
-- Name: TABLE question_type; Type: COMMENT; Schema: public; Owner: developer
--

COMMENT ON TABLE public.question_type IS 'Type of questions';


--
-- Name: question_type_id_seq; Type: SEQUENCE; Schema: public; Owner: developer
--

CREATE SEQUENCE public.question_type_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.question_type_id_seq OWNER TO developer;

--
-- Name: question_type_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: developer
--

ALTER SEQUENCE public.question_type_id_seq OWNED BY public.question_type.id;


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
-- Name: survey; Type: TABLE; Schema: public; Owner: developer
--

CREATE TABLE public.survey (
    id integer NOT NULL,
    name text NOT NULL,
    description text,
    active boolean DEFAULT false NOT NULL
);


ALTER TABLE public.survey OWNER TO developer;

--
-- Name: TABLE survey; Type: COMMENT; Schema: public; Owner: developer
--

COMMENT ON TABLE public.survey IS 'Survey information';


--
-- Name: survey_info_code_seq; Type: SEQUENCE; Schema: public; Owner: developer
--

CREATE SEQUENCE public.survey_info_code_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.survey_info_code_seq OWNER TO developer;

--
-- Name: survey_info_code_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: developer
--

ALTER SEQUENCE public.survey_info_code_seq OWNED BY public.survey.id;


--
-- Name: survey_question; Type: TABLE; Schema: public; Owner: developer
--

CREATE TABLE public.survey_question (
    survey_id integer NOT NULL,
    question_id integer NOT NULL
);


ALTER TABLE public.survey_question OWNER TO developer;

--
-- Name: TABLE survey_question; Type: COMMENT; Schema: public; Owner: developer
--

COMMENT ON TABLE public.survey_question IS 'Intermediate table for survey and question';


--
-- Name: user_answer; Type: TABLE; Schema: public; Owner: developer
--

CREATE TABLE public.user_answer (
    email text NOT NULL,
    question integer NOT NULL,
    answer text,
    survey integer NOT NULL,
    document_type character varying(5) NOT NULL,
    document text NOT NULL
);


ALTER TABLE public.user_answer OWNER TO developer;

--
-- Name: TABLE user_answer; Type: COMMENT; Schema: public; Owner: developer
--

COMMENT ON TABLE public.user_answer IS 'Answers for the user';


--
-- Name: COLUMN user_answer.answer; Type: COMMENT; Schema: public; Owner: developer
--

COMMENT ON COLUMN public.user_answer.answer IS 'Can be an answer_option or something';


--
-- Name: answer_option id; Type: DEFAULT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.answer_option ALTER COLUMN id SET DEFAULT nextval('public.answer_option_id_seq'::regclass);


--
-- Name: answer_options id; Type: DEFAULT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.answer_options ALTER COLUMN id SET DEFAULT nextval('public.answer_options_id_seq'::regclass);


--
-- Name: institution_type id; Type: DEFAULT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.institution_type ALTER COLUMN id SET DEFAULT nextval('public.institution_type_id_seq'::regclass);


--
-- Name: question id; Type: DEFAULT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.question ALTER COLUMN id SET DEFAULT nextval('public.question_id_seq'::regclass);


--
-- Name: question_type id; Type: DEFAULT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.question_type ALTER COLUMN id SET DEFAULT nextval('public.question_type_id_seq'::regclass);


--
-- Name: study_level id; Type: DEFAULT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.study_level ALTER COLUMN id SET DEFAULT nextval('public.study_level_id_seq'::regclass);


--
-- Name: survey id; Type: DEFAULT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.survey ALTER COLUMN id SET DEFAULT nextval('public.survey_info_code_seq'::regclass);


--
-- Data for Name: answer_option; Type: TABLE DATA; Schema: public; Owner: developer
--

COPY public.answer_option (id, option) FROM stdin;
1	Ingeniería de sistemas
2	Ingeniería mecánica
3	1
4	2
5	3
6	4
7	5
8	Sí
9	No
10	Ingeniería electrónica
11	Bioingeniería
12	Indufácil
13	Ingeniería eléctrica
14	Tal vez
\.


--
-- Data for Name: answer_options; Type: TABLE DATA; Schema: public; Owner: developer
--

COPY public.answer_options (id, code, answer_option) FROM stdin;
1	1	8
2	1	9
3	2	3
4	2	4
5	2	5
6	2	6
7	2	7
8	3	1
9	3	2
10	3	10
11	3	11
12	3	12
13	3	13
14	1	14
\.


--
-- Data for Name: country; Type: TABLE DATA; Schema: public; Owner: developer
--

COPY public.country (iso_code, name) FROM stdin;
CO	COLOMBIA
\.


--
-- Data for Name: department; Type: TABLE DATA; Schema: public; Owner: developer
--

COPY public.department (code, name, country_code) FROM stdin;
70	SUCRE	CO
05	ANTIOQUIA	CO
\.


--
-- Data for Name: document_type; Type: TABLE DATA; Schema: public; Owner: developer
--

COPY public.document_type (id, value) FROM stdin;
CC   	Cédula de Ciudadanía
TI   	Tarjeta de identidad
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
-- Data for Name: log_user; Type: TABLE DATA; Schema: public; Owner: developer
--

COPY public.log_user (operation, time_stamp, email, document_type, document, first_name, second_name, first_surname, second_surname, password, institution_name, study_level, institution_type, registry_confirmed, department_code, municipality_code, country_code) FROM stdin;
I	2020-11-21 06:59:04.228874	xephelsax@gmail.com	CC	1037656066	Leonardo	\N	Perez	Castilla	$2a$06$QKJMdh8TVNf89g4yCWK2Q.pdeuZlyIbhHSICkh.BOiEv4LCeMW4gu	Liceo Panamericano Campestre	4	2	f	70	001	CO
U	2020-11-21 06:59:04.228874	xephelsax@gmail.com	CC	1037656066	Leonardo	\N	Perez	Castilla	$2a$06$hfFCyJADUolN/DJ/ZrTStOSd0fingiNVl3h8P8vMhX5wzwFfxrTi.	Liceo Panamericano Campestre	4	2	f	70	001	CO
D	2020-11-21 07:03:50.21651	xephelsax@gmail.com	CC	1037656066	Leonardo	\N	Perez	Castilla	$2a$06$fQyf1ZOT2kcDKHlw/iJ6KeSAZG/IMmxmLQDd.NDo3Fy8Jffo09QVW	Liceo Panamericano Campestre	4	2	f	70	001	CO
I	2020-11-21 07:04:20.245328	xephelsax@gmail.com	CC	1037656066	Leonardo	\N	Perez	Castilla	$2a$06$qvCYBatFBEFTD60QZJi0u.NPTLCtV2MaJyTxqx5os89PF2TS0ahEm	Liceo Panamericano Campestre	4	2	f	70	001	CO
U	2020-11-21 07:04:20.245328	xephelsax@gmail.com	CC	1037656066	Leonardo	\N	Perez	Castilla	$2a$06$JkB68Ho1QOWcnG40aTlPRee0k4Bk72.Vq/IUMhj9ljxqexWkuq94u	Liceo Panamericano Campestre	4	2	f	70	001	CO
D	2020-11-21 08:24:38.180446	xephelsax@gmail.com	CC	1037656066	Leonardo	\N	Perez	Castilla	$2a$06$NT7hn1.jtFcIipP3VDu/6.oVL1QWC0VoWqSJz0RxJYyePhbKUEFUO	Liceo Panamericano Campestre	4	2	f	70	001	CO
I	2020-11-21 08:24:40.067754	xephelsax@gmail.com	CC	1037656066	Leonardo	Andres	Perez	Castilla	$2a$06$vPXxhY3o385XkBGUQC2Iv.iBdabS9VGCbSOnfliYgeEKNl0ygV9ge	Liceo Panamericano Campestre	4	2	f	70	001	CO
U	2020-11-21 08:24:40.067754	xephelsax@gmail.com	CC	1037656066	Leonardo	Andres	Perez	Castilla	$2a$06$yqUQ0/3WtbM4Jbx4Y.OMYOmXqjJ01MSjmDj4H5sKVhXTCCg7p5dd.	Liceo Panamericano Campestre	4	2	f	70	001	CO
D	2020-11-21 08:24:52.370495	xephelsax@gmail.com	CC	1037656066	Leonardo	Andres	Perez	Castilla	$2a$06$LrmKGiSX8doJs7ld20YCd.RipfkSs5dUc3QkrwPKi8VwjBA5v2SDW	Liceo Panamericano Campestre	4	2	f	70	001	CO
I	2020-11-21 08:25:34.92914	xephelsax@gmail.com	CC	1037656066	Leonardo	Andres	Perez	Castilla	$2a$06$MY2Q6xhcMeX67D5ok5xs9uwSEs/vJQUJlaot39wguCd0BFw/auj0u	Liceo Panamericano Campestre	4	2	f	70	001	CO
U	2020-11-21 08:25:34.92914	xephelsax@gmail.com	CC	1037656066	Leonardo	Andres	Perez	Castilla	$2a$06$3sJ9ZiwZu8z3iRsBIKcGJei20w0PAVRzq3mE7WLsx2OIsWabUmbyu	Liceo Panamericano Campestre	4	2	f	70	001	CO
D	2020-11-21 08:26:01.781104	xephelsax@gmail.com	CC	1037656066	Leonardo	Andres	Perez	Castilla	$2a$06$.FsbuItieQkoJCKsabpnFupJaenLnDwgf4Q0cW7iqwH4kFgXyzSjG	Liceo Panamericano Campestre	4	2	f	70	001	CO
I	2020-11-21 08:26:03.491667	xephelsax@gmail.com	CC	1037656066	Leonardo	Andres	Perez	Castilla	$2a$06$g7IVO0dBP/KPGAfoFHrQPOMSyXuqPW5mmtPNVNR6YB190CfQDrk.e	Liceo Panamericano Campestre	4	2	f	70	001	CO
U	2020-11-21 08:26:03.491667	xephelsax@gmail.com	CC	1037656066	Leonardo	Andres	Perez	Castilla	$2a$06$xJDA5EWSixI1Cjrv4.sO4uwNJcBWMwQYSz1k6UArKh0KCBudPOEZq	Liceo Panamericano Campestre	4	2	f	70	001	CO
D	2020-11-22 00:45:43.748477	xephelsax@gmail.com	CC	1037656066	Leonardo	Andres	Perez	Castilla	$2a$06$0yfn/okdVjwrQGEm013yKOmLOLgviXWQIB52gc.wk7OlNyd3br3iO	Liceo Panamericano Campestre	4	2	f	70	001	CO
I	2020-11-22 00:45:45.338048	xephelsax@gmail.com	CC	1037656066	Leonardo	Andres	Perez	Castilla	$2a$06$vyjke/8rwpN0fl9Iwews1upTfHLaM4FAQeZaAau1Dw7lPSxyMgFES	Liceo Panamericano Campestre	4	2	f	70	001	CO
U	2020-11-22 00:45:45.338048	xephelsax@gmail.com	CC	1037656066	Leonardo	Andres	Perez	Castilla	$2a$06$EAIjDQqQuhPW3EsIdF3ObOYJxwxTbHtizQtK5xLkG37hLipgR80pm	Liceo Panamericano Campestre	4	2	f	70	001	CO
U	2020-11-22 05:20:22.266248	xephelsax@gmail.com	CC	1037656066	Leonardo	Andres	Perez	Castilla	$2a$06$L9P/gEOn7F1h26PHZslHme2w8wW93f/k253vToS8ucBSuLJsW9fPa	Liceo Panamericano Campestre	4	2	t	70	001	CO
I	2020-11-23 05:39:52.919289	landres.perez@gmail.com	CC	23123	Leonardo	Andres	Perez	Castilla	$2a$06$HgZYoWqtbJhT52GYb3f4c.hVfbC0.RLJr.eNzT0WEQHyZ8q/iDFbe	Liceo Panamericano Campestre	4	2	f	70	001	CO
U	2020-11-23 05:39:52.919289	landres.perez@gmail.com	CC	23123	Leonardo	Andres	Perez	Castilla	$2a$06$FOvLwASzgXdoYQ7/jODuMe54MIG0p8eaDi.MmVHZgtQbgRDumPKBK	Liceo Panamericano Campestre	4	2	f	70	001	CO
I	2020-11-24 07:03:06.798534	ephelsa@hotmail.com	CC	1037656066	Leonardo		Pérez	Castilla	$2a$06$4CeSbR489tXye4D2.WKxWeeoAob8o4XYWxA1P7no7Ubo0VTUifFN2	Liceo Panamericano Campestre	4	2	f	70	001	CO
U	2020-11-24 07:03:06.798534	ephelsa@hotmail.com	CC	1037656066	Leonardo		Pérez	Castilla	$2a$06$zlvbIsnyLmfLaVcGtbNMx.Gqqr2iKUy2Zin11ZyAHWOlo2cVOwYDe	Liceo Panamericano Campestre	4	2	f	70	001	CO
U	2020-11-24 07:09:23.612914	ephelsa@hotmail.com	CC	1037656066	Leonardo		Pérez	Castilla	$2a$06$zkdlY7bPWoIJIxdQIVI.WuX4kW0P6utl9fKZa6OESPaIjvYig0fGG	Liceo Panamericano Campestre	4	2	t	70	001	CO
\.


--
-- Data for Name: log_user_answer; Type: TABLE DATA; Schema: public; Owner: developer
--

COPY public.log_user_answer (email, survey, question, answer, time_stamp, operation, document_type, document) FROM stdin;
xephelsax@gmail.com	1	1	8	2020-12-03 05:33:52.141387	I	CC	1037656066
xephelsax@gmail.com	1	1	xephelsax@gmail.com	2020-12-03 05:56:09.201606	U	CC	1037656066
xephelsax@gmail.com	1	1	8	2020-12-03 05:56:18.894657	U	CC	1037656066
xephelsax@gmail.com	1	2	8	2020-12-03 05:57:26.075741	I	CC	1037656066
xephelsax@gmail.com	1	2	9	2020-12-03 05:57:34.2466	U	CC	1037656066
xephelsax@gmail.com	1	1	8	2020-12-03 06:08:31.499158	U	CC	1037656066
xephelsax@gmail.com	1	1	8	2020-12-03 06:08:33.423087	U	CC	1037656066
xephelsax@gmail.com	1	1	8	2020-12-03 06:08:33.924898	U	CC	1037656066
xephelsax@gmail.com	1	1	8	2020-12-03 06:08:34.160088	U	CC	1037656066
xephelsax@gmail.com	1	1	8	2020-12-03 06:17:06.432437	U	CC	1037656066
xephelsax@gmail.com	1	1	8	2020-12-03 06:17:08.065597	U	CC	1037656066
xephelsax@gmail.com	1	1	8	2020-12-03 06:17:08.747803	U	CC	1037656066
xephelsax@gmail.com	1	1	8	2020-12-03 06:17:09.420431	U	CC	1037656066
xephelsax@gmail.com	1	1	8	2020-12-03 06:17:10.095956	U	CC	1037656066
xephelsax@gmail.com	1	2	1	2020-12-03 06:22:22.374415	U	CC	1037656066
xephelsax@gmail.com	1	1	8	2020-12-03 06:22:52.670989	U	CC	1037656066
xephelsax@gmail.com	1	1	Something	2020-12-03 06:23:03.239079	U	CC	1037656066
xephelsax@gmail.com	1	1	9	2020-12-03 06:23:10.783856	U	CC	1037656066
xephelsax@gmail.com	1	1	9	2020-12-03 06:23:46.354626	U	CC	1037656066
xephelsax@gmail.com	1	1	9	2020-12-03 06:24:42.984976	U	CC	1037656066
xephelsax@gmail.com	1	1	9	2020-12-03 06:28:15.981238	U	CC	1037656066
xephelsax@gmail.com	1	1	9	2020-12-03 06:28:19.328196	U	CC	1037656066
xephelsax@gmail.com	1	1	9	2020-12-03 14:19:23.143758	D	CC	1037656066
xephelsax@gmail.com	1	2	1	2020-12-03 14:19:23.143758	D	CC	1037656066
xephelsax@gmail.com	1	1	9	2020-12-03 14:19:26.556914	I	CC	1037656066
xephelsax@gmail.com	1	5	Liceo Campestre	2020-12-03 14:19:55.874435	I	CC	1037656066
xephelsax@gmail.com	1	5	Liceo Panamericano Campestre	2020-12-03 14:20:11.576963	U	CC	1037656066
xephelsax@gmail.com	1	5	Liceo Campestre	2020-12-03 14:20:30.205594	U	CC	1037656066
xephelsax@gmail.com	1	5	Liceo Panamericano Campestre	2020-12-03 14:20:41.858486	U	CC	1037656066
\.


--
-- Data for Name: municipality; Type: TABLE DATA; Schema: public; Owner: developer
--

COPY public.municipality (code, department_code, name, country_code) FROM stdin;
110	70	BUENAVISTA	CO
124	70	CAIMITO	CO
204	70	COLOSO	CO
215	70	COROZAL	CO
221	70	COVEÑAS	CO
230	70	CHALAN	CO
233	70	EL ROBLE	CO
235	70	GALERAS	CO
265	70	GUARANDA	CO
400	70	LA UNION	CO
418	70	LOS PALMITOS	CO
429	70	MAJAGUAL	CO
473	70	MORROA	CO
508	70	OVEJAS	CO
523	70	PALMITO	CO
670	70	SAMPUES	CO
678	70	SAN BENITO ABAD	CO
702	70	SAN JUAN DE BETULIA	CO
708	70	SAN MARCOS	CO
713	70	SAN ONOFRE	CO
717	70	SAN PEDRO	CO
742	70	SAN LUIS DE SINCE	CO
771	70	SUCRE	CO
820	70	SANTIAGO DE TOLU	CO
823	70	TOLU VIEJO	CO
001	70	SINCELEJO	CO
001	05	MEDELLIN	CO
002	05	ABEJORRAL	CO
004	05	ABRIAQUI	CO
021	05	ALEJANDRIA	CO
030	05	AMAGA	CO
031	05	AMALFI	CO
034	05	ANDES	CO
036	05	ANGELOPOLIS	CO
038	05	ANGOSTURA	CO
040	05	ANORI	CO
042	05	SANTAFE DE ANTIOQUIA	CO
044	05	ANZA	CO
045	05	APARTADO	CO
051	05	ARBOLETES	CO
055	05	ARGELIA	CO
059	05	ARMENIA	CO
079	05	BARBOSA	CO
086	05	BELMIRA	CO
088	05	BELLO	CO
091	05	BETANIA	CO
093	05	BETULIA	CO
101	05	CIUDAD BOLIVAR	CO
107	05	BRICEÑO	CO
113	05	BURITICA	CO
120	05	CACERES	CO
125	05	CAICEDO	CO
129	05	CALDAS	CO
134	05	CAMPAMENTO	CO
138	05	CAÑASGORDAS	CO
142	05	CARACOLI	CO
145	05	CARAMANTA	CO
147	05	CAREPA	CO
148	05	EL CARMEN DE VIBORAL	CO
150	05	CAROLINA	CO
154	05	CAUCASIA	CO
172	05	CHIGORODO	CO
190	05	CISNEROS	CO
197	05	COCORNA	CO
206	05	CONCEPCION	CO
209	05	CONCORDIA	CO
212	05	COPACABANA	CO
234	05	DABEIBA	CO
237	05	DON MATIAS	CO
240	05	EBEJICO	CO
250	05	EL BAGRE	CO
264	05	ENTRERRIOS	CO
266	05	ENVIGADO	CO
282	05	FREDONIA	CO
284	05	FRONTINO	CO
306	05	GIRALDO	CO
308	05	GIRARDOTA	CO
310	05	GOMEZ PLATA	CO
313	05	GRANADA	CO
315	05	GUADALUPE	CO
318	05	GUARNE	CO
321	05	GUATAPE	CO
347	05	HELICONIA	CO
353	05	HISPANIA	CO
360	05	ITAGUI	CO
361	05	ITUANGO	CO
364	05	JARDIN	CO
368	05	JERICO	CO
376	05	LA CEJA	CO
380	05	LA ESTRELLA	CO
390	05	LA PINTADA	CO
400	05	LA UNION	CO
411	05	LIBORINA	CO
425	05	MACEO	CO
440	05	MARINILLA	CO
467	05	MONTEBELLO	CO
475	05	MURINDO	CO
480	05	MUTATA	CO
483	05	NARIÑO	CO
490	05	NECOCLI	CO
495	05	NECHI	CO
501	05	OLAYA	CO
541	05	PEÐOL	CO
543	05	PEQUE	CO
576	05	PUEBLORRICO	CO
579	05	PUERTO BERRIO	CO
585	05	PUERTO NARE	CO
591	05	PUERTO TRIUNFO	CO
604	05	REMEDIOS	CO
607	05	RETIRO	CO
615	05	RIONEGRO	CO
628	05	SABANALARGA	CO
631	05	SABANETA	CO
642	05	SALGAR	CO
647	05	SAN ANDRES DE CUERQUIA	CO
649	05	SAN CARLOS	CO
652	05	SAN FRANCISCO	CO
656	05	SAN JERONIMO	CO
658	05	SAN JOSE DE LA MONTAÑA	CO
659	05	SAN JUAN DE URABA	CO
660	05	SAN LUIS	CO
664	05	SAN PEDRO	CO
665	05	SAN PEDRO DE URABA	CO
667	05	SAN RAFAEL	CO
670	05	SAN ROQUE	CO
674	05	SAN VICENTE	CO
679	05	SANTA BARBARA	CO
686	05	SANTA ROSA DE OSOS	CO
690	05	SANTO DOMINGO	CO
697	05	EL SANTUARIO	CO
736	05	SEGOVIA	CO
756	05	SONSON	CO
761	05	SOPETRAN	CO
789	05	TAMESIS	CO
790	05	TARAZA	CO
792	05	TARSO	CO
809	05	TITIRIBI	CO
819	05	TOLEDO	CO
837	05	TURBO	CO
842	05	URAMITA	CO
847	05	URRAO	CO
854	05	VALDIVIA	CO
856	05	VALPARAISO	CO
858	05	VEGACHI	CO
861	05	VENECIA	CO
873	05	VIGIA DEL FUERTE	CO
885	05	YALI	CO
887	05	YARUMAL	CO
890	05	YOLOMBO	CO
893	05	YONDO	CO
895	05	ZARAGOZA	CO
\.


--
-- Data for Name: question; Type: TABLE DATA; Schema: public; Owner: developer
--

COPY public.question (id, question, question_type, answer_options) FROM stdin;
1	¿Te gusta el queso?	1	1
2	¿Te gusta la patilla?	1	1
3	¿Qué tanto te gusta leer?	1	2
5	Institución donde estudias	2	\N
6	¿Qué carrera estudias?	1	3
\.


--
-- Data for Name: question_type; Type: TABLE DATA; Schema: public; Owner: developer
--

COPY public.question_type (id, type) FROM stdin;
1	select
2	text
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
-- Data for Name: survey; Type: TABLE DATA; Schema: public; Owner: developer
--

COPY public.survey (id, name, description, active) FROM stdin;
1	Perfilamiento de prueba	Esta es una simple pruebita	t
2	Otro questionario	\N	f
\.


--
-- Data for Name: survey_question; Type: TABLE DATA; Schema: public; Owner: developer
--

COPY public.survey_question (survey_id, question_id) FROM stdin;
1	1
1	2
2	3
2	1
1	5
1	6
\.


--
-- Data for Name: user; Type: TABLE DATA; Schema: public; Owner: developer
--

COPY public."user" (first_name, second_name, first_surname, second_surname, email, password, document_type, institution_name, study_level, institution_type, registry_confirmed, department_code, municipality_code, country_code, document) FROM stdin;
Leonardo	Andres	Perez	Castilla	xephelsax@gmail.com	$2a$06$Ya0biuVS63LTX4XQ0qABruXfLQ5NWo4ksOw6kD139b6rzcN5L9vYK	CC   	Liceo Panamericano Campestre	4	2	t	70	001	CO	1037656066
Leonardo	Andres	Perez	Castilla	landres.perez@gmail.com	$2a$06$.shV05nG810Eqx2pv4fI1.7VBZEqBxQr59/b.0/xN/ECbfPzZ43By	CC   	Liceo Panamericano Campestre	4	2	f	70	001	CO	23123
Leonardo		Pérez	Castilla	ephelsa@hotmail.com	$2a$06$Ng88bxzLnKixb8nxX7Gg4.QblbDw2H7UEF2I6kjAu3hd8KZjJ18US	CC   	Liceo Panamericano Campestre	4	2	t	70	001	CO	1037656066
\.


--
-- Data for Name: user_answer; Type: TABLE DATA; Schema: public; Owner: developer
--

COPY public.user_answer (email, question, answer, survey, document_type, document) FROM stdin;
xephelsax@gmail.com	1	9	1	CC	1037656066
xephelsax@gmail.com	5	Liceo Panamericano Campestre	1	CC	1037656066
\.


--
-- Name: answer_option_id_seq; Type: SEQUENCE SET; Schema: public; Owner: developer
--

SELECT pg_catalog.setval('public.answer_option_id_seq', 14, true);


--
-- Name: answer_options_id_seq; Type: SEQUENCE SET; Schema: public; Owner: developer
--

SELECT pg_catalog.setval('public.answer_options_id_seq', 14, true);


--
-- Name: institution_type_id_seq; Type: SEQUENCE SET; Schema: public; Owner: developer
--

SELECT pg_catalog.setval('public.institution_type_id_seq', 3, true);


--
-- Name: question_id_seq; Type: SEQUENCE SET; Schema: public; Owner: developer
--

SELECT pg_catalog.setval('public.question_id_seq', 6, true);


--
-- Name: question_type_id_seq; Type: SEQUENCE SET; Schema: public; Owner: developer
--

SELECT pg_catalog.setval('public.question_type_id_seq', 2, true);


--
-- Name: study_level_id_seq; Type: SEQUENCE SET; Schema: public; Owner: developer
--

SELECT pg_catalog.setval('public.study_level_id_seq', 6, true);


--
-- Name: survey_info_code_seq; Type: SEQUENCE SET; Schema: public; Owner: developer
--

SELECT pg_catalog.setval('public.survey_info_code_seq', 2, true);


--
-- Name: answer_option answer_option_pk; Type: CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.answer_option
    ADD CONSTRAINT answer_option_pk PRIMARY KEY (id);


--
-- Name: answer_options answer_options_pk; Type: CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.answer_options
    ADD CONSTRAINT answer_options_pk PRIMARY KEY (id);


--
-- Name: country country_pk; Type: CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.country
    ADD CONSTRAINT country_pk PRIMARY KEY (iso_code);


--
-- Name: department department_pk; Type: CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.department
    ADD CONSTRAINT department_pk PRIMARY KEY (country_code, code);


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
    ADD CONSTRAINT municipality_pk PRIMARY KEY (country_code, department_code, code);


--
-- Name: question question_pk; Type: CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.question
    ADD CONSTRAINT question_pk PRIMARY KEY (id);


--
-- Name: question_type question_type_pk; Type: CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.question_type
    ADD CONSTRAINT question_type_pk PRIMARY KEY (id);


--
-- Name: study_level study_level_pk; Type: CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.study_level
    ADD CONSTRAINT study_level_pk PRIMARY KEY (id);


--
-- Name: survey survey_info_pk; Type: CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.survey
    ADD CONSTRAINT survey_info_pk PRIMARY KEY (id);


--
-- Name: survey_question survey_question_pk; Type: CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.survey_question
    ADD CONSTRAINT survey_question_pk PRIMARY KEY (question_id, survey_id);


--
-- Name: user_answer user_answer_pk; Type: CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.user_answer
    ADD CONSTRAINT user_answer_pk PRIMARY KEY (email, document_type, document, survey, question);


--
-- Name: user user_pk; Type: CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public."user"
    ADD CONSTRAINT user_pk PRIMARY KEY (email, document, document_type);


--
-- Name: answer_option_option_uindex; Type: INDEX; Schema: public; Owner: developer
--

CREATE UNIQUE INDEX answer_option_option_uindex ON public.answer_option USING btree (option);


--
-- Name: department_code_uindex; Type: INDEX; Schema: public; Owner: developer
--

CREATE UNIQUE INDEX department_code_uindex ON public.department USING btree (code);


--
-- Name: document_type_id_uindex; Type: INDEX; Schema: public; Owner: developer
--

CREATE UNIQUE INDEX document_type_id_uindex ON public.document_type USING btree (id);


--
-- Name: question_type_type_uindex; Type: INDEX; Schema: public; Owner: developer
--

CREATE UNIQUE INDEX question_type_type_uindex ON public.question_type USING btree (type);


--
-- Name: user_email_uindex; Type: INDEX; Schema: public; Owner: developer
--

CREATE UNIQUE INDEX user_email_uindex ON public."user" USING btree (email);


--
-- Name: user log_user; Type: TRIGGER; Schema: public; Owner: developer
--

CREATE TRIGGER log_user AFTER INSERT OR DELETE OR UPDATE ON public."user" FOR EACH ROW EXECUTE FUNCTION public.process_user_audit();


--
-- Name: user_answer log_user_answer; Type: TRIGGER; Schema: public; Owner: developer
--

CREATE TRIGGER log_user_answer AFTER INSERT OR DELETE OR UPDATE ON public.user_answer FOR EACH ROW EXECUTE FUNCTION public.process_user_answer_audit();


--
-- Name: user new_user; Type: TRIGGER; Schema: public; Owner: developer
--

CREATE TRIGGER new_user AFTER INSERT ON public."user" FOR EACH ROW EXECUTE FUNCTION public.cypher_new_user_pass();


--
-- Name: answer_options answer_options_answer_option_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.answer_options
    ADD CONSTRAINT answer_options_answer_option_id_fk FOREIGN KEY (answer_option) REFERENCES public.answer_option(id);


--
-- Name: department department_country_code_fk; Type: FK CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.department
    ADD CONSTRAINT department_country_code_fk FOREIGN KEY (country_code) REFERENCES public.country(iso_code);


--
-- Name: log_user_answer log_user_answer_user_email_document_document_type_fk; Type: FK CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.log_user_answer
    ADD CONSTRAINT log_user_answer_user_email_document_document_type_fk FOREIGN KEY (email, document, document_type) REFERENCES public."user"(email, document, document_type);


--
-- Name: log_user log_user_institution_type_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.log_user
    ADD CONSTRAINT log_user_institution_type_id_fk FOREIGN KEY (institution_type) REFERENCES public.institution_type(id);


--
-- Name: log_user log_user_municipality_code_country_code_department_code_fk; Type: FK CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.log_user
    ADD CONSTRAINT log_user_municipality_code_country_code_department_code_fk FOREIGN KEY (municipality_code, country_code, department_code) REFERENCES public.municipality(code, country_code, department_code);


--
-- Name: log_user log_user_study_level_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.log_user
    ADD CONSTRAINT log_user_study_level_id_fk FOREIGN KEY (study_level) REFERENCES public.study_level(id);


--
-- Name: municipality municipality_department_code_country_code_fk; Type: FK CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.municipality
    ADD CONSTRAINT municipality_department_code_country_code_fk FOREIGN KEY (department_code, country_code) REFERENCES public.department(code, country_code);


--
-- Name: question question_question_type_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.question
    ADD CONSTRAINT question_question_type_id_fk FOREIGN KEY (question_type) REFERENCES public.question_type(id);


--
-- Name: survey_question survey_question_question_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.survey_question
    ADD CONSTRAINT survey_question_question_id_fk FOREIGN KEY (question_id) REFERENCES public.question(id);


--
-- Name: survey_question survey_question_survey_info_code_fk; Type: FK CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.survey_question
    ADD CONSTRAINT survey_question_survey_info_code_fk FOREIGN KEY (survey_id) REFERENCES public.survey(id);


--
-- Name: user_answer user_answer_survey_question_question_id_survey_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.user_answer
    ADD CONSTRAINT user_answer_survey_question_question_id_survey_id_fk FOREIGN KEY (question, survey) REFERENCES public.survey_question(question_id, survey_id);


--
-- Name: user_answer user_answer_user_document_type_document_email_fk; Type: FK CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.user_answer
    ADD CONSTRAINT user_answer_user_document_type_document_email_fk FOREIGN KEY (document_type, document, email) REFERENCES public."user"(document_type, document, email);


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
-- Name: user user_municipality_code_country_code_department_code_fk; Type: FK CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public."user"
    ADD CONSTRAINT user_municipality_code_country_code_department_code_fk FOREIGN KEY (municipality_code, country_code, department_code) REFERENCES public.municipality(code, country_code, department_code);


--
-- Name: user user_study_level_fk; Type: FK CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public."user"
    ADD CONSTRAINT user_study_level_fk FOREIGN KEY (study_level) REFERENCES public.study_level(id);


--
-- PostgreSQL database dump complete
--

