CREATE TABLE fuse_tbl(a Int8, b Int8) Engine = Log;
SET optimize_syntax_fuse_functions = 1;
EXPLAIN SYNTAX run_query_tree_passes = 1 SELECT sum(a), sum(b), count(b), avg(b) from fuse_tbl FORMAT TSVRaw;
