-- ✗ Nombres de columna en SELECT (usar Identifier con cuidado)
SELECT {column: Identifier} FROM users;  -- Soporte limitado

-- ✗ Fragmentos SQL arbitrarios
SELECT * FROM users {where_clause: String};  -- NO SOPORTADO

-- ✗ Sentencias ALTER TABLE
ALTER TABLE {table: Identifier} ADD COLUMN new_col String;  -- NO SOPORTADO

-- ✗ Múltiples sentencias
{statements: String};  -- NO SOPORTADO
