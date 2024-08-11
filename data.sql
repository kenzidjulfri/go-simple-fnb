--
-- PostgreSQL database dump
--

-- Dumped from database version 16.2
-- Dumped by pg_dump version 16.2

-- Started on 2024-08-12 02:47:39

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
-- TOC entry 4872 (class 1262 OID 38934)
-- Name: simple_fnb; Type: DATABASE; Schema: -; Owner: postgres
--

CREATE DATABASE simple_fnb WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'English_Indonesia.1252';


ALTER DATABASE simple_fnb OWNER TO postgres;

\connect simple_fnb

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
-- TOC entry 4 (class 2615 OID 2200)
-- Name: public; Type: SCHEMA; Schema: -; Owner: pg_database_owner
--

CREATE SCHEMA public;


ALTER SCHEMA public OWNER TO pg_database_owner;

--
-- TOC entry 4873 (class 0 OID 0)
-- Dependencies: 4
-- Name: SCHEMA public; Type: COMMENT; Schema: -; Owner: pg_database_owner
--

COMMENT ON SCHEMA public IS 'standard public schema';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 230 (class 1259 OID 39076)
-- Name: order_details; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.order_details (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    order_id bigint,
    item_id bigint,
    item_name character varying(255),
    item_type character varying(255),
    base_price numeric(10,2),
    qty smallint,
    sum_price numeric(10,2)
);


ALTER TABLE public.order_details OWNER TO postgres;

--
-- TOC entry 229 (class 1259 OID 39075)
-- Name: order_details_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.order_details_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.order_details_id_seq OWNER TO postgres;

--
-- TOC entry 4874 (class 0 OID 0)
-- Dependencies: 229
-- Name: order_details_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.order_details_id_seq OWNED BY public.order_details.id;


--
-- TOC entry 228 (class 1259 OID 39062)
-- Name: orders; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.orders (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    table_id bigint,
    is_closed boolean DEFAULT false
);


ALTER TABLE public.orders OWNER TO postgres;

--
-- TOC entry 227 (class 1259 OID 39061)
-- Name: orders_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.orders_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.orders_id_seq OWNER TO postgres;

--
-- TOC entry 4875 (class 0 OID 0)
-- Dependencies: 227
-- Name: orders_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.orders_id_seq OWNED BY public.orders.id;


--
-- TOC entry 224 (class 1259 OID 38977)
-- Name: printers; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.printers (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    code character varying(10),
    name character varying(255),
    usage character varying(255)
);


ALTER TABLE public.printers OWNER TO postgres;

--
-- TOC entry 223 (class 1259 OID 38976)
-- Name: printers_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.printers_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.printers_id_seq OWNER TO postgres;

--
-- TOC entry 4876 (class 0 OID 0)
-- Dependencies: 223
-- Name: printers_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.printers_id_seq OWNED BY public.printers.id;


--
-- TOC entry 218 (class 1259 OID 38944)
-- Name: products; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.products (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name character varying(255),
    category character varying(255),
    price numeric(10,2)
);


ALTER TABLE public.products OWNER TO postgres;

--
-- TOC entry 217 (class 1259 OID 38943)
-- Name: products_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.products_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.products_id_seq OWNER TO postgres;

--
-- TOC entry 4877 (class 0 OID 0)
-- Dependencies: 217
-- Name: products_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.products_id_seq OWNED BY public.products.id;


--
-- TOC entry 226 (class 1259 OID 39039)
-- Name: promo_details; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.promo_details (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    promo_id bigint,
    variant_id bigint,
    product_id bigint
);


ALTER TABLE public.promo_details OWNER TO postgres;

--
-- TOC entry 225 (class 1259 OID 39038)
-- Name: promo_details_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.promo_details_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.promo_details_id_seq OWNER TO postgres;

--
-- TOC entry 4878 (class 0 OID 0)
-- Dependencies: 225
-- Name: promo_details_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.promo_details_id_seq OWNED BY public.promo_details.id;


--
-- TOC entry 222 (class 1259 OID 38969)
-- Name: promos; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.promos (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name character varying(255),
    price numeric(10,2)
);


ALTER TABLE public.promos OWNER TO postgres;

--
-- TOC entry 221 (class 1259 OID 38968)
-- Name: promos_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.promos_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.promos_id_seq OWNER TO postgres;

