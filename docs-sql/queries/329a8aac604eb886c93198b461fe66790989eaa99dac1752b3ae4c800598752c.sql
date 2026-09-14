  SELECT
      Origin,
      toStringCutToZero(Dest) AS Destination,
      groupArray(DepDelayMinutes) AS delays,
      round(arrayReduce('avg', groupArray(DepDelayMinutes)), 2) AS avg_delay,
      round(arrayReduce('max', groupArray(DepDelayMinutes)), 2) AS worst_delay
  FROM ontime.ontime
  WHERE Origin = 'DEN'
      AND FlightDate = '2024-01-01'
      AND DepDelayMinutes IS NOT NULL
  GROUP BY Origin, Destination
  ORDER BY avg_delay DESC
