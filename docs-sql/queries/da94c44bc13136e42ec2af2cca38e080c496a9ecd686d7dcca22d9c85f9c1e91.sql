SET allow_experimental_database_iceberg = 1;

CREATE DATABASE lake
ENGINE = DataLakeCatalog('http://seaweedfs:8181/v1', 'tutorialkey', 'tutorialsecret')
SETTINGS catalog_type = 'rest',
    warehouse = 's3://analytics',
    storage_endpoint = 'http://seaweedfs:8333/analytics',
    catalog_credential = 'tutorialkey:tutorialsecret',
    oauth_server_uri = 'http://seaweedfs:8181/v1/oauth/tokens'
