SET allow_deprecated_error_prone_window_functions = 1;
SELECT number, neighbor(number, 2) FROM system.numbers LIMIT 10;
