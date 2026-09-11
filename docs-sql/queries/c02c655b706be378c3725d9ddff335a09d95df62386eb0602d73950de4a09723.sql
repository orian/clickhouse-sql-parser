CREATE TABLE [IF NOT EXISTS] [db.]table_name [ON CLUSTER cluster]
(
    name1 [type1] [DEFAULT|MATERIALIZED|ALIAS expr1],
    name2 [type2] [DEFAULT|MATERIALIZED|ALIAS expr2],
    ...
) ENGINE = FileLog('path_to_logs', 'format_name') SETTINGS
    [poll_timeout_ms = 0,]
    [poll_max_batch_size = 0,]
    [max_block_size = 0,]
    [max_threads = 0,]
    [poll_directory_watch_events_backoff_init = 500,]
    [poll_directory_watch_events_backoff_max = 32000,]
    [poll_directory_watch_events_backoff_factor = 2,]
    [handle_error_mode = 'default']
