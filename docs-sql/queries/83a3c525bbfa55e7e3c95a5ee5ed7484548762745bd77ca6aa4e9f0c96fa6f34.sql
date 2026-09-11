SELECT * \
FROM  \
( \
    SELECT 1 AS x \
    UNION ALL \
    SELECT 1 \
    UNION DISTINCT \
    SELECT 3 \
)
