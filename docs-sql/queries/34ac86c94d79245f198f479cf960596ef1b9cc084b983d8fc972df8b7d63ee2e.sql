    CREATE FILE FORMAT my_parquet_format TYPE = parquet;

    -- Создаем внешний stage, указывающий S3 бакет, в который нужно копировать данные
    CREATE OR REPLACE STAGE external_stage
    URL='s3://mybucket/mydataset'
    CREDENTIALS=(AWS_KEY_ID='<key>' AWS_SECRET_KEY='<secret>')
    FILE_FORMAT = my_parquet_format;

    -- Добавляем префикс "mydataset" ко всем файлам и задаем максимальный размер файла 150 МБ
    -- Параметр `header=true` обязателен, чтобы получить имена столбцов
    COPY INTO @external_stage/mydataset from mydataset max_file_size=157286400 header=true;
