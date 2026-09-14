GRANT SELECT ON db.* TO john -- верно
GRANT SELECT ON db*.* TO john -- верно

GRANT SELECT ON *.my_table TO john -- неверно
GRANT SELECT ON foo*bar TO john -- неверно
GRANT SELECT ON *suffix TO john -- неверно
GRANT SELECT(foo) ON db.table* TO john -- неверно
