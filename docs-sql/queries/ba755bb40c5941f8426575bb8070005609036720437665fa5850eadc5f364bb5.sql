  WITH range(0, 100, 10) AS delay, 
       toStringCutToZero(Dest) AS Destination

  SELECT    
      'Up to ' || del || ' minutes' AS delayTime,
      countIf(DepDelayMinutes >= del) flightsDelayed
  FROM ontime.ontime
  ARRAY JOIN delay AS del
  WHERE Origin = 'DEN' AND Destination = 'MIA' AND FlightDate = '2024-01-01'
  GROUP BY ALL
  ORDER BY flightsDelayed DESC
