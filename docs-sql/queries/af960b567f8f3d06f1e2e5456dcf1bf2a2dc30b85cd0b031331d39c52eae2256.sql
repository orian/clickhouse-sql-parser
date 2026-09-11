  -- Criar tabela com todos os tipos de dados necessários
  CREATE TABLE all_types_test (
      `id` UInt32,
      
      -- Tipo String
      `String_value` String,
      
      -- Tipos inteiros sem sinal
      `UInt8_value` UInt8,
      `UInt16_value` UInt16,
      `UInt32_value` UInt32,
      `UInt64_value` UInt64,
      
      -- Tipos inteiros com sinal
      `Int8_value` Int8,
      `Int16_value` Int16,
      `Int32_value` Int32,
      `Int64_value` Int64,
      
      -- Tipos de ponto flutuante
      `Float32_value` Float32,
      `Float64_value` Float64,
      
      -- Tipos de data/hora
      `Date_value` Date,
      `DateTime_value` DateTime,
      
      -- Tipos de rede
      `IPv4_value` IPv4,
      `IPv6_value` IPv6,
      
      -- Tipo UUID
      `UUID_value` UUID
  ) ENGINE = MergeTree() 
  ORDER BY id;
