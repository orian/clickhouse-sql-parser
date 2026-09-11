nonNegativeDerivative(metric_column, timestamp_column[, INTERVAL X UNITS])
  OVER ([[PARTITION BY grouping_column] [ORDER BY sorting_column]
        [ROWS, RANGE, or GROUPS expression_to_bound_rows_within_the_group]] | [window_name])
FROM table_name
WINDOW window_name AS ([PARTITION BY grouping_column] [ORDER BY sorting_column] [ROWS, RANGE, or GROUPS expression_to_bound_rows_within_the_group])
