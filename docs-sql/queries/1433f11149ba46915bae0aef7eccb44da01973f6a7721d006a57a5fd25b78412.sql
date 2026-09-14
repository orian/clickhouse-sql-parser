-- Filters reads of local_table on this server, including reads shipped by a Distributed table over it.
CREATE ROW POLICY filter ON mydb.local_table USING a < 1000 TO john;
