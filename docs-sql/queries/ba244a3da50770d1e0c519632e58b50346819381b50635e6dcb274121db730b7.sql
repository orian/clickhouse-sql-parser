  SELECT key, last_value(value_int), last_value(value_string), last_value(value_date)  FROM test_table GROUP BY key; -- غير موصى به.
