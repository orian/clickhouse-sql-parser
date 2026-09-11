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
    SETTINGS input_format_null_as_default = 1, -- Veiller à ce que les colonnes soient insérées avec leur valeur par défaut si les valeurs sont NULL
    input_format_parquet_case_insensitive_column_matching = 1 -- La correspondance des colonnes entre les données sources et la table cible doit être insensible à la casse
