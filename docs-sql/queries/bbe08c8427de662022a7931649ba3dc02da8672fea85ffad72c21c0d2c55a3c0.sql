CREATE DATABASE glue_test2
ENGINE = DataLakeCatalog
SETTINGS 
    catalog_type = 'glue', 
    region = {region}, 
    aws_role_arn = 'arn:aws:iam::{account-id}:role/iceberg-role-for-clickhouse',
    aws_role_session_name = {session-name},
    SETTINGS
    allow_database_glue_catalog = 1;
