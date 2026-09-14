CREATE DATABASE postgres_db
ENGINE = MaterializedPostgreSQL('postgres-host:5432', 'postgres_database', 'postgres_user', 'postgres_password',
                                sslmode = 'verify-full', sslrootcert_pem = '-----BEGIN CERTIFICATE-----
...
-----END CERTIFICATE-----');
