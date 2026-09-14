    CREATE FILE FORMAT my_parquet_format TYPE = parquet;

    -- 创建外部 stage，并指定要复制到的 S3 bucket
    CREATE OR REPLACE STAGE external_stage
    URL='s3://mybucket/mydataset'
    CREDENTIALS=(AWS_KEY_ID='<key>' AWS_SECRET_KEY='<secret>')
    FILE_FORMAT = my_parquet_format;

    -- 为所有文件添加 "mydataset" 前缀，并指定最大文件大小为 150mb
    -- `header=true` 参数是获取列名所必需的
    COPY INTO @external_stage/mydataset from mydataset max_file_size=157286400 header=true;
