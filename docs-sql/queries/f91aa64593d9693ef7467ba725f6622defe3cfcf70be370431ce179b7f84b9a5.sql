SET dialect = 'kusto';

StormEvents
| extend Bucket = toStartOfHour(StartTime), Fingerprint = cityHash64(EventType)
| summarize Events = count() by Bucket
