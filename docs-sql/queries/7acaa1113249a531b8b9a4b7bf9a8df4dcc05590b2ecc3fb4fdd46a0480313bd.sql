SELECT JSONLength('{"a": "hello", "b": [-100, 200.0, 300]}', 'b') = 3;
SELECT JSONLength('{"a": "hello", "b": [-100, 200.0, 300]}') = 2;
