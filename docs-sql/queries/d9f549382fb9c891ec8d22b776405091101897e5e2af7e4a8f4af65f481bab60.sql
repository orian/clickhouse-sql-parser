CREATE DATABASE my_lake
ENGINE = DataLakeCatalog
SETTINGS
    catalog_type = 'glue',
    region = 'us-east-1',
    aws_access_key_id = '<key>',
    aws_secret_access_key = '<secret>'

SELECT count() FROM my_lake.`analytics.events`
