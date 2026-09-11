SELECT JSONHas('{"a": "hello", "b": [-100, 200.0, 300]}', 0);
SELECT JSONHas('{"a": "hello", "b": [-100, 200.0, 300]}', 1);
SELECT JSONHas('{"a": "hello", "b": [-100, 200.0, 300]}', 2);
SELECT JSONHas('{"a": "hello", "b": [-100, 200.0, 300]}', -1);
SELECT JSONHas('{"a": "hello", "b": [-100, 200.0, 300]}', -2);
SELECT JSONHas('{"a": "hello", "b": [-100, 200.0, 300]}', 3);
