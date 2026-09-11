CREATE DATABASE polaris_catalog
ENGINE = DataLakeCatalog('https://<catalog_uri>/api/catalog/v1')
SETTINGS
    catalog_type = 'rest',
    catalog_credential = '<client-id>:<client-secret>',
    warehouse = 'snowflake',
    auth_scope = 'PRINCIPAL_ROLE:ALL',
    oauth_server_uri = 'https://<catalog_uri>/api/catalog/v1/oauth/tokens',
    storage_endpoint = '<storage_endpoint>'
