
    INSERT INTO tab
    SELECT
        number,
        trimRight(if(number < 100000,
            repeat('clickhouse is a fast column oriented database system ', 4),
            repeat('the quick brown fox jumps over the lazy dog near the river ', 4)))
    FROM numbers(10000000)
  