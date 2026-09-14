CREATE TABLE test (d Dynamic) ENGINE = Memory;
INSERT INTO test VALUES ('world'), (123), (456);

-- Por defecto (lanzar excepción en caso de incompatibilidad): length() no acepta enteros, por lo que la consulta lanza una excepción.
SELECT length(d) FROM test;  -- lanza ILLEGAL_TYPE_OF_ARGUMENT

-- Con throw deshabilitado: las filas incompatibles devuelven NULL.
SET dynamic_throw_on_type_mismatch = false;
SELECT d, length(d) FROM test ORDER BY d::String NULLS LAST;
