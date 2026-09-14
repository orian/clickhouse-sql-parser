
        WITH (rand64() % 2_000_000, toString(rand64() % 2_000_000)) as key
        SELECT dictHas('default.complex_key_{layout_suffix}_dictionary_s{shards}', key)
        FROM numbers(2_000_000)
        FORMAT Null
    