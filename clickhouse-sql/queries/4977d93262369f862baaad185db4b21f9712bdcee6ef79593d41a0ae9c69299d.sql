
        INSERT INTO source_table
        SELECT
        number AS id,
        if(number % 3 = 0, NULL, toInt8(number % 100)) AS col_int8,
        if(number % 4 = 0, NULL, toUInt8(number % 200)) AS col_uint8,
        if(number % 5 = 0, NULL, toInt16(number % 10000)) AS col_int16,
        if(number % 6 = 0, NULL, toUInt16(number % 20000)) AS col_uint16,
        if(number % 7 = 0, NULL, toInt32(number)) AS col_int32,
        if(number % 8 = 0, NULL, toUInt32(number)) AS col_uint32,
        if(number % 9 = 0, NULL, toInt64(number)) AS col_int64,
        if(number % 10 = 0, NULL, toUInt64(number)) AS col_uint64,
        if(number % 11 = 0, NULL, toFloat32(number / 10.0)) AS col_float32,
        if(number % 12 = 0, NULL, toFloat64(number / 100.0)) AS col_float64,
        if(number % 13 = 0, NULL, concat('String', toString(number % 1000))) AS col_string,
        if(number % 14 = 0, NULL, toFixedString(concat('F', toString(number % 100)), 10)) AS col_fixed_string,
        if(number % 15 = 0, NULL, toDate('1970-01-01') + number % 1000) AS col_date,
        if(number % 16 = 0, NULL, toDateTime('1970-01-01 00:00:00') + number % 100000) AS col_datetime,
        if(number % 17 = 0, NULL, toDateTime64('1970-01-01 00:00:00.000', 3) + number % 100000 / 1000.0) AS col_datetime64,
        if(number % 18 = 0, NULL, toDecimal32(number % 10000, 5)) AS col_decimal32,
        if(number % 19 = 0, NULL, toDecimal64(number % 1000000, 10)) AS col_decimal64,
        if(number % 20 = 0, NULL, toDecimal128(number % 1000000, 20)) AS col_decimal128,
        [if(number % 3 = 0, NULL, number % 10000), if(number % 5 = 0, NULL, (number % 10000)+1), if(number % 7 = 0, NULL, (number % 10000)+2)] AS col_array,
        if(number % 22 = 0, NULL, concat('LC', toString(number % 1000))) AS col_low_cardinality
        FROM numbers(1000000)
    