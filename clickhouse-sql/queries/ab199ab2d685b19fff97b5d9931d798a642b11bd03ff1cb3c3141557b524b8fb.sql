
        CREATE TABLE projection_materialize (
            key UInt64,
            val UInt64
        ) ENGINE = MergeTree()
        ORDER BY key
    