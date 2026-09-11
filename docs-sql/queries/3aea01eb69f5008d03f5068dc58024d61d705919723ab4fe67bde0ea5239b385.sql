CREATE MATERIALIZED VIEW nginxdb.access_logs_view
(
  RemoteAddr String,
  Client String,
  RemoteUser String,
  TimeLocal DateTime,
  RequestMethod String,
  Request String,
  HttpVersion String,
  Status Int32,
  BytesSent Int64,
  UserAgent String
)
ENGINE = MergeTree()
ORDER BY RemoteAddr
POPULATE AS
WITH
  splitByWhitespace(message) as split,
  splitByRegexp('\S \d+ "([^"]*)"', message) as referer
SELECT
  split[1] AS RemoteAddr,
  split[2] AS Client,
  split[3] AS RemoteUser,
  parseDateTimeBestEffort(replaceOne(trim(LEADING '[' FROM split[4]), ':', ' ')) AS TimeLocal,
  trim(LEADING '"' FROM split[6]) AS RequestMethod,
  split[7] AS Request,
  trim(TRAILING '"' FROM split[8]) AS HttpVersion,
  split[9] AS Status,
  split[10] AS BytesSent,
  trim(BOTH '"' from referer[2]) AS UserAgent
FROM
  (SELECT message FROM nginxdb.access_logs)
