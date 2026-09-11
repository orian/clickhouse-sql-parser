
        INSERT INTO t_in_clustered
        SELECT
            intDiv(number, 20000) AS k,
            reinterpretAsUUID(sipHash128(k)) AS u,
            number AS v
        FROM numbers_mt(40000000)
    