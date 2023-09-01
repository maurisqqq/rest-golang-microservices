-- TOC entry 209 (class 1259 OID 19397)
-- Name: hargas; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.hargas (
    harga_id text NOT NULL,
    admin_id character varying(255) NOT NULL,
    harga_topup bigint NOT NULL,
    harga_buyback bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


--
-- TOC entry 211 (class 1259 OID 19413)
-- Name: rekenings; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.rekenings (
    rekening_id text NOT NULL,
    norek character varying(255) NOT NULL,
    saldo numeric NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


--
-- TOC entry 210 (class 1259 OID 19405)
-- Name: transaksis; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.transaksis (
    transaski_id text NOT NULL,
    type character varying(255) NOT NULL,
    gram numeric NOT NULL,
    harga_topup bigint NOT NULL,
    harga_buyback bigint NOT NULL,
    rekening_id character varying(255) NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


--
-- TOC entry 3363 (class 0 OID 19397)
-- Dependencies: 209
-- Data for Name: hargas; Type: TABLE DATA; Schema: public; Owner: -
--

INSERT INTO public.hargas VALUES ('oksoakq', 'a001', 120000, 120000, '2023-09-01 13:28:08.775+07', '2023-09-01 13:28:08.775372+07', NULL);


--
-- TOC entry 3365 (class 0 OID 19413)
-- Dependencies: 211
-- Data for Name: rekenings; Type: TABLE DATA; Schema: public; Owner: -
--

INSERT INTO public.rekenings VALUES ('oksokaok', 'r001', 0.013, '2023-09-01 13:28:08.775+07', '2023-09-01 14:48:43.091948+07', NULL);


--
-- TOC entry 3364 (class 0 OID 19405)
-- Dependencies: 210
-- Data for Name: transaksis; Type: TABLE DATA; Schema: public; Owner: -
--



--
-- TOC entry 3215 (class 2606 OID 19403)
-- Name: hargas hargas_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.hargas
    ADD CONSTRAINT hargas_pkey PRIMARY KEY (harga_id);


--
-- TOC entry 3222 (class 2606 OID 19419)
-- Name: rekenings rekenings_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.rekenings
    ADD CONSTRAINT rekenings_pkey PRIMARY KEY (rekening_id);


--
-- TOC entry 3219 (class 2606 OID 19411)
-- Name: transaksis transaksis_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.transaksis
    ADD CONSTRAINT transaksis_pkey PRIMARY KEY (transaski_id);


--
-- TOC entry 3216 (class 1259 OID 19404)
-- Name: idx_hargas_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_hargas_deleted_at ON public.hargas USING btree (deleted_at);


--
-- TOC entry 3220 (class 1259 OID 19420)
-- Name: idx_rekenings_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_rekenings_deleted_at ON public.rekenings USING btree (deleted_at);


--
-- TOC entry 3217 (class 1259 OID 19412)
-- Name: idx_transaksis_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_transaksis_deleted_at ON public.transaksis USING btree (deleted_at);


--
-- TOC entry 3223 (class 2606 OID 19421)
-- Name: transaksis fk_rekenings_transaksi; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.transaksis
    ADD CONSTRAINT fk_rekenings_transaksi FOREIGN KEY (rekening_id) REFERENCES public.rekenings(rekening_id) ON UPDATE CASCADE ON DELETE CASCADE;


-- Completed on 2023-09-01 14:56:55 WIB

--
-- PostgreSQL database dump complete
--