--
-- TOC entry 4879 (class 0 OID 0)
-- Dependencies: 221
-- Name: promos_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.promos_id_seq OWNED BY public.promos.id;


--
-- TOC entry 216 (class 1259 OID 38936)
-- Name: tables; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.tables (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name character varying(255)
);


ALTER TABLE public.tables OWNER TO postgres;

--
-- TOC entry 215 (class 1259 OID 38935)
-- Name: tables_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.tables_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.tables_id_seq OWNER TO postgres;

--
-- TOC entry 4880 (class 0 OID 0)
-- Dependencies: 215
-- Name: tables_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.tables_id_seq OWNED BY public.tables.id;


--
-- TOC entry 220 (class 1259 OID 38954)
-- Name: variants; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.variants (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name character varying(255),
    product_id bigint,
    additional_price numeric
);


ALTER TABLE public.variants OWNER TO postgres;

--
-- TOC entry 219 (class 1259 OID 38953)
-- Name: variants_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.variants_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.variants_id_seq OWNER TO postgres;

--
-- TOC entry 4881 (class 0 OID 0)
-- Dependencies: 219
-- Name: variants_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.variants_id_seq OWNED BY public.variants.id;


--
-- TOC entry 4677 (class 2604 OID 39079)
-- Name: order_details id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.order_details ALTER COLUMN id SET DEFAULT nextval('public.order_details_id_seq'::regclass);


--
-- TOC entry 4675 (class 2604 OID 39065)
-- Name: orders id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.orders ALTER COLUMN id SET DEFAULT nextval('public.orders_id_seq'::regclass);


--
-- TOC entry 4673 (class 2604 OID 38980)
-- Name: printers id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.printers ALTER COLUMN id SET DEFAULT nextval('public.printers_id_seq'::regclass);


--
-- TOC entry 4670 (class 2604 OID 38947)
-- Name: products id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.products ALTER COLUMN id SET DEFAULT nextval('public.products_id_seq'::regclass);


--
-- TOC entry 4674 (class 2604 OID 39042)
-- Name: promo_details id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.promo_details ALTER COLUMN id SET DEFAULT nextval('public.promo_details_id_seq'::regclass);


--
-- TOC entry 4672 (class 2604 OID 38972)
-- Name: promos id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.promos ALTER COLUMN id SET DEFAULT nextval('public.promos_id_seq'::regclass);


--
-- TOC entry 4669 (class 2604 OID 38939)
-- Name: tables id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tables ALTER COLUMN id SET DEFAULT nextval('public.tables_id_seq'::regclass);


--
-- TOC entry 4671 (class 2604 OID 38957)
-- Name: variants id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.variants ALTER COLUMN id SET DEFAULT nextval('public.variants_id_seq'::regclass);


--
-- TOC entry 4866 (class 0 OID 39076)
-- Dependencies: 230
-- Data for Name: order_details; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.order_details VALUES (1, NULL, NULL, NULL, 1, 2, 'Jeruk Panas', 'VARIANT', 10000.00, 1, 10000.00);
INSERT INTO public.order_details VALUES (2, NULL, NULL, NULL, 1, 6, 'Kopi Panas', 'VARIANT', 6000.00, 1, 6000.00);
INSERT INTO public.order_details VALUES (3, NULL, NULL, NULL, 1, 1, 'Nasi Goreng + Jeruk Dingin', 'PROMO', 23000.00, 2, 46000.00);
INSERT INTO public.order_details VALUES (4, NULL, NULL, NULL, 1, 3, 'Teh Manis', 'VARIANT', 8000.00, 1, 8000.00);
INSERT INTO public.order_details VALUES (5, NULL, NULL, NULL, 1, 7, 'Mie Goreng', 'VARIANT', 15000.00, 1, 15000.00);


--
-- TOC entry 4864 (class 0 OID 39062)
-- Dependencies: 228
-- Data for Name: orders; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.orders VALUES (1, NULL, NULL, NULL, 1, false);


--
-- TOC entry 4860 (class 0 OID 38977)
-- Dependencies: 224
-- Data for Name: printers; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.printers VALUES (1, NULL, NULL, NULL, 'A', 'Printer Kasir', 'BILL');
INSERT INTO public.printers VALUES (2, NULL, NULL, NULL, 'B', 'Printer Dapur (Makanan)', 'FOOD');
INSERT INTO public.printers VALUES (3, NULL, NULL, NULL, 'C', 'Printer Bar (Minuman)', 'BEVERAGE');


