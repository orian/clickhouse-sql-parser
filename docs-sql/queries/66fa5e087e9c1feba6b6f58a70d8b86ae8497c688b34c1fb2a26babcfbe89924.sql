SET allow_experimental_database_iceberg = 1;

CREATE DATABASE horizon_catalog
ENGINE = DataLakeCatalog('https://<org>-<account>.snowflakecomputing.com/polaris/api/catalog')
SETTINGS
    catalog_type = 'horizon',
    warehouse = 'ICEBERG_TEST_DB',
    catalog_credential = '<PAT>',
    auth_scope = 'session:role:DATA_ENGINEER',
    oauth_server_uri = 'https://<org>-<account>.snowflakecomputing.com/polaris/api/catalog/v1/oauth/tokens',
    vended_credentials = 1;
