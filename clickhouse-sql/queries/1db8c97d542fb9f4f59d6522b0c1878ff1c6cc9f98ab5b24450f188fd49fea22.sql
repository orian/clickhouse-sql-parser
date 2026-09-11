
        ALTER TABLE t_materialize_uniq_v2_i64
        CLEAR STATISTICS ALL, MATERIALIZE STATISTICS ALL
        SETTINGS mutations_sync = 1
    