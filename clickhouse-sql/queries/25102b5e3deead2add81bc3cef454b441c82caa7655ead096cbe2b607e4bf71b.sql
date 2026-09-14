
CREATE TABLE squash_performance
(
    s1 String,
    s2 Nullable(String),
    a1 Array(Array(String)),
    a2 Array(Array(UInt32)),
    m1 Map(String, Array(String)),
    m2 Map(String, Array(UInt64)),
    t Tuple(String, Array(String), Map(String, String))
)
ENGINE = Null;
    