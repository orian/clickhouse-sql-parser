SELECT
    toInt32(x) AS x,
    toInt32(y) AS y,
    toUInt8(intensity) AS v
FROM points
INTO OUTFILE 'points.png'
FORMAT PNG
SETTINGS output_format_image_width = 512, output_format_image_height = 512;
