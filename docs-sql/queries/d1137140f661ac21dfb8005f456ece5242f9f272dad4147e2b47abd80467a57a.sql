CREATE TABLE test.test_parquet
(
    `f_tinyint` Int8,
    `f_smallint` Int16,
    `f_int` Int32,
    `f_integer` Int32,
    `f_bigint` Int64,
    `f_float` Float32,
    `f_double` Float64,
    `f_decimal` Float64,
    `f_timestamp` DateTime,
    `f_date` Date,
    `f_string` String,
    `f_varchar` String,
    `f_char` String,
    `f_bool` Bool,
    `f_binary` String,
    `f_array_int` Array(Int32),
    `f_array_string` Array(String),
    `f_array_float` Array(Float32),
    `f_array_array_int` Array(Array(Int32)),
    `f_array_array_string` Array(Array(String)),
    `f_array_array_float` Array(Array(Float32)),
    `day` String
)
ENGINE = Hive('thrift://localhost:9083', 'test', 'test_parquet')
PARTITION BY day
