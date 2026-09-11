-- ✗ Nomes de colunas no SELECT (use Identifier com cuidado)
SELECT {column: Identifier} FROM users;  -- Suporte limitado

-- ✗ Fragmentos SQL arbitrários
SELECT * FROM users {where_clause: String};  -- NÃO SUPORTADO

-- ✗ Instruções ALTER TABLE
ALTER TABLE {table: Identifier} ADD COLUMN new_col String;  -- NÃO SUPORTADO

-- ✗ Múltiplas instruções
{statements: String};  -- NÃO SUPORTADO
