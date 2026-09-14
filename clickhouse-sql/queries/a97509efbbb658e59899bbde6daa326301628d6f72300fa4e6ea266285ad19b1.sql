
        CREATE TABLE sum_map_{scale} ENGINE Memory AS
        SELECT
            arrayMap(x -> (x % 23), range(50)) AS key,
            arrayMap(x -> intDiv(number, x + 1), range(50)) AS val
        FROM numbers({scale})
    