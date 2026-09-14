            CREATE USER `clickstack-ingest` IDENTIFIED WITH sha256_password BY '<password>';
            GRANT SELECT, INSERT, CREATE DATABASE, CREATE TABLE, CREATE VIEW ON default.* TO `clickstack-ingest`;
