CREATE TABLE dict (key UInt64, value1 UInt64, value2 Float64, value3 String,
        value4 String, value5 String, value6 String, value7 String, value8 String, value9 String,
        value10 String) ENGINE = Join(ANY, LEFT, key);