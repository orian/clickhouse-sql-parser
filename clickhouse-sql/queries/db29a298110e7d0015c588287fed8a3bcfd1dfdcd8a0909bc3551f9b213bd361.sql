
        INSERT INTO ttl_group_by_src
        SELECT intDiv(number, 500000), now() - INTERVAL 1 DAY,
               number, number + 1, number + 2, number + 3, number + 4, number + 5, number + 6, number + 7
        FROM numbers(3000000)
    