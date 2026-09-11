    CREATE DATABASE my_lake
    ENGINE = DataLakeCatalog
    SETTINGS
        catalog_type = 'glue',
        region = '<your-region>',
        aws_access_key_id = '<your-access-key>',
        aws_secret_access_key = '<your-secret-key>'
