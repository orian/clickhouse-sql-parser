CREATE DATABASE glue
ENGINE = DataLakeCatalog
SETTINGS
    catalog_type = 'glue',
    region = 'us-west-2',
    aws_role_arn = 'arn:aws:iam::1111111111:role/glue_role',
    aws_role_session_name = 'clickhouse-glue-session' 
