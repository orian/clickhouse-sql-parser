CREATE TABLE test (d Dynamic) ENGINE = Memory;
INSERT INTO test VALUES ('world'), (123), (456);

-- Padrão (lançar exceção em caso de incompatibilidade): length() não aceita inteiros, então a consulta lança uma exceção.
SELECT length(d) FROM test;  -- lança ILLEGAL_TYPE_OF_ARGUMENT

-- Com o lançamento de exceção desativado: linhas incompatíveis retornam NULL.
SET dynamic_throw_on_type_mismatch = false;
SELECT d, length(d) FROM test ORDER BY d::String NULLS LAST;
