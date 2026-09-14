
        WITH (range(10), range(10))::Map(String, UInt64) AS m1, (range(3), range(3))::Map(String, UInt64) AS m2
        SELECT count() FROM numbers(500000)
        WHERE NOT ignore(mapUpdate(materialize(m1), materialize(m2)))
    