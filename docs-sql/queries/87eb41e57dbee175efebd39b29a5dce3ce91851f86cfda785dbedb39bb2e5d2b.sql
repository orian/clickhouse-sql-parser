CREATE ROLE read_only_role;
GRANT SELECT ON db1.* TO read_only_role;

CREATE ROLE maint_role;
GRANT SELECT, INSERT, ALTER on db1.* TO maint_role;

GRANT read_only_role, maint_role TO service_user;
