
        INSERT INTO map_10000_keys
        SELECT number AS id, mapFromArrays(arrayMap(i -> 'key' || toString(i), range(10000)), arrayMap(i -> number + i, range(10000))) AS m
        FROM numbers(1000)
    