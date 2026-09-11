CREATE TABLE sample_oid
(
    _id String,
    another_oid_column String
) ENGINE = MongoDB('mongodb://user:pass@host/db', 'sample_oid');

SELECT count() FROM sample_oid WHERE _id = '67bf6cc44ebc466d33d42fb2'; --将输出 1。
SELECT count() FROM sample_oid WHERE another_oid_column = '67bf6cc40000000000ea41b1'; --将输出 0
