
        INSERT INTO map_1000_keys
        SELECT number AS id, mapFromArrays(arrayMap(i -> 'key' || toString(i), range(1000)), arrayMap(i -> number + i, range(1000))) AS m
        FROM numbers(100000)
    