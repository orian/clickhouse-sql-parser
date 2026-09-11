-- Create table with all the required data types
CREATE TABLE all_types_test (
    `id` UInt32,
    
    -- String type
    `String_value` String,
    
    -- Unsigned integer types
    `UInt8_value` UInt8,
    `UInt16_value` UInt16,
    `UInt32_value` UInt32,
    `UInt64_value` UInt64,
    
    -- Signed integer types
    `Int8_value` Int8,
    `Int16_value` Int16,
    `Int32_value` Int32,
    `Int64_value` Int64,
    
    -- Floating point types
    `Float32_value` Float32,
    `Float64_value` Float64,
    
    -- Date/time types
    `Date_value` Date,
    `DateTime_value` DateTime,
    
    -- Network types
    `IPv4_value` IPv4,
    `IPv6_value` IPv6,
    
    -- UUID type
    `UUID_value` UUID
) ENGINE = MergeTree() 
ORDER BY id;
