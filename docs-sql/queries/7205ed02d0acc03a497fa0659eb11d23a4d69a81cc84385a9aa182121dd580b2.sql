CREATE TABLE t_enum_nullable
(
    x Nullable(Enum8('hello' = 1, 'world' = 2))
)
ENGINE = TinyLog
