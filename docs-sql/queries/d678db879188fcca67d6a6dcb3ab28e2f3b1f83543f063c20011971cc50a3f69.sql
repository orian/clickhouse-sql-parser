WITH 'hello' || toString(number) AS str
SELECT str,
isConstant(length(str)) AS str_length_is_constant,
isConstant(length(str::FixedString(6))) AS fixed_str_length_is_constant
FROM numbers(3)
