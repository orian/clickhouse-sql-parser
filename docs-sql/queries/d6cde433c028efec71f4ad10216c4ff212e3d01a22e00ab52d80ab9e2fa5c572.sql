CREATE TABLE postgres_table (id UInt64, value String)
ENGINE = PostgreSQL('localhost:5432', 'database', 'table', 'user', 'password',
                    sslmode = 'verify-full', sslrootcert_pem = '-----BEGIN CERTIFICATE-----
...
-----END CERTIFICATE-----');