--
-- TOC entry 4854 (class 0 OID 38944)
-- Dependencies: 218
-- Data for Name: products; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.products VALUES (1, NULL, NULL, NULL, 'Jeruk', 'BEVERAGE', 10000.00);
INSERT INTO public.products VALUES (2, NULL, NULL, NULL, 'Teh', 'BEVERAGE', 5000.00);
INSERT INTO public.products VALUES (3, NULL, NULL, NULL, 'Kopi', 'BEVERAGE', 6000.00);
INSERT INTO public.products VALUES (4, NULL, NULL, NULL, 'Extra Es Batu', 'BEVERAGE', 2000.00);
INSERT INTO public.products VALUES (5, NULL, NULL, NULL, 'Mie', 'FOOD', 15000.00);
INSERT INTO public.products VALUES (6, NULL, NULL, NULL, 'Nasi Goreng', 'FOOD', 15000.00);


--
-- TOC entry 4862 (class 0 OID 39039)
-- Dependencies: 226
-- Data for Name: promo_details; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.promo_details VALUES (1, NULL, NULL, NULL, 1, NULL, 6);
INSERT INTO public.promo_details VALUES (2, NULL, NULL, NULL, 1, 1, 1);


--
-- TOC entry 4858 (class 0 OID 38969)
-- Dependencies: 222
-- Data for Name: promos; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.promos VALUES (1, NULL, NULL, NULL, 'Nasi Goreng + Jeruk Dingin', 23000.00);


--
-- TOC entry 4852 (class 0 OID 38936)
-- Dependencies: 216
-- Data for Name: tables; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.tables VALUES (1, NULL, NULL, NULL, 'MEJA NO 1');
INSERT INTO public.tables VALUES (2, NULL, NULL, NULL, 'MEJA NO 2');
INSERT INTO public.tables VALUES (3, NULL, NULL, NULL, 'MEJA NO 3');


--
-- TOC entry 4856 (class 0 OID 38954)
-- Dependencies: 220
-- Data for Name: variants; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.variants VALUES (1, NULL, NULL, NULL, 'Jeruk Dingin', 1, 2000);
INSERT INTO public.variants VALUES (2, NULL, NULL, NULL, 'Jeruk Panas', 1, 0);
INSERT INTO public.variants VALUES (3, NULL, NULL, NULL, 'Teh Manis', 2, 3000);
INSERT INTO public.variants VALUES (4, NULL, NULL, NULL, 'Teh Tawar', 2, 0);
INSERT INTO public.variants VALUES (5, NULL, NULL, NULL, 'Kopi Dingin', 3, 2000);
INSERT INTO public.variants VALUES (6, NULL, NULL, NULL, 'Kopi Panas', 3, 0);
INSERT INTO public.variants VALUES (7, NULL, NULL, NULL, 'Mie Goreng', 5, 0);
INSERT INTO public.variants VALUES (8, NULL, NULL, NULL, 'Mie Kuah', 5, 0);


--
-- TOC entry 4882 (class 0 OID 0)
-- Dependencies: 229
-- Name: order_details_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.order_details_id_seq', 5, true);


--
-- TOC entry 4883 (class 0 OID 0)
-- Dependencies: 227
-- Name: orders_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.orders_id_seq', 1, true);


--
-- TOC entry 4884 (class 0 OID 0)
-- Dependencies: 223
-- Name: printers_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.printers_id_seq', 3, true);


--
-- TOC entry 4885 (class 0 OID 0)
-- Dependencies: 217
-- Name: products_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.products_id_seq', 6, true);


--
-- TOC entry 4886 (class 0 OID 0)
-- Dependencies: 225
-- Name: promo_details_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.promo_details_id_seq', 2, true);


--
-- TOC entry 4887 (class 0 OID 0)
-- Dependencies: 221
-- Name: promos_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.promos_id_seq', 1, true);


--
-- TOC entry 4888 (class 0 OID 0)
-- Dependencies: 215
-- Name: tables_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.tables_id_seq', 3, true);


--
-- TOC entry 4889 (class 0 OID 0)
-- Dependencies: 219
-- Name: variants_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.variants_id_seq', 8, true);


