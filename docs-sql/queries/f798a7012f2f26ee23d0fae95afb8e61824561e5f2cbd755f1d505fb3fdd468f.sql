SELECT id, sequenceNextNode('backward', 'tail')(dt, page, page = 'Basket', page = 'Basket', page = 'Gift') FROM test_flow GROUP BY id;

                 dt   id   page
1970-01-01 09:00:01    1   Home
1970-01-01 09:00:02    1   Gift
1970-01-01 09:00:03    1   Exit // Punto base, no coincide con Basket

1970-01-01 09:00:01    2   Home
1970-01-01 09:00:02    2   Home // El resultado
1970-01-01 09:00:03    2   Gift // Coincide con Gift
1970-01-01 09:00:04    2   Basket // Punto base, coincide con Basket

1970-01-01 09:00:01    3   Gift
1970-01-01 09:00:02    3   Home // El resultado
1970-01-01 09:00:03    3   Gift // Punto base, coincide con Gift
1970-01-01 09:00:04    3   Basket // Punto base, coincide con Basket
