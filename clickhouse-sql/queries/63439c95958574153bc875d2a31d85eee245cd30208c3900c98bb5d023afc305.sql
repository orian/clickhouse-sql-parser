
        CREATE TABLE IF NOT EXISTS phj_asof_right (
            key UInt64,
            ts Nullable(DateTime),
            val String
        ) ENGINE = Memory
    