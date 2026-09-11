
        INSERT INTO map_100_keys
        SELECT number AS id, mapFromArrays(arrayMap(i -> 'key' || toString(i), range(100)), arrayMap(i -> number + i, range(100))) AS m
        FROM numbers(1000000)
    