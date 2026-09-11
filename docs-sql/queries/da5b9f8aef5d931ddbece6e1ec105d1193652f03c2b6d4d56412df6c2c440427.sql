  SELECT
      toHour(EventTime) AS hour_of_day,
      -- Use sumForEach para contar las visitas de escritorio y móviles en una sola pasada
      sumForEach([
          IsMobile = 0, -- Visitas de escritorio (IsMobile = 0)
          IsMobile = 1  -- Visitas móviles (IsMobile = 1)
      ]) AS device_counts
  FROM metrica.hits
  GROUP BY hour_of_day
  ORDER BY hour_of_day;
