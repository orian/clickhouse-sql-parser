
        ALTER TABLE t_materialize_uniq_v2_f64
        CLEAR STATISTICS ALL, MATERIALIZE STATISTICS ALL
        SETTINGS mutations_sync = 1
    