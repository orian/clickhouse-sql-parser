CREATE TABLE test.test_orc
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
ENGINE = Hive('thrift://202.168.117.26:9083', 'test', 'test_orc')
PARTITION BY day

