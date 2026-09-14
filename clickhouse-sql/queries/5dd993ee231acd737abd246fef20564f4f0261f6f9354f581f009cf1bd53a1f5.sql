
        CREATE TABLE t_chain_derived ENGINE = MergeTree ORDER BY tuple() AS
        SELECT
            repeat('p', 100) || 'a' || toString(number % 100) AS s1,
            repeat('p', 100) || 'b' || toString(number % 100) AS s2,
            repeat('p', 100) || 'c' || toString(number % 100) AS s3,
            repeat('p', 100) || 'd' || toString(number % 100) AS s4,
            if(number % 100 = 0, 'a', 'z') || repeat('q', 200) || toString(number) AS b,
            leftPad(toString(number % 100000), 8, '0') AS x,
            number AS k
        FROM numbers_mt(10000000)
    