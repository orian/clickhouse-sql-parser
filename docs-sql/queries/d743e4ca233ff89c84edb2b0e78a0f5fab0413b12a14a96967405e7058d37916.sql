CREATE TABLE tab
(
    a UInt64,
    b UInt64,
    INDEX idx_a a TYPE minmax,
    INDEX idx_b b TYPE set(3)
)
ENGINE = MergeTree ORDER BY tuple();

SET exclude_materialize_skip_indexes_on_insert='idx_a'; -- idx_a will be not be updated upon insert
--SET exclude_materialize_skip_indexes_on_insert='idx_a, idx_b'; -- neither index would be updated on insert

INSERT INTO tab SELECT number, number / 50 FROM numbers(100); -- only idx_b is updated

-- since it is a session setting it can be set on a per-query level
INSERT INTO tab SELECT number, number / 50 FROM numbers(100, 100) SETTINGS exclude_materialize_skip_indexes_on_insert='idx_b';

ALTER TABLE tab MATERIALIZE INDEX idx_a; -- this query can be used to explicitly materialize the index

SET exclude_materialize_skip_indexes_on_insert = DEFAULT; -- reset setting to default
