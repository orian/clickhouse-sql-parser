  -- 필요한 모든 데이터 타입이 포함된 테이블 생성
  CREATE TABLE all_types_test (
      `id` UInt32,
      
      -- String 타입
      `String_value` String,
      
      -- 부호 없는 정수 타입
      `UInt8_value` UInt8,
      `UInt16_value` UInt16,
      `UInt32_value` UInt32,
      `UInt64_value` UInt64,
      
      -- 부호 있는 정수 타입
      `Int8_value` Int8,
      `Int16_value` Int16,
      `Int32_value` Int32,
      `Int64_value` Int64,
      
      -- 부동 소수점 타입
      `Float32_value` Float32,
      `Float64_value` Float64,
      
      -- 날짜/시간 타입
      `Date_value` Date,
      `DateTime_value` DateTime,
      
      -- 네트워크 타입
      `IPv4_value` IPv4,
      `IPv6_value` IPv6,
      
      -- UUID 타입
      `UUID_value` UUID
  ) ENGINE = MergeTree() 
  ORDER BY id;
