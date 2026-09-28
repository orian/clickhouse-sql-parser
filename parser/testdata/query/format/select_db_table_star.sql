-- Origin SQL:
SELECT db.t.* FROM db.t;
SELECT db.t.* EXCEPT (a) FROM db.t;
SELECT db.t.* REPLACE(a + 1 AS a) APPLY(toString) FROM db.t;
SELECT db.t.*, t2.x FROM db.t, t2;
SELECT count(db.t.*) FROM db.t;


-- Format SQL:
SELECT db.t.* FROM db.t;
SELECT db.t.* EXCEPT (a) FROM db.t;
SELECT db.t.* REPLACE (a + 1 AS a) APPLY (toString) FROM db.t;
SELECT db.t.*, t2.x FROM db.t, t2;
SELECT count(db.t.*) FROM db.t;
