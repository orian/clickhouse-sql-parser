    -- 授予读取 schema 信息的权限
    GRANT SELECT ON information_schema.* TO your_databrain_user;

    -- 授予对您的 database 和表的只读权限
    GRANT SELECT ON your_database.* TO your_databrain_user;
