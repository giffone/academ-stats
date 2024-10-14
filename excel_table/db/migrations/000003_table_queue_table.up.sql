CREATE TABLE IF NOT EXISTS "excel_table"."table_queue" (
	value varchar(50) NOT NULL,
	value_title varchar(50) DEFAULT 'other'::character varying NOT NULL,
	queue int8 DEFAULT 1000 NOT NULL,
	value_type varchar(20) DEFAULT ''::character varying NOT NULL,
	CONSTRAINT table_queue_pkey PRIMARY KEY (value)
);

-- Permissions

ALTER TABLE "excel_table"."table_queue" OWNER TO postgres;
GRANT ALL ON TABLE "excel_table"."table_queue" TO postgres;
GRANT ALL ON TABLE "excel_table"."table_queue" TO excel_table;