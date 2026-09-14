GRANT SELECT ON db.* TO john -- correto
GRANT SELECT ON db*.* TO john -- correto

GRANT SELECT ON *.my_table TO john -- incorreto
GRANT SELECT ON foo*bar TO john -- incorreto
GRANT SELECT ON *suffix TO john -- incorreto
GRANT SELECT(foo) ON db.table* TO john -- incorreto
