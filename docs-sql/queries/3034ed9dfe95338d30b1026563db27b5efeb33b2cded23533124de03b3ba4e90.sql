SELECT id, sequenceNextNode('backward', 'last_match')(dt, page, ref = 'ref2', page = 'B') FROM test_flow_basecond GROUP BY id;

                  dt   id   page   ref
 1970-01-01 09:00:01    1   A      ref4
 1970-01-01 09:00:02    1   A      ref3 // El resultado
 1970-01-01 09:00:03    1   B      ref2 // Punto base
 1970-01-01 09:00:04    1   B      ref1 // Esta fila no puede ser punto base porque la columna ref no coincide con 'ref2'.
