    CREATE FILE FORMAT my_parquet_format TYPE = parquet;

    -- Crie o stage externo que especifica o bucket do S3 para o qual copiar
    CREATE OR REPLACE STAGE external_stage
    URL='s3://mybucket/mydataset'
    CREDENTIALS=(AWS_KEY_ID='<key>' AWS_SECRET_KEY='<secret>')
    FILE_FORMAT = my_parquet_format;

    -- Aplique o prefixo "mydataset" a todos os arquivos e especifique um tamanho máximo de arquivo de 150mb
    -- O parâmetro `header=true` é necessário para obter os nomes das colunas
    COPY INTO @external_stage/mydataset from mydataset max_file_size=157286400 header=true;
