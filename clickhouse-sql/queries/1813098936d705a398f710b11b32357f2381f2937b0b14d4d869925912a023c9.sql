
CREATE TABLE column_map_test
ENGINE = MergeTree ORDER BY number
AS
SELECT number, map
FROM
(
    SELECT
        number,
        arrayMap(x -> toString(x), range(100)) AS keys,
        arrayMap(x -> toString(x * x), range(100)) AS values,
        cast((keys, values), 'Map(String, String)') AS map
    FROM numbers_mt(10000)
)
SETTINGS max_insert_threads = 8
    