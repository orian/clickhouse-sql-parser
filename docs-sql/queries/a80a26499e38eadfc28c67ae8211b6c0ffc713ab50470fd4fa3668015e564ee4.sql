   -- Target database access: CREATE TABLE, DROP TABLE, SELECT, INSERT, TRUNCATE, etc.
   GRANT ALL ON <database>.* TO <user>;

   -- System table access for metadata discovery and partition management.
   -- These are NOT covered by the database grant above.
   GRANT SELECT ON system.columns TO <user>;
   GRANT SELECT ON system.parts TO <user>;
   GRANT SELECT ON system.tables TO <user>;
