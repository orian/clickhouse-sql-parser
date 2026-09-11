  -- 必要なすべてのデータ型を持つテーブルを作成
  CREATE TABLE all_types_test (
      `id` UInt32,
      
      -- String 型
      `String_value` String,
      
      -- 符号なし整数型
      `UInt8_value` UInt8,
      `UInt16_value` UInt16,
      `UInt32_value` UInt32,
      `UInt64_value` UInt64,
      
      -- 符号付き整数型
      `Int8_value` Int8,
      `Int16_value` Int16,
      `Int32_value` Int32,
      `Int64_value` Int64,
      
      -- 浮動小数点型
      `Float32_value` Float32,
      `Float64_value` Float64,
      
      -- 日付/時刻型
      `Date_value` Date,
      `DateTime_value` DateTime,
      
      -- ネットワーク型
      `IPv4_value` IPv4,
      `IPv6_value` IPv6,
      
      -- UUID 型
      `UUID_value` UUID
  ) ENGINE = MergeTree() 
  ORDER BY id;
