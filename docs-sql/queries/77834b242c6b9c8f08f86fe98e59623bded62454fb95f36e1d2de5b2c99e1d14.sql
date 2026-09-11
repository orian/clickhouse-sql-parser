CREATE SETTINGS PROFILE OR REPLACE p
    SETTINGS max_execution_time = 10, enable_lazy_columns_replication = 1;

ALTER SETTINGS PROFILE p SETTINGS max_memory_usage = 16106127360;

SHOW CREATE SETTINGS PROFILE p;
-- → CREATE SETTINGS PROFILE p SETTINGS max_memory_usage = 16106127360
-- max_execution_time and enable_lazy_columns_replication are gone.