--
-- TOC entry 4701 (class 2606 OID 39083)
-- Name: order_details order_details_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.order_details
    ADD CONSTRAINT order_details_pkey PRIMARY KEY (id);


--
-- TOC entry 4698 (class 2606 OID 39068)
-- Name: orders orders_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.orders
    ADD CONSTRAINT orders_pkey PRIMARY KEY (id);


--
-- TOC entry 4692 (class 2606 OID 38984)
-- Name: printers printers_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.printers
    ADD CONSTRAINT printers_pkey PRIMARY KEY (id);


--
-- TOC entry 4683 (class 2606 OID 38951)
-- Name: products products_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT products_pkey PRIMARY KEY (id);


--
-- TOC entry 4695 (class 2606 OID 39044)
-- Name: promo_details promo_details_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.promo_details
    ADD CONSTRAINT promo_details_pkey PRIMARY KEY (id);


--
-- TOC entry 4689 (class 2606 OID 38974)
-- Name: promos promos_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.promos
    ADD CONSTRAINT promos_pkey PRIMARY KEY (id);


--
-- TOC entry 4680 (class 2606 OID 38941)
-- Name: tables tables_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tables
    ADD CONSTRAINT tables_pkey PRIMARY KEY (id);


--
-- TOC entry 4686 (class 2606 OID 38961)
-- Name: variants variants_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.variants
    ADD CONSTRAINT variants_pkey PRIMARY KEY (id);


--
-- TOC entry 4699 (class 1259 OID 39089)
-- Name: idx_order_details_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_order_details_deleted_at ON public.order_details USING btree (deleted_at);


--
-- TOC entry 4696 (class 1259 OID 39074)
-- Name: idx_orders_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_orders_deleted_at ON public.orders USING btree (deleted_at);


--
-- TOC entry 4690 (class 1259 OID 38985)
-- Name: idx_printers_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_printers_deleted_at ON public.printers USING btree (deleted_at);


--
-- TOC entry 4681 (class 1259 OID 38952)
-- Name: idx_products_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_products_deleted_at ON public.products USING btree (deleted_at);


--
-- TOC entry 4693 (class 1259 OID 39060)
-- Name: idx_promo_details_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_promo_details_deleted_at ON public.promo_details USING btree (deleted_at);


--
-- TOC entry 4687 (class 1259 OID 38975)
-- Name: idx_promos_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_promos_deleted_at ON public.promos USING btree (deleted_at);


--
-- TOC entry 4678 (class 1259 OID 38942)
-- Name: idx_tables_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_tables_deleted_at ON public.tables USING btree (deleted_at);


--
-- TOC entry 4684 (class 1259 OID 38967)
-- Name: idx_variants_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_variants_deleted_at ON public.variants USING btree (deleted_at);


--
-- TOC entry 4707 (class 2606 OID 39084)
-- Name: order_details fk_orders_details; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.order_details
    ADD CONSTRAINT fk_orders_details FOREIGN KEY (order_id) REFERENCES public.orders(id);


--
-- TOC entry 4706 (class 2606 OID 39069)
-- Name: orders fk_orders_table; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.orders
    ADD CONSTRAINT fk_orders_table FOREIGN KEY (table_id) REFERENCES public.tables(id);


--
-- TOC entry 4703 (class 2606 OID 39050)
-- Name: promo_details fk_promo_details_product; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.promo_details
    ADD CONSTRAINT fk_promo_details_product FOREIGN KEY (product_id) REFERENCES public.products(id);


--
-- TOC entry 4704 (class 2606 OID 39045)
-- Name: promo_details fk_promo_details_variant; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.promo_details
    ADD CONSTRAINT fk_promo_details_variant FOREIGN KEY (variant_id) REFERENCES public.variants(id);


--
-- TOC entry 4705 (class 2606 OID 39055)
-- Name: promo_details fk_promos_details; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.promo_details
    ADD CONSTRAINT fk_promos_details FOREIGN KEY (promo_id) REFERENCES public.promos(id);


--
-- TOC entry 4702 (class 2606 OID 38962)
-- Name: variants fk_variants_product; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.variants
    ADD CONSTRAINT fk_variants_product FOREIGN KEY (product_id) REFERENCES public.products(id);


-- Completed on 2024-08-12 02:47:39

--
-- PostgreSQL database dump complete
--

