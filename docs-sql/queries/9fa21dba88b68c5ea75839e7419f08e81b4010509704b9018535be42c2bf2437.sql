CREATE TABLE tab (
    ordinary_col UInt32,
    nullable_col Nullable(UInt32)
)
ENGINE = MergeTree
ORDER BY tuple();
INSERT INTO tab (ordinary_col, nullable_col) VALUES (1,1), (2, 2), (3,3);
SELECT isNullable(ordinary_col), isNullable(nullable_col) FROM tab;
