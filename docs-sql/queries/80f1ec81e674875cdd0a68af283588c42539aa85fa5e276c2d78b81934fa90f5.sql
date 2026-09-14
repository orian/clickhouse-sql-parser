       -- 访问目标数据库的权限：CREATE TABLE、DROP TABLE、SELECT、INSERT、TRUNCATE 等。
       GRANT ALL ON <database>.* TO <user>;

       -- 用于元数据发现和分区管理的系统表访问权限。
       -- 上述数据库授权不包含这些权限。
       GRANT SELECT ON system.columns TO <user>;
       GRANT SELECT ON system.parts TO <user>;
       GRANT SELECT ON system.tables TO <user>;
