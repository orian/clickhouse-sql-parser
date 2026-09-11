  -- Créer une table avec tous les types de données requis
  CREATE TABLE all_types_test (
      `id` UInt32,
      
      -- Type String
      `String_value` String,
      
      -- Types d'entiers non signés
      `UInt8_value` UInt8,
      `UInt16_value` UInt16,
      `UInt32_value` UInt32,
      `UInt64_value` UInt64,
      
      -- Types d'entiers signés
      `Int8_value` Int8,
      `Int16_value` Int16,
      `Int32_value` Int32,
      `Int64_value` Int64,
      
      -- Types à virgule flottante
      `Float32_value` Float32,
      `Float64_value` Float64,
      
      -- Types date/heure
      `Date_value` Date,
      `DateTime_value` DateTime,
      
      -- Types réseau
      `IPv4_value` IPv4,
      `IPv6_value` IPv6,
      
      -- Type UUID
      `UUID_value` UUID
  ) ENGINE = MergeTree() 
  ORDER BY id;
