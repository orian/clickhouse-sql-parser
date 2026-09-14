  SELECT
      [1, 'ClickHouse', ['Another', 'Array']] AS array,
      array.UInt8,
      array.String,
      array.`Array(String)`
  SETTINGS use_variant_as_common_type = 1;
