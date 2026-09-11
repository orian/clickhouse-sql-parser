CREATE ROLE IF NOT EXISTS sql_console_developer;

GRANT SELECT, INSERT, CREATE TABLE
ON my_database.*
TO sql_console_developer;
