  SELECT
      FlightDate,
      Origin,
      groupArray(toStringCutToZero(Dest)) AS Destinations
  FROM ontime.ontime
  WHERE Origin IN ('ATL', 'ORD', 'DFW', 'DEN', 'LAX', 'JFK', 'LAS', 'CLT', 'SFO', 'SEA') AND FlightDate='2024-01-01'
  GROUP BY FlightDate, Origin
  ORDER BY length(Destinations)
