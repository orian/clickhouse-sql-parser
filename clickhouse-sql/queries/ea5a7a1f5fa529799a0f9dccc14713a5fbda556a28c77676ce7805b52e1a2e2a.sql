
        CREATE TABLE t_low_card_str
        (
            s8 String,
            s16 String,
            s24 String,
            s32 String
        )
        ENGINE = MergeTree
        ORDER BY tuple()
    