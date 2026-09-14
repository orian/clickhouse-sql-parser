
    create table t(
        f1 Nullable(String),
        f2 Nullable(Int32),
        f3 Nullable(String),
        f4 Nullable(Int32),
        f5 Nullable(UInt64),
        f6 Nullable(UInt64),
        f7 Nullable(String),
        f8 Nullable(String)
    ) ENGINE = File(JSONEachRow)