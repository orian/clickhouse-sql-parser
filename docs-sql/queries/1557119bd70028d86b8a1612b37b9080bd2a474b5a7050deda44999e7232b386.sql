CREATE DATABASE horizon_catalog
ENGINE = DataLakeCatalog('https://<org>-<account>.snowflakecomputing.com/polaris/api/catalog')
SETTINGS
    catalog_type = 'horizon',
    warehouse = 'ICEBERG_TEST_DB',
    auth_header = 'Authorization: Bearer <ACCESS_TOKEN>',
    vended_credentials = 1;
