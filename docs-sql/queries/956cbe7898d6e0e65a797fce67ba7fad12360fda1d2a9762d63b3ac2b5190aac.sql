GRANT SELECT ON db.* TO john -- correcto
GRANT SELECT ON db*.* TO john -- correcto

GRANT SELECT ON *.my_table TO john -- incorrecto
GRANT SELECT ON foo*bar TO john -- incorrecto
GRANT SELECT ON *suffix TO john -- incorrecto
GRANT SELECT(foo) ON db.table* TO john -- incorrecto
