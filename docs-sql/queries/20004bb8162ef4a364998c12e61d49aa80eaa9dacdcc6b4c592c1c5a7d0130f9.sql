server_id = 1               -- o mayor; cualquier valor excepto 0
log_bin = ON
binlog_format = ROW
binlog_row_image = FULL
binlog_row_metadata = FULL  -- introducido en 10.5.0
expire_logs_days = 1        -- o mayor; 0 significaría que los logs se conservan indefinidamente
