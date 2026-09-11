SELECT id, sequenceNextNode('backward', 'tail')(dt, page, page = 'Basket', page = 'Basket', page = 'Gift') FROM test_flow GROUP BY id;

                 dt   id   page
1970-01-01 09:00:01    1   Home
1970-01-01 09:00:02    1   Gift
1970-01-01 09:00:03    1   Exit // Базовая точка, не совпадает с Basket

1970-01-01 09:00:01    2   Home
1970-01-01 09:00:02    2   Home // Результат
1970-01-01 09:00:03    2   Gift // Совпадает с Gift
1970-01-01 09:00:04    2   Basket // Базовая точка, совпадает с Basket

1970-01-01 09:00:01    3   Gift
1970-01-01 09:00:02    3   Home // Результат
1970-01-01 09:00:03    3   Gift // Базовая точка, совпадает с Gift
1970-01-01 09:00:04    3   Basket // Базовая точка, совпадает с Basket
