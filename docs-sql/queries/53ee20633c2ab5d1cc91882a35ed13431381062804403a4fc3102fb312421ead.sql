    CREATE FILE FORMAT my_parquet_format TYPE = parquet;

    -- Créer le stage externe qui spécifie le bucket S3 de destination
    CREATE OR REPLACE STAGE external_stage
    URL='s3://mybucket/mydataset'
    CREDENTIALS=(AWS_KEY_ID='<key>' AWS_SECRET_KEY='<secret>')
    FILE_FORMAT = my_parquet_format;

    -- Appliquer le préfixe "mydataset" à tous les fichiers et spécifier une taille de fichier maximale de 150mb
    -- Le paramètre `header=true` est requis pour obtenir les noms de colonnes
    COPY INTO @external_stage/mydataset from mydataset max_file_size=157286400 header=true;
