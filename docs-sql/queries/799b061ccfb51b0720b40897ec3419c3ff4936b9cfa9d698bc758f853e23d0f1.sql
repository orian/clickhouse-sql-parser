SELECT
    number % 60 AS t,
    toInt32(intDiv(number, 60) % 64) AS x,
    toInt32((number * 7) % 64) AS y,
    toUInt8(255) AS v
FROM numbers(60 * 64)
INTO OUTFILE 'animation.png'
FORMAT PNG
SETTINGS output_format_image_width = 64, output_format_image_height = 64;
