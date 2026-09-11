
        INSERT INTO dynamic_compact_read SELECT
            concat('k-', toString(number))::Dynamic AS key,
            concat(repeat('x', 100), '-', toString(number))::Dynamic AS blob1,
            concat(repeat('y', 100), '-', toString(number % 500))::Dynamic AS blob2,
            concat('K', toString(number % 20)) AS grp_kind,
            concat('t', toString(number % 200)) AS tenant,
            false AS flag_a,
            true AS flag_b,
            concat('r-', toString(number)) AS rid
        FROM numbers(2000000)
    