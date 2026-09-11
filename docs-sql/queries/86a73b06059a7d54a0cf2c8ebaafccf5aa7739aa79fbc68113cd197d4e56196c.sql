CREATE TABLE my_table_dead_letter (
    raw_message   String,
    error_message String,
    stack_trace   String,
    failed_at     DateTime
) ENGINE = MergeTree()
ORDER BY failed_at;
