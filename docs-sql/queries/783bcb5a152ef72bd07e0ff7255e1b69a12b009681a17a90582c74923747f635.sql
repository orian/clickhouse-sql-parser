CREATE TABLE prod
(
  uid Int16,
  name String,
  age Int16
)
ENGINE=MergeTree
ORDER BY ();

CREATE TABLE staging
(
  uid Int16,
  name String,
  age Int16
)
ENGINE=MergeTree
ORDER BY ();

-- Datos iniciales
INSERT INTO prod VALUES (123, 'John', 33);
INSERT INTO prod VALUES (456, 'Ksenia', 48);
-- Cargar datos
INSERT INTO staging VALUES (8811, 'Alice', 50);
INSERT INTO staging VALUES (8812, 'Bob', 23);

-- Validar importación
SELECT 'Staging count:', COUNT() FROM staging;
-- Mover partición
ALTER TABLE staging MOVE PARTITION tuple() TO TABLE prod; -- operación atómica

-- Verificar datos
SELECT 'Prod count:', COUNT() FROM prod;
SELECT * FROM prod;
