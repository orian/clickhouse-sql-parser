SYSTEM START MERGES test;
OPTIMIZE TABLE test FINAL;
SELECT count(), dynamicType(d), isDynamicElementInSharedData(d), _part FROM test GROUP BY _part, dynamicType(d), isDynamicElementInSharedData(d) ORDER BY _part, count() desc;
