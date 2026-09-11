  SELECT
      [1, 'ClickHouse', ['Another', 'Array']] AS array,
      toTypeName(array)
  SETTINGS use_variant_as_common_type = 1;
