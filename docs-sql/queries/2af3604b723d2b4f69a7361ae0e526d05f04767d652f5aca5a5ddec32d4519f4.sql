SELECT toUInt8(x * 25) AS r, toUInt8(y * 25) AS g, toUInt8((x + y) * 12) AS b
FROM (SELECT number % 10 AS x, intDiv(number, 10) AS y FROM numbers(100))
FORMAT PNG
SETTINGS output_format_image_width = 10, output_format_image_height = 10, output_format_image_terminal_mode = 'auto';
