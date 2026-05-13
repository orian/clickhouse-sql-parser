SELECT
  coalesce(
    very_long_function_name_a(a, b, c),
    very_long_function_name_b(d, e, f),
    some_default_value
  )
FROM
  t
