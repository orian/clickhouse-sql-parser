
        insert into function file('test_orc_tfp.orc', 'ORC') select * from generateRandom('tuple_column Tuple(a Nullable(String), b Nullable(Float64), c Nullable(Int64)), array_tuple_column Nested(a Nullable(String), b Nullable(Float64), c Nullable(Int64)), map_tuple_column Map(String, Tuple(a Nullable(String), b Nullable(Float64), c Nullable(Int64)))') limit 1000000
    