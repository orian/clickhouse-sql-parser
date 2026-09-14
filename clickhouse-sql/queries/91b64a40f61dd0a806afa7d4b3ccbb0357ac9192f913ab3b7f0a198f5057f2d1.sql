
        WITH (range(100), range(100))::Map(String, UInt64) AS m1, (range(30), range(30))::Map(String, UInt64) AS m2
        SELECT count() FROM numbers(50000)
        WHERE NOT ignore(mapUpdate(materialize(m1), m2))
    