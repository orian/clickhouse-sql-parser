SELECT JSONHas('{"a": "hello", "b": [-100, 200.0, 300]}', 'b') = 1;
SELECT JSONHas('{"a": "hello", "b": [-100, 200.0, 300]}', 'b', 4) = 0;
