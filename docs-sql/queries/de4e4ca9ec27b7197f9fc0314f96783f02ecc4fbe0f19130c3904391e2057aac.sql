server_id = 1            -- o mayor; cualquier valor excepto 0
log_bin = ON
binlog_format = ROW      -- valor predeterminado
binlog_row_image = FULL  -- valor predeterminado
expire_logs_days = 1     -- o mayor; 0 significaría que los logs se conservan indefinidamente
