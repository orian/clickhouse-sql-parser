SELECT * FROM {tbl:Identifier};
SELECT * FROM {db:Identifier}.{tbl:Identifier};
SELECT * FROM db.{tbl:Identifier} AS x;
SELECT * FROM t JOIN {u:Identifier} USING (x);
SELECT {col:Identifier} FROM t;
INSERT INTO {db:Identifier}.t VALUES (1);
USE {db:Identifier};
