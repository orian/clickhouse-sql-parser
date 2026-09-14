server_id = 1            -- ou maior; qualquer valor exceto 0
log_bin = ON
binlog_format = ROW      -- valor padrão
binlog_row_image = FULL  -- valor padrão
expire_logs_days = 1     -- ou maior; 0 significa que os logs serão preservados indefinidamente
