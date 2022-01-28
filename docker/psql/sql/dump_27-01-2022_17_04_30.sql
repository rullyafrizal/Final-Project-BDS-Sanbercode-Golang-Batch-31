--
-- PostgreSQL database cluster dump
--

SET default_transaction_read_only = off;

SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;

--
-- Drop databases (except postgres and template1)
--

DROP DATABASE api_blog_b31;




--
-- Drop roles
--

DROP ROLE blog;
DROP ROLE postgres;
DROP ROLE rully;


--
-- Roles
--

CREATE ROLE blog;
ALTER ROLE blog WITH SUPERUSER INHERIT NOCREATEROLE CREATEDB LOGIN NOREPLICATION NOBYPASSRLS PASSWORD 'SCRAM-SHA-256$4096:CUzGk0OFqYJwt+Bq1MUmaA==$q9PYHfhzkLWyD7xRNfTc9irJ+G4Mq6h6RBP6P9yx0Lg=:tHHWNmm4o2img39X/ooNMpKBcNTdWbwAPL726+mO+UU=';
CREATE ROLE postgres;
ALTER ROLE postgres WITH SUPERUSER INHERIT CREATEROLE CREATEDB LOGIN REPLICATION BYPASSRLS PASSWORD 'SCRAM-SHA-256$4096:dfHH1OKjkkG9hz8cb0frjw==$r5ijQsdkL3tqel4ebDr4l+W+ZODsvFNMVUlKYNCAcgM=:+g3xGmPt5ZWnuyRi+b24n3/ByZjjBUL26IYZM7BuvvM=';
CREATE ROLE rully;
ALTER ROLE rully WITH SUPERUSER INHERIT NOCREATEROLE CREATEDB LOGIN NOREPLICATION NOBYPASSRLS;






--
-- Databases
--

--
-- Database "template1" dump
--

--
-- PostgreSQL database dump
--

-- Dumped from database version 14.1
-- Dumped by pg_dump version 14.1

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

UPDATE pg_catalog.pg_database SET datistemplate = false WHERE datname = 'template1';
DROP DATABASE template1;
--
-- Name: template1; Type: DATABASE; Schema: -; Owner: postgres
--

CREATE DATABASE template1 WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE = 'en_US.utf8';


ALTER DATABASE template1 OWNER TO postgres;

\connect template1

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
-- Name: DATABASE template1; Type: COMMENT; Schema: -; Owner: postgres
--

COMMENT ON DATABASE template1 IS 'default template for new databases';


--
-- Name: template1; Type: DATABASE PROPERTIES; Schema: -; Owner: postgres
--

ALTER DATABASE template1 IS_TEMPLATE = true;


\connect template1

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
-- Name: DATABASE template1; Type: ACL; Schema: -; Owner: postgres
--

REVOKE CONNECT,TEMPORARY ON DATABASE template1 FROM PUBLIC;
GRANT CONNECT ON DATABASE template1 TO PUBLIC;


--
-- PostgreSQL database dump complete
--

--
-- Database "api_blog_b31" dump
--

--
-- PostgreSQL database dump
--

-- Dumped from database version 14.1
-- Dumped by pg_dump version 14.1

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
-- Name: api_blog_b31; Type: DATABASE; Schema: -; Owner: blog
--

CREATE DATABASE api_blog_b31 WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE = 'en_US.utf8';


ALTER DATABASE api_blog_b31 OWNER TO blog;

\connect api_blog_b31

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
-- Name: post_images; Type: TABLE; Schema: public; Owner: blog
--

CREATE TABLE public.post_images (
    id bigint NOT NULL,
    url text,
    post_id bigint
);


ALTER TABLE public.post_images OWNER TO blog;

--
-- Name: post_images_id_seq; Type: SEQUENCE; Schema: public; Owner: blog
--

CREATE SEQUENCE public.post_images_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.post_images_id_seq OWNER TO blog;

--
-- Name: post_images_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: blog
--

ALTER SEQUENCE public.post_images_id_seq OWNED BY public.post_images.id;


