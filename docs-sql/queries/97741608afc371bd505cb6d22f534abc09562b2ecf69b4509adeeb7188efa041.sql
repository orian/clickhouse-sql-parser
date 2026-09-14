    CREATE FILE FORMAT my_parquet_format TYPE = parquet;

    -- Crear el stage externo que especifica el bucket de S3 al que se copiarán los datos
    CREATE OR REPLACE STAGE external_stage
    URL='s3://mybucket/mydataset'
    CREDENTIALS=(AWS_KEY_ID='<key>' AWS_SECRET_KEY='<secret>')
    FILE_FORMAT = my_parquet_format;

    -- Aplicar el prefijo "mydataset" a todos los archivos y especificar un tamaño máximo de archivo de 150mb
    -- El parámetro `header=true` es obligatorio para obtener los nombres de las columnas
    COPY INTO @external_stage/mydataset from mydataset max_file_size=157286400 header=true;
