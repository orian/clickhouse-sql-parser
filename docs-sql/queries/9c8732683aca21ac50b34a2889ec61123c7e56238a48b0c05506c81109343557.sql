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
    SETTINGS input_format_null_as_default = 1, -- 値が null の場合はカラムにデフォルト値を挿入する
    input_format_parquet_case_insensitive_column_matching = 1 -- ソースデータとターゲットテーブルのカラム照合で大文字と小文字を区別しない
