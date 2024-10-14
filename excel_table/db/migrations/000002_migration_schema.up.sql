CREATE TABLE IF NOT EXISTS excel_table.schema_migrations (
    version bigint PRIMARY KEY,
    dirty boolean NOT NULL
);

ALTER TABLE IF EXISTS excel_table.schema_migrations OWNER to postgres;
GRANT ALL ON TABLE excel_table.schema_migrations TO postgres;
GRANT ALL ON TABLE excel_table.schema_migrations TO excel_table;