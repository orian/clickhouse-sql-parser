SELECT JSONType('{"a": "hello", "b": [-100, 200.0, 300]}') = 'Object';
SELECT JSONType('{"a": "hello", "b": [-100, 200.0, 300]}', 'a') = 'String';
SELECT JSONType('{"a": "hello", "b": [-100, 200.0, 300]}', 'b') = 'Array';
