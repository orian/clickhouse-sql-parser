CREATE DATABASE onelake_catalog
ENGINE = DataLakeCatalog('https://onelake.table.fabric.microsoft.com/iceberg')
SETTINGS
catalog_type = 'onelake',
warehouse = 'warehouse_id/data_item_id',
onelake_tenant_id = '<tenant_id>',
oauth_server_uri = 'https://login.microsoftonline.com/<tenant_id>/oauth2/v2.0/token',
auth_scope = 'https://storage.azure.com/.default',
onelake_client_id = '<client_id>',
onelake_client_secret = '<client_secret>'
