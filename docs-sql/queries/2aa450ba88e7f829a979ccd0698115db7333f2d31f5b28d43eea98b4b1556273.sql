SELECT isValidJSON('{"a": "hello", "b": [-100, 200.0, 300]}') = 1;
SELECT isValidJSON('not JSON') = 0;
