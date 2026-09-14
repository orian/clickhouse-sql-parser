-- (1) Pausar las inserciones
-- (2) Crear un duplicado de nuestra tabla de destino

CREATE TABLE pypi_v2 AS pypi

SELECT count() FROM pypi_v2
