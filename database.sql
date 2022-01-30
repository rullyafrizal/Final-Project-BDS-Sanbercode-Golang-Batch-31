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
    is_published boolean DEFAULT false,
    user_id bigint,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    vote_count bigint,
    published_at timestamp with time zone
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
    comment text,
    user_id bigint,
    post_id bigint,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP
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
    avatar text,
    role_id bigint,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP
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
-- Name: votes; Type: TABLE; Schema: public; Owner: blog
--

CREATE TABLE public.votes (
    user_id bigint,
    post_id bigint,
    state bigint
);


ALTER TABLE public.votes OWNER TO blog;

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
2	https://picsum.photos/200/300	2
1	https://picsum.photos/200/300	\N
3	https://random.imagecdn.app/500/150	1
\.


--
-- Data for Name: post_tags; Type: TABLE DATA; Schema: public; Owner: blog
--

COPY public.post_tags (post_id, tag_id) FROM stdin;
1	3
2	1
2	4
1	5
\.


--
-- Data for Name: posts; Type: TABLE DATA; Schema: public; Owner: blog
--

COPY public.posts (id, title, content, is_published, user_id, created_at, updated_at, vote_count, published_at) FROM stdin;
2	Postingan dari John Doe kedua	Quisque velit nisi, pretium ut lacinia in, elementum id enim. Pellentesque in ipsum id orci porta dapibus.	f	2	2022-01-28 08:34:10.933547+00	2022-01-28 08:34:10.933547+00	0	\N
1	Ini adalah title	Quisque velit nisi, pretium ut lacinia in, elementum id enim. Pellentesque in ipsum id orci porta dapibus.	t	2	2022-01-28 01:50:08.956709+00	2022-01-28 09:19:23.920354+00	0	2022-01-28 08:29:39.70096+00
\.


--
-- Data for Name: reviews; Type: TABLE DATA; Schema: public; Owner: blog
--

COPY public.reviews (id, comment, user_id, post_id, created_at, updated_at) FROM stdin;
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
1	fashion
2	olahraga
3	misteri
4	pemrograman
5	sci-fi
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: blog
--

COPY public.users (id, name, email, password, avatar, role_id, created_at, updated_at) FROM stdin;
1	Admin	admin@admin.com	$2a$14$SCwlDi49s0bEBpDm2JQ5A.bu83b6e39NGgmy8aPwKOyNDvnXykzyq		1	2022-01-28 01:23:36.999529+00	2022-01-28 01:23:36.999529+00
3	Zidan	zidan@gmail.com	$2a$14$zBvngya6ZsHs5d8CiVpTUO4UFFbfC3y5O6aAkR1icqqEQBVsC2K4m		2	2022-01-28 02:27:05.798588+00	2022-01-28 02:27:05.798588+00
2	John Doe	johndoe@gmail.com	$2a$14$AZQf1gRDXSsuqwH93KDAIuHJL/jjNzjkPdowWkAYK3x7pBegAqYpG		2	2022-01-28 01:49:01.926589+00	2022-01-28 08:54:38.299737+00
\.


--
-- Data for Name: votes; Type: TABLE DATA; Schema: public; Owner: blog
--

COPY public.votes (user_id, post_id, state) FROM stdin;
2	1	1
3	1	-1
\.


--
-- Name: post_images_id_seq; Type: SEQUENCE SET; Schema: public; Owner: blog
--

SELECT pg_catalog.setval('public.post_images_id_seq', 3, true);


--
-- Name: posts_id_seq; Type: SEQUENCE SET; Schema: public; Owner: blog
--

SELECT pg_catalog.setval('public.posts_id_seq', 2, true);


--
-- Name: reviews_id_seq; Type: SEQUENCE SET; Schema: public; Owner: blog
--

SELECT pg_catalog.setval('public.reviews_id_seq', 1, false);


--
-- Name: roles_id_seq; Type: SEQUENCE SET; Schema: public; Owner: blog
--

SELECT pg_catalog.setval('public.roles_id_seq', 2, true);


--
-- Name: tags_id_seq; Type: SEQUENCE SET; Schema: public; Owner: blog
--

SELECT pg_catalog.setval('public.tags_id_seq', 5, true);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: blog
--

SELECT pg_catalog.setval('public.users_id_seq', 3, true);


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
-- Name: votes fk_votes_post; Type: FK CONSTRAINT; Schema: public; Owner: blog
--

ALTER TABLE ONLY public.votes
    ADD CONSTRAINT fk_votes_post FOREIGN KEY (post_id) REFERENCES public.posts(id);


--
-- Name: votes fk_votes_user; Type: FK CONSTRAINT; Schema: public; Owner: blog
--

ALTER TABLE ONLY public.votes
    ADD CONSTRAINT fk_votes_user FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- PostgreSQL database dump complete
--

