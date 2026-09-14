
        CREATE TABLE limit_by_in_pk_order_merge
        (a Int32, b UInt32, z UInt32)
        ENGINE = Merge(currentDatabase(), '^limit_by_in_pk_order_merge_[12]$')
    