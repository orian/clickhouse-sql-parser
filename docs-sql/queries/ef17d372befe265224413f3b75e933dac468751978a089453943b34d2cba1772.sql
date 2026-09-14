    CREATE FILE FORMAT my_parquet_format TYPE = parquet;

    -- 복사 대상인 S3 버킷을 지정하는 외부 스테이지를 생성합니다
    CREATE OR REPLACE STAGE external_stage
    URL='s3://mybucket/mydataset'
    CREDENTIALS=(AWS_KEY_ID='<key>' AWS_SECRET_KEY='<secret>')
    FILE_FORMAT = my_parquet_format;

    -- 모든 파일에 "mydataset" 접두사를 적용하고 최대 파일 크기를 150MB로 지정합니다
    -- 컬럼 이름을 가져오려면 `header=true` 매개변수가 필요합니다
    COPY INTO @external_stage/mydataset from mydataset max_file_size=157286400 header=true;
