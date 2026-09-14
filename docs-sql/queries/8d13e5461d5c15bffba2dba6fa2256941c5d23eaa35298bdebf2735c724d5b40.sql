-- Primero, crear la tabla e insertar algunos datos
CREATE TABLE hilbert_numbers(
    n1 UInt32,
    n2 UInt32
)
ENGINE=MergeTree()
ORDER BY n1;
insert into hilbert_numbers (*) values(1, 2);

-- Usar nombres de columna en lugar de constantes como argumentos de la función
SELECT hilbertEncode(n1, n2) FROM hilbert_numbers;
