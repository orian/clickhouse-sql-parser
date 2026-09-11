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
    SETTINGS input_format_null_as_default = 1, -- 값이 null인 경우 컬럼이 기본값으로 삽입되도록 합니다
    input_format_parquet_case_insensitive_column_matching = 1 -- 소스 데이터와 대상 테이블 간 컬럼 매칭은 대소문자를 구분하지 않아야 합니다
