CREATE FUNCTION str_reverse LANGUAGE WASM ABI BUFFERED_V1
FROM 'str_reverse' :: 'str_reverse'
ARGUMENTS (s String) RETURNS String
SETTINGS serialization_format = 'RowBinary';

SELECT str_reverse(toString(number + 100)) FROM numbers(3);
