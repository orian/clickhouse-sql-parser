  WITH airport_routes AS (
      SELECT 
          Origin,
          arraySort(groupArray(DISTINCT toStringCutToZero(Dest))) AS destinations
      FROM ontime.ontime
      WHERE FlightDate = '2024-01-01'
      GROUP BY Origin
  )
  SELECT 
      a1.Origin AS airport1,
      a2.Origin AS airport2,
      length(arrayIntersect(a1.destinations, a2.destinations)) AS common_destinations
  FROM airport_routes a1
  CROSS JOIN airport_routes a2
  WHERE a1.Origin < a2.Origin
      AND a1.Origin IN ('DEN', 'ATL', 'DFW', 'ORD', 'LAS')
      AND a2.Origin IN ('DEN', 'ATL', 'DFW', 'ORD', 'LAS')
  ORDER BY common_destinations DESC
  LIMIT 10
