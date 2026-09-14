-- Crear una tabla e insertar datos
CREATE TABLE test_lazy (json JSON) ENGINE = MergeTree ORDER BY tuple();
INSERT INTO test_lazy VALUES ('{"user_id": "123", "score": "95.5"}');

-- Habilitar la configuración experimental
SET allow_experimental_json_lazy_type_hints = 1;

-- Agregar indicaciones de tipo - esto se completa de forma instantánea sin mutación
ALTER TABLE test_lazy MODIFY COLUMN json JSON(user_id UInt64, score Float64);

-- Consultar los datos - las indicaciones de tipo se aplican en el momento de la lectura
SELECT json.user_id, toTypeName(json.user_id), json.score, toTypeName(json.score) FROM test_lazy;
