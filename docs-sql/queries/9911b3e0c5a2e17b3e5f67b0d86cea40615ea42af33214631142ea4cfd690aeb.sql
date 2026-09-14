-- La expansión de rango puede ser útil cuando se necesita una distribución similar para
-- argumentos con rangos (o cardinalidad) muy diferentes.
-- Por ejemplo: 'IP Address' (0...FFFFFFFF) y 'Country code' (0...FF).
-- Nota: el tamaño del tuple debe ser igual al número de los demás argumentos.
SELECT hilbertEncode((10, 6), 1024, 16)
