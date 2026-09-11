
        CREATE TABLE t_chain_derived_join ENGINE = MergeTree ORDER BY tuple() AS
        SELECT number AS k, number - (number % 2) AS v FROM numbers_mt(10000000)
    