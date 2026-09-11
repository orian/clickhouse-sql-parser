
        INSERT INTO map_subcol_pw_compact
        SELECT number AS id,
               mapFromArrays(
                   arrayMap(i -> 'k' || toString(i), range(16)),
                   arrayMap(i -> toString(number + i), range(16))) AS tags
        FROM numbers_mt(5000000)
    