CREATE USER clickpipes_user PASSWORD '<password>';

GRANT USAGE ON SCHEMA public TO clickpipes_user;
GRANT SELECT ON ALL TABLES IN SCHEMA public TO clickpipes_user;
ALTER DEFAULT PRIVILEGES IN SCHEMA public
GRANT SELECT ON TABLES TO clickpipes_user;
ALTER USER clickpipes_user WITH REPLICATION;

