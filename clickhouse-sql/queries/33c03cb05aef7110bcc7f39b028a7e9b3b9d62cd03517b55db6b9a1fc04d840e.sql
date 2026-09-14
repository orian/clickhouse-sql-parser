
        INSERT INTO minmax_insert
        SELECT
            toUInt8(number), toUInt8(number), toUInt8(number), toUInt8(number), toUInt8(number),
            toUInt16(number), toUInt16(number), toUInt16(number), toUInt16(number), toUInt16(number),
            toUInt32(number), toUInt32(number), toUInt32(number), toUInt32(number), toUInt32(number),
            number, number, number, number, number,
            toInt8(number), toInt8(number), toInt8(number), toInt8(number), toInt8(number),
            toInt16(number), toInt16(number), toInt16(number), toInt16(number), toInt16(number),
            toInt32(number), toInt32(number), toInt32(number), toInt32(number), toInt32(number),
            toInt64(number), toInt64(number), toInt64(number), toInt64(number), toInt64(number),
            toFloat32(number), toFloat32(number), toFloat32(number), toFloat32(number), toFloat32(number),
            toFloat64(number), toFloat64(number), toFloat64(number), toFloat64(number), toFloat64(number)
        FROM numbers(1000000)
        SETTINGS insert_deduplicate = 0, max_insert_threads = 1, min_insert_block_size_rows = 0, min_insert_block_size_bytes = 0
    