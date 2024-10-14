CREATE TABLE IF NOT EXISTS academ_stats.snapshot (
    id serial NOT NULL,
	stat_name varchar(20) NOT NULL,
	body jsonb NULL,
	createdAt timestamptz DEFAULT LOCALTIMESTAMP NOT NULL,
	CONSTRAINT id_pkey PRIMARY KEY (id)
);

ALTER TABLE IF EXISTS academ_stats.snapshot OWNER TO postgres;
GRANT ALL ON TABLE academ_stats.snapshot TO postgres;
GRANT ALL ON TABLE academ_stats.snapshot TO session_manager;