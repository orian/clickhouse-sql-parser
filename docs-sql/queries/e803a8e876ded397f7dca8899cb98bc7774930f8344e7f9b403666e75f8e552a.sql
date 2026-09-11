nth_value (x, offset)
  OVER ([[PARTITION BY grouping_column] [ORDER BY sorting_column] 
        [ROWS, RANGE, or GROUPS expression_to_bound_rows_withing_the_group]] | [window_name])
FROM table_name
WINDOW window_name as ([[PARTITION BY grouping_column] [ORDER BY sorting_column])
