
        INSERT INTO variant_compact SELECT number,
            if(number % 2, number::UInt64::Variant(String, UInt64), randomString(1000)::String::Variant(String, UInt64))
        FROM numbers(800000)
    