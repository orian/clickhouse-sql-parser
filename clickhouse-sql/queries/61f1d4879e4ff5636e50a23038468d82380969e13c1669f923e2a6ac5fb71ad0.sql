
        INSERT INTO array_join_filter_fusion
        SELECT number, arrayMap(x -> concat('e', toString(x % 40)), range(40)), repeat('x', 256)
        FROM numbers(3000000)
    