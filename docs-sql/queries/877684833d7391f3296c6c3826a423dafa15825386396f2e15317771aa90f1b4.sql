CREATE TABLE tab
(
a UInt64,
b UInt64,
INDEX idx_a a TYPE minmax,
INDEX idx_b b TYPE set(3)
)
ENGINE = MergeTree ORDER BY tuple() SETTINGS exclude_materialize_skip_indexes_on_merge = 'idx_a';

INSERT INTO tab SELECT number, number / 50 FROM numbers(100); -- setting has no effect on INSERTs

-- idx_a will be excluded from update during background or explicit merge via OPTIMIZE TABLE FINAL

-- can exclude multiple indexes by providing a list
ALTER TABLE tab MODIFY SETTING exclude_materialize_skip_indexes_on_merge = 'idx_a, idx_b';

-- default setting, no indexes excluded from being updated during merge
ALTER TABLE tab MODIFY SETTING exclude_materialize_skip_indexes_on_merge = '';
