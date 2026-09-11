-- (1) Pausar inserts
-- (2) Criar uma cópia da tabela de destino

CREATE TABLE pypi_v2 AS pypi

SELECT count() FROM pypi_v2
