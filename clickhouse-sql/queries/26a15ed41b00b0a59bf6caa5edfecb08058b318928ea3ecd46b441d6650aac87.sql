
        create table test_orc_fpd
        (
            a Nullable(Int64),
            b Nullable(String)
        ) Engine=File(ORC)
    