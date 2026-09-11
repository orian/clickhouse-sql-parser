  -- 创建包含所有必需数据类型的表
  CREATE TABLE all_types_test (
      `id` UInt32,
      
      -- String 类型
      `String_value` String,
      
      -- 无符号整数类型
      `UInt8_value` UInt8,
      `UInt16_value` UInt16,
      `UInt32_value` UInt32,
      `UInt64_value` UInt64,
      
      -- 有符号整数类型
      `Int8_value` Int8,
      `Int16_value` Int16,
      `Int32_value` Int32,
      `Int64_value` Int64,
      
      -- 浮点类型
      `Float32_value` Float32,
      `Float64_value` Float64,
      
      -- 日期/时间类型
      `Date_value` Date,
      `DateTime_value` DateTime,
      
      -- 网络类型
      `IPv4_value` IPv4,
      `IPv6_value` IPv6,
      
      -- UUID 类型
      `UUID_value` UUID
  ) ENGINE = MergeTree() 
  ORDER BY id;
