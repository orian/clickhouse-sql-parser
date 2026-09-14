  SELECT
      Origin,
      OriginCityName,
      length(arrayFilter(d -> d >= 30, groupArray(ArrDelayMinutes))) AS num_delays_30_min_or_more
  FROM ontime.ontime
  WHERE Origin IN ('DEN', 'ATL', 'DFW')
      AND FlightDate = '2024-01-01'
  GROUP BY Origin, OriginCityName
  ORDER BY num_delays_30_min_or_more DESC
