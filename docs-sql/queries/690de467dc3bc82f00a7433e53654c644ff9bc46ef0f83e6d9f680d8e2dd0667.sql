SELECT id, sequenceNextNode('backward', 'last_match')(dt, page, page = 'Gift', page = 'Gift', page = 'Home') FROM test_flow GROUP BY id;

                 dt   id   page
1970-01-01 09:00:01    1   Home // Coincide con Home, el resultado es null
1970-01-01 09:00:02    1   Gift // Punto base
1970-01-01 09:00:03    1   Exit

1970-01-01 09:00:01    2   Home // El resultado
1970-01-01 09:00:02    2   Home // Coincide con Home
1970-01-01 09:00:03    2   Gift // Punto base
1970-01-01 09:00:04    2   Basket

1970-01-01 09:00:01    3   Gift // El resultado
1970-01-01 09:00:02    3   Home // Coincide con Home
1970-01-01 09:00:03    3   Gift // Punto base
1970-01-01 09:00:04    3   Basket
