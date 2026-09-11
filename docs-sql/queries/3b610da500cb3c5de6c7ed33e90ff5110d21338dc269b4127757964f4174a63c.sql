SET default_temporary_table_engine = 'Log';

CREATE TEMPORARY TABLE my_table (
    x UInt32,
    y UInt32
);

SHOW CREATE TEMPORARY TABLE my_table;
