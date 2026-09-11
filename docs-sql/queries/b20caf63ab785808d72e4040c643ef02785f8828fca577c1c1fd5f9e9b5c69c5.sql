    CREATE FILE FORMAT my_parquet_format TYPE = parquet;

    -- コピー先の S3 bucket を指定する外部ステージを作成する
    CREATE OR REPLACE STAGE external_stage
    URL='s3://mybucket/mydataset'
    CREDENTIALS=(AWS_KEY_ID='<key>' AWS_SECRET_KEY='<secret>')
    FILE_FORMAT = my_parquet_format;

    -- すべてのファイルに "mydataset" プレフィックスを適用し、最大ファイルサイズを 150MB に指定する
    -- カラム名を取得するには `header=true` パラメータが必要
    COPY INTO @external_stage/mydataset from mydataset max_file_size=157286400 header=true;
