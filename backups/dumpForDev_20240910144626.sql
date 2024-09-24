--
-- PostgreSQL database dump
--

-- Dumped from database version 16.4 (Debian 16.4-1.pgdg120+1)
-- Dumped by pg_dump version 16.4 (Debian 16.4-1.pgdg120+1)

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
-- Name: albums; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.albums (
    id_album integer NOT NULL,
    titre character varying(100) NOT NULL,
    annee_sortie integer,
    id_artiste integer,
    id_genre integer
);


--
-- Name: albums_id_album_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.albums_id_album_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: albums_id_album_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.albums_id_album_seq OWNED BY public.albums.id_album;


--
-- Name: artistes; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.artistes (
    id_artiste integer NOT NULL,
    nom character varying(100) NOT NULL,
    pays_origine character varying(50)
);


--
-- Name: artistes_id_artiste_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.artistes_id_artiste_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: artistes_id_artiste_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.artistes_id_artiste_seq OWNED BY public.artistes.id_artiste;


--
-- Name: genres; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.genres (
    id_genre integer NOT NULL,
    nom_genre character varying(50) NOT NULL
);


--
-- Name: genres_id_genre_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.genres_id_genre_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: genres_id_genre_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.genres_id_genre_seq OWNED BY public.genres.id_genre;


--
-- Name: albums id_album; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.albums ALTER COLUMN id_album SET DEFAULT nextval('public.albums_id_album_seq'::regclass);


--
-- Name: artistes id_artiste; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.artistes ALTER COLUMN id_artiste SET DEFAULT nextval('public.artistes_id_artiste_seq'::regclass);


--
-- Name: genres id_genre; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.genres ALTER COLUMN id_genre SET DEFAULT nextval('public.genres_id_genre_seq'::regclass);


--
-- Data for Name: albums; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.albums (id_album, titre, annee_sortie, id_artiste, id_genre) FROM stdin;
1	Abbey Road	1969	1	1
2	Kind of Blue	1959	2	2
3	Discovery	2001	3	3
\.


--
-- Data for Name: artistes; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.artistes (id_artiste, nom, pays_origine) FROM stdin;
1	The Beatles	Royaume-Uni
2	Miles Davis	États-Unis
3	Daft Punk	France
\.


--
-- Data for Name: genres; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.genres (id_genre, nom_genre) FROM stdin;
1	Rock
2	Jazz
3	Électronique
\.


--
-- Name: albums_id_album_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.albums_id_album_seq', 3, true);


--
-- Name: artistes_id_artiste_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.artistes_id_artiste_seq', 3, true);


--
-- Name: genres_id_genre_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.genres_id_genre_seq', 3, true);


--
-- Name: albums albums_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.albums
    ADD CONSTRAINT albums_pkey PRIMARY KEY (id_album);


--
-- Name: artistes artistes_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.artistes
    ADD CONSTRAINT artistes_pkey PRIMARY KEY (id_artiste);


--
-- Name: genres genres_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.genres
    ADD CONSTRAINT genres_pkey PRIMARY KEY (id_genre);


--
-- Name: albums albums_id_artiste_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.albums
    ADD CONSTRAINT albums_id_artiste_fkey FOREIGN KEY (id_artiste) REFERENCES public.artistes(id_artiste);


--
-- Name: albums albums_id_genre_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.albums
    ADD CONSTRAINT albums_id_genre_fkey FOREIGN KEY (id_genre) REFERENCES public.genres(id_genre);


--
-- PostgreSQL database dump complete
--

