
        WITH rand64() % 3_000_000 as key
        SELECT dictHas('default.simple_key_{layout_suffix}_dictionary_s{shards}', key)
        FROM numbers(3_000_000)
        FORMAT Null
    