CREATE TABLE test
(
    `id` UInt8,
    `tag` Array(String),
    `version` Array(String)
)
ENGINE = MergeTree
ORDER BY id;

INSERT INTO test VALUES (1, ['release-stable', 'dev', 'security'], ['2.4.0', '2.6.0-alpha', '2.4.0-sec1']);

SELECT
    id,
    tag,
    version,
    seq
FROM test
ARRAY JOIN
    tag,
    version,
    arrayEnumerate(tag) AS seq
