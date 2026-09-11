  -- Создать таблицу со всеми необходимыми типами данных
  CREATE TABLE all_types_test (
      `id` UInt32,
      
      -- Тип String
      `String_value` String,
      
      -- Беззнаковые целочисленные типы
      `UInt8_value` UInt8,
      `UInt16_value` UInt16,
      `UInt32_value` UInt32,
      `UInt64_value` UInt64,
      
      -- Знаковые целочисленные типы
      `Int8_value` Int8,
      `Int16_value` Int16,
      `Int32_value` Int32,
      `Int64_value` Int64,
      
      -- Типы с плавающей точкой
      `Float32_value` Float32,
      `Float64_value` Float64,
      
      -- Типы даты и времени
      `Date_value` Date,
      `DateTime_value` DateTime,
      
      -- Сетевые типы
      `IPv4_value` IPv4,
      `IPv6_value` IPv6,
      
      -- Тип UUID
      `UUID_value` UUID
  ) ENGINE = MergeTree() 
  ORDER BY id;
