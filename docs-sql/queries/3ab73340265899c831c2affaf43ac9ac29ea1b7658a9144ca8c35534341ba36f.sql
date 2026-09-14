GRANT SELECT ON db.* TO john -- 正しい
GRANT SELECT ON db*.* TO john -- 正しい

GRANT SELECT ON *.my_table TO john -- 誤り
GRANT SELECT ON foo*bar TO john -- 誤り
GRANT SELECT ON *suffix TO john -- 誤り
GRANT SELECT(foo) ON db.table* TO john -- 誤り
