INSERT INTO perf_nested_lc_functions
        SELECT
            number AS id,
            arrayMap(x -> concat('k', toString((id + x) % 128)), range(16)) AS arr_lc,
            mapFromArrays(
                arrayMap(x -> concat('k', toString((id + x) % 128)), range(16)),
                arrayMap(x -> concat('v', toString((id + x) % 128)), range(16))) AS map_key_lc,
            mapFromArrays(
                arrayMap(x -> concat('k', toString((id + x) % 128)), range(16)),
                arrayMap(x -> concat('v', toString((id + x) % 128)), range(16))) AS map_value_lc
        FROM numbers(1000000)
    