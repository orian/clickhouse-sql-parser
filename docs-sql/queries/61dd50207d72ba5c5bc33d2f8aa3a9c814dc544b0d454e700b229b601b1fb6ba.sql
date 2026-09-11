CREATE TABLE table_1
(
    `id` UInt32,
    `value` String
)
ENGINE = MergeTree
ORDER BY id;

CREATE TABLE table_2
(
    `id` UInt32,
    `value` String
)
ENGINE = MergeTree
ORDER BY id;

INSERT INTO table_1 VALUES (1, 'a'), (2, 'b'), (3, 'c');
INSERT INTO table_2 VALUES (1, 'x'), (2, 'y'), (3, 'z');
