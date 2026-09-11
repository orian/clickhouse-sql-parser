CREATE TABLE postgresql_db.postgresql_replica (key UInt64, value UInt64)
ENGINE = MaterializedPostgreSQL('postgres1:5432', 'postgres_database', 'postgresql_table', 'postgres_user', 'postgres_password',
                                sslmode = 'verify-full', sslrootcert_pem = '-----BEGIN CERTIFICATE-----
...
-----END CERTIFICATE-----')
PRIMARY KEY key;
