SELECT id, sequenceNextNode('backward', 'tail')(dt, page, page = 'Basket', page = 'Basket', page = 'Gift') FROM test_flow GROUP BY id;

                 dt   id   page
1970-01-01 09:00:01    1   Home
1970-01-01 09:00:02    1   Gift
1970-01-01 09:00:03    1   Exit // Ponto base, não corresponde a Basket

1970-01-01 09:00:01    2   Home
1970-01-01 09:00:02    2   Home // O resultado
1970-01-01 09:00:03    2   Gift // Corresponde a Gift
1970-01-01 09:00:04    2   Basket // Ponto base, corresponde a Basket

1970-01-01 09:00:01    3   Gift
1970-01-01 09:00:02    3   Home // O resultado
1970-01-01 09:00:03    3   Gift // Ponto base, Matched with Gift
1970-01-01 09:00:04    3   Basket // Ponto base, corresponde a Basket
