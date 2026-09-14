
        INSERT INTO variant_wide SELECT number, intDiv(number, 2048) % 2,
            if(number % 2, number::UInt64::Variant(String, UInt64), randomString(1000)::String::Variant(String, UInt64))
        FROM numbers(1500000)
    