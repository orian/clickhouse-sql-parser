  GRANT USAGE ON SCHEMA <schema_i> TO subscriber_user;
  GRANT CREATE ON DATABASE destination_db TO subscriber_user;
  GRANT pg_create_subscription TO subscriber_user;

  -- Accorder des droits sur les tables
  GRANT INSERT, UPDATE, DELETE, TRUNCATE ON ALL TABLES IN SCHEMA <schema_i> TO subscriber_user;
