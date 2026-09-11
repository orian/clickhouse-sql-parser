SELECT 8374283493092 AS id,
       timeSeriesStoreTags(id, [('region', 'eu'), ('env', 'dev')], '__name__', 'http_requests_count') AS same_id,
       throwIf(same_id != id),
       timeSeriesIdToTags(same_id),
       timeSeriesGroupToTags(timeSeriesIdToGroup(same_id))
