SET allow_database_iceberg = 1;

CREATE DATABASE places
ENGINE = DataLakeCatalog('https://catalog.h3-hub.foursquare.com/iceberg')
SETTINGS
    catalog_type = 'rest',
    warehouse = 'places',
    auth_header = 'Authorization: Bearer <YOUR_ACCESS_TOKEN>',
    vended_credentials = 1;
