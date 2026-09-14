GRANT SELECT ON db.* TO john -- 正确
GRANT SELECT ON db*.* TO john -- 正确

GRANT SELECT ON *.my_table TO john -- 错误
GRANT SELECT ON foo*bar TO john -- 错误
GRANT SELECT ON *suffix TO john -- 错误
GRANT SELECT(foo) ON db.table* TO john -- 错误
