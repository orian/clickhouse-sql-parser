    INSERT INTO mydataset
    SELECT
      timestamp,
      some_text,
      JSONExtract(
        ifNull(some_file, '{}'),
        'Tuple(filename String, version String)'
      ) AS some_file,
      JSONExtract(
        ifNull(complex_data, '{}'),
        'Tuple(filename String, description String)'
      ) AS complex_data,
    FROM s3('https://mybucket.s3.amazonaws.com/mydataset/mydataset*.parquet')
    SETTINGS input_format_null_as_default = 1, -- Garantiza que las columnas se inserten con el valor predeterminado si los valores son NULL
    input_format_parquet_case_insensitive_column_matching = 1 -- La coincidencia entre las columnas de los datos de origen y la tabla de destino no debe distinguir entre mayúsculas y minúsculas
