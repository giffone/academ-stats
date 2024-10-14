CREATE TABLE IF NOT EXISTS academ_stats.schema_migrations (
    version bigint PRIMARY KEY,
    dirty boolean NOT NULL
);

ALTER TABLE IF EXISTS academ_stats.schema_migrations OWNER to postgres;
GRANT ALL ON TABLE academ_stats.schema_migrations TO postgres;
GRANT ALL ON TABLE academ_stats.schema_migrations TO session_manager;