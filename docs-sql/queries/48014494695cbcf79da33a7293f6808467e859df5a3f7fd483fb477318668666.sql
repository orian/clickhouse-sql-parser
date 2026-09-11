-- Crie uma tabela e insira os dados
CREATE TABLE test_lazy (json JSON) ENGINE = MergeTree ORDER BY tuple();
INSERT INTO test_lazy VALUES ('{"user_id": "123", "score": "95.5"}');

-- Habilite a configuração experimental
SET allow_experimental_json_lazy_type_hints = 1;

-- Adicione type hints - isso é concluído instantaneamente, sem mutação
ALTER TABLE test_lazy MODIFY COLUMN json JSON(user_id UInt64, score Float64);

-- Consulte os dados - os type hints são aplicados no momento da leitura
SELECT json.user_id, toTypeName(json.user_id), json.score, toTypeName(json.score) FROM test_lazy;
