WITH arrayMap(
              d -> if(d >= 30, 'DELAYED', if(d >= 15, 'WARNING', 'ON-TIME')),
              groupArray(DepDelayMinutes)
    ) AS statuses

SELECT
    Origin,
    toStringCutToZero(Dest) AS Destination,
    arrayZip(groupArray(Tail_Number), statuses) as tailNumberStatuses
FROM ontime.ontime
WHERE Origin = 'DEN'
  AND FlightDate = '2024-01-01'
  AND DepTime IS NOT NULL
  AND DepDelayMinutes IS NOT NULL
GROUP BY ALL
