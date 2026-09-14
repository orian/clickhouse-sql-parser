CREATE DATABASE glue
ENGINE = DataLakeCatalog
SETTINGS 
    catalog_type = 'glue', 
    region = 'us-west-2', 
    aws_access_key_id = '<access-key>', 
    aws_secret_access_key = '<secret-key>'
