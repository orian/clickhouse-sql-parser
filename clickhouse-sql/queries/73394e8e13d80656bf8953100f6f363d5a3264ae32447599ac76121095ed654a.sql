
        CREATE TABLE nums ENGINE Memory AS
        WITH number as x
        SELECT
            toUInt8(x)   u8,
            toUInt16(x)  u16,
            toUInt32(x)  u32,
            toUInt64(x)  u64,
            toInt8(x)    i8,
            toInt16(x)   i16,
            toInt32(x)   i32,
            toInt64(x)   i64,
            toFloat32(x) f32,
            toFloat64(x) f64
        FROM numbers_mt(200000000)
        SETTINGS max_threads = 4
        ;
    