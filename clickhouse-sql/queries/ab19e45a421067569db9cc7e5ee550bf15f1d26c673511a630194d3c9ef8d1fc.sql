
    SELECT a, b, ROW_NUMBER() OVER (PARTITION BY a) AS rn
    FROM t
    SETTINGS query_plan_reuse_storage_ordering_for_window_functions=1
    FORMAT Null
  