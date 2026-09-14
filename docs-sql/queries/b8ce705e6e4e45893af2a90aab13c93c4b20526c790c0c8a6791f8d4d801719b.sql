SELECT id, sequenceNextNode('forward', 'first_match')(dt, page, ref = 'ref3', page = 'A') FROM test_flow_basecond GROUP BY id;

                  dt   id   page   ref
 1970-01-01 09:00:01    1   A      ref4 // Эта строка не может быть базовой точкой, так как столбец ref не совпадает с 'ref3'.
 1970-01-01 09:00:02    1   A      ref3 // Базовая точка
 1970-01-01 09:00:03    1   B      ref2 // Результат
 1970-01-01 09:00:04    1   B      ref1
