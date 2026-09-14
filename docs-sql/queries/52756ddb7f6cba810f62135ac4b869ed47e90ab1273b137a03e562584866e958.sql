CREATE TABLE test(tags Array(Tuple(String, String))) engine=Memory;

INSERT INTO test VALUES ([('__name__', 'up')]);

SELECT timeSeriesThrowDuplicateSeriesIf(count() > 1, timeSeriesTagsToGroup(tags))
FROM test
GROUP BY timeSeriesTagsToGroup(tags);
