SET default_table_engine = 'Log';

SELECT name, value, changed FROM system.settings WHERE name = 'default_table_engine';
