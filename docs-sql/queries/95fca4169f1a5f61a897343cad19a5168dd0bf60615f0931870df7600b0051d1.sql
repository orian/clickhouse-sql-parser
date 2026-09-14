SELECT id, sequenceNextNode('forward', 'first_match')(dt, page, ref = 'ref3', page = 'A') FROM test_flow_basecond GROUP BY id;

                  dt   id   page   ref
 1970-01-01 09:00:01    1   A      ref4 // Esta linha não pode ser ponto base porque a coluna ref não corresponde a 'ref3'.
 1970-01-01 09:00:02    1   A      ref3 // Ponto base
 1970-01-01 09:00:03    1   B      ref2 // O resultado
 1970-01-01 09:00:04    1   B      ref1
