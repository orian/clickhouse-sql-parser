SELECT regionHierarchy(number::UInt32) AS arr, arrayMap(id -> regionToName(id, 'en'), arr) FROM numbers(5);