--
-- Name: post_tags; Type: TABLE; Schema: public; Owner: blog
--

CREATE TABLE public.post_tags (
    post_id bigint NOT NULL,
    tag_id bigint NOT NULL
);


ALTER TABLE public.post_tags OWNER TO blog;

--
-- Name: posts; Type: TABLE; Schema: public; Owner: blog
--

CREATE TABLE public.posts (
    id bigint NOT NULL,
    title text,
    content text,
    is_published boolean,
    user_id bigint,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.posts OWNER TO blog;

--
-- Name: posts_id_seq; Type: SEQUENCE; Schema: public; Owner: blog
--

CREATE SEQUENCE public.posts_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.posts_id_seq OWNER TO blog;

--
-- Name: posts_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: blog
--

ALTER SEQUENCE public.posts_id_seq OWNED BY public.posts.id;


--
-- Name: reviews; Type: TABLE; Schema: public; Owner: blog
--

CREATE TABLE public.reviews (
    id bigint NOT NULL,
    user_id bigint,
    post_id bigint,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    comment text,
    rating numeric
);


ALTER TABLE public.reviews OWNER TO blog;

--
-- Name: reviews_id_seq; Type: SEQUENCE; Schema: public; Owner: blog
--

CREATE SEQUENCE public.reviews_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.reviews_id_seq OWNER TO blog;

--
-- Name: reviews_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: blog
--

ALTER SEQUENCE public.reviews_id_seq OWNED BY public.reviews.id;


--
-- Name: roles; Type: TABLE; Schema: public; Owner: blog
--

CREATE TABLE public.roles (
    id bigint NOT NULL,
    name text
);


ALTER TABLE public.roles OWNER TO blog;

--
-- Name: roles_id_seq; Type: SEQUENCE; Schema: public; Owner: blog
--

CREATE SEQUENCE public.roles_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.roles_id_seq OWNER TO blog;

--
-- Name: roles_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: blog
--

ALTER SEQUENCE public.roles_id_seq OWNED BY public.roles.id;


--
-- Name: tags; Type: TABLE; Schema: public; Owner: blog
--

CREATE TABLE public.tags (
    id bigint NOT NULL,
    name text
);


ALTER TABLE public.tags OWNER TO blog;

--
-- Name: tags_id_seq; Type: SEQUENCE; Schema: public; Owner: blog
--

CREATE SEQUENCE public.tags_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.tags_id_seq OWNER TO blog;

--
-- Name: tags_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: blog
--

ALTER SEQUENCE public.tags_id_seq OWNED BY public.tags.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: blog
--

CREATE TABLE public.users (
    id bigint NOT NULL,
    name text,
    email text,
    password text,
    role_id bigint,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    avatar text
);


ALTER TABLE public.users OWNER TO blog;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: blog
--

CREATE SEQUENCE public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO blog;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: blog
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: post_images id; Type: DEFAULT; Schema: public; Owner: blog
--

ALTER TABLE ONLY public.post_images ALTER COLUMN id SET DEFAULT nextval('public.post_images_id_seq'::regclass);


--
-- Name: posts id; Type: DEFAULT; Schema: public; Owner: blog
--

ALTER TABLE ONLY public.posts ALTER COLUMN id SET DEFAULT nextval('public.posts_id_seq'::regclass);


--
-- Name: reviews id; Type: DEFAULT; Schema: public; Owner: blog
--

ALTER TABLE ONLY public.reviews ALTER COLUMN id SET DEFAULT nextval('public.reviews_id_seq'::regclass);


--
-- Name: roles id; Type: DEFAULT; Schema: public; Owner: blog
--

ALTER TABLE ONLY public.roles ALTER COLUMN id SET DEFAULT nextval('public.roles_id_seq'::regclass);


--
-- Name: tags id; Type: DEFAULT; Schema: public; Owner: blog
--

ALTER TABLE ONLY public.tags ALTER COLUMN id SET DEFAULT nextval('public.tags_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: blog
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Data for Name: post_images; Type: TABLE DATA; Schema: public; Owner: blog
--

COPY public.post_images (id, url, post_id) FROM stdin;
1	fasfafa	2
2	dfasfaf	2
3	https://picsum.photos/200/300	3
6	https://picsum.photos/200/300	5
7	https://picsum.photos/200/300	5
4	https://picsum.photos/200/300	4
5	https://picsum.photos/200/300	4
\.


--
-- Data for Name: post_tags; Type: TABLE DATA; Schema: public; Owner: blog
--

COPY public.post_tags (post_id, tag_id) FROM stdin;
2	5
2	6
3	5
3	6
4	7
4	8
5	7
5	8
5	9
\.


--
-- Data for Name: posts; Type: TABLE DATA; Schema: public; Owner: blog
--

COPY public.posts (id, title, content, is_published, user_id, created_at, updated_at) FROM stdin;
2	Ini adalah title	Quisque velit nisi, pretium ut lacinia in, elementum id enim. Pellentesque in ipsum id orci porta dapibus.	f	4	2022-01-26 06:57:26.940142+00	2022-01-26 06:57:26.940142+00
3	Ini adalah title	Quisque velit nisi, pretium ut lacinia in, elementum id enim. Pellentesque in ipsum id orci porta dapibus.	f	4	2022-01-26 07:38:05.696573+00	2022-01-26 07:38:05.696573+00
5	Postingan dari John Doe Yang Kedua	Quisque velit nisi, pretium ut lacinia in, elementum id enim. Pellentesque in ipsum id orci porta dapibus.	f	6	2022-01-27 08:06:44.903473+00	2022-01-27 08:06:44.903473+00
4	Postingan dari John Doe	Quisque velit nisi, pretium ut lacinia in, elementum id enim. Pellentesque in ipsum id orci porta dapibus.	t	6	2022-01-27 08:06:27.56561+00	2022-01-27 08:45:02.445362+00
\.


--
-- Data for Name: reviews; Type: TABLE DATA; Schema: public; Owner: blog
--

COPY public.reviews (id, user_id, post_id, created_at, updated_at, comment, rating) FROM stdin;
2	6	4	2022-01-27 08:10:06.786067+00	2022-01-27 08:10:06.786067+00	Postingan ini keren	4
\.


--
-- Data for Name: roles; Type: TABLE DATA; Schema: public; Owner: blog
--

COPY public.roles (id, name) FROM stdin;
1	admin
2	user
\.


--
-- Data for Name: tags; Type: TABLE DATA; Schema: public; Owner: blog
--

COPY public.tags (id, name) FROM stdin;
5	halo
6	bagus
7	fashion
8	olahraga
9	misteri
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: blog
--

COPY public.users (id, name, email, password, role_id, created_at, updated_at, avatar) FROM stdin;
8	Admin	admin@admin.com	$2a$14$MmndCeaq7z7/KxzAk3uLOOm4BINctwAOCNx4DBBGe5VVPSFd5.7dy	1	2022-01-27 05:26:46.018017+00	2022-01-27 05:26:46.018017+00	https://picsum.photos/200
4	Rully Afrizal	rullyafrizal4@gmail.com	$2a$14$wY/7KTi21bJzAC7UsuC5PuDzgf0uLqgHVngHeGLjYMweV3zvMQ8La	2	2022-01-25 08:44:12.175378+00	2022-01-25 08:44:12.175378+00	\N
6	John Doe Edited	johndoe@gmail.com	$2a$14$kxDVuWcL45dmuUmyAase3u4uO6HurlT9SLXdVfoorWeBRG/AO2rlu	2	2022-01-27 03:37:33.785747+00	2022-01-27 05:53:58.432186+00	https://picsum.photos/200/300
\.


--
-- Name: post_images_id_seq; Type: SEQUENCE SET; Schema: public; Owner: blog
--

SELECT pg_catalog.setval('public.post_images_id_seq', 7, true);


--
-- Name: posts_id_seq; Type: SEQUENCE SET; Schema: public; Owner: blog
--

SELECT pg_catalog.setval('public.posts_id_seq', 5, true);


--
-- Name: reviews_id_seq; Type: SEQUENCE SET; Schema: public; Owner: blog
--

SELECT pg_catalog.setval('public.reviews_id_seq', 2, true);


--
-- Name: roles_id_seq; Type: SEQUENCE SET; Schema: public; Owner: blog
--

SELECT pg_catalog.setval('public.roles_id_seq', 4, true);


--
-- Name: tags_id_seq; Type: SEQUENCE SET; Schema: public; Owner: blog
--

SELECT pg_catalog.setval('public.tags_id_seq', 9, true);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: blog
--

SELECT pg_catalog.setval('public.users_id_seq', 8, true);


--
-- Name: post_images post_images_pkey; Type: CONSTRAINT; Schema: public; Owner: blog
--

ALTER TABLE ONLY public.post_images
    ADD CONSTRAINT post_images_pkey PRIMARY KEY (id);


--
-- Name: post_tags post_tags_pkey; Type: CONSTRAINT; Schema: public; Owner: blog
--

ALTER TABLE ONLY public.post_tags
    ADD CONSTRAINT post_tags_pkey PRIMARY KEY (post_id, tag_id);


--
-- Name: posts posts_pkey; Type: CONSTRAINT; Schema: public; Owner: blog
--

ALTER TABLE ONLY public.posts
    ADD CONSTRAINT posts_pkey PRIMARY KEY (id);


--
-- Name: reviews reviews_pkey; Type: CONSTRAINT; Schema: public; Owner: blog
--

ALTER TABLE ONLY public.reviews
    ADD CONSTRAINT reviews_pkey PRIMARY KEY (id);


--
-- Name: roles roles_pkey; Type: CONSTRAINT; Schema: public; Owner: blog
--

ALTER TABLE ONLY public.roles
    ADD CONSTRAINT roles_pkey PRIMARY KEY (id);


--
-- Name: tags tags_pkey; Type: CONSTRAINT; Schema: public; Owner: blog
--

ALTER TABLE ONLY public.tags
    ADD CONSTRAINT tags_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: blog
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: post_tags fk_post_tags_post; Type: FK CONSTRAINT; Schema: public; Owner: blog
--

ALTER TABLE ONLY public.post_tags
    ADD CONSTRAINT fk_post_tags_post FOREIGN KEY (post_id) REFERENCES public.posts(id);


--
-- Name: post_tags fk_post_tags_tag; Type: FK CONSTRAINT; Schema: public; Owner: blog
--

ALTER TABLE ONLY public.post_tags
    ADD CONSTRAINT fk_post_tags_tag FOREIGN KEY (tag_id) REFERENCES public.tags(id);


--
-- Name: post_images fk_posts_post_images; Type: FK CONSTRAINT; Schema: public; Owner: blog
--

ALTER TABLE ONLY public.post_images
    ADD CONSTRAINT fk_posts_post_images FOREIGN KEY (post_id) REFERENCES public.posts(id);


--
-- Name: reviews fk_posts_reviews; Type: FK CONSTRAINT; Schema: public; Owner: blog
--

ALTER TABLE ONLY public.reviews
    ADD CONSTRAINT fk_posts_reviews FOREIGN KEY (post_id) REFERENCES public.posts(id);


--
-- Name: users fk_roles_user; Type: FK CONSTRAINT; Schema: public; Owner: blog
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT fk_roles_user FOREIGN KEY (role_id) REFERENCES public.roles(id);


--
-- Name: posts fk_users_posts; Type: FK CONSTRAINT; Schema: public; Owner: blog
--

ALTER TABLE ONLY public.posts
    ADD CONSTRAINT fk_users_posts FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: reviews fk_users_reviews; Type: FK CONSTRAINT; Schema: public; Owner: blog
--

ALTER TABLE ONLY public.reviews
    ADD CONSTRAINT fk_users_reviews FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- PostgreSQL database dump complete
--

--
-- PostgreSQL database cluster dump complete
--

