-- La expansión de rango puede ser beneficiosa cuando se necesita una distribución similar para
-- argumentos con rangos (o cardinalidad) muy distintos
-- Por ejemplo: 'IP Address' (0...FFFFFFFF) y 'Country code' (0...FF).
-- Nota: el tamaño del Tuple debe ser igual al número de los demás argumentos.
SELECT mortonEncode((1,2), 1024, 16)
