server_id = 1               -- ou maior; qualquer valor exceto 0
log_bin = ON
binlog_format = ROW
binlog_row_image = FULL
binlog_row_metadata = FULL  -- introduzido na versão 10.5.0
expire_logs_days = 1        -- ou maior; 0 significa que os logs serão preservados indefinidamente
