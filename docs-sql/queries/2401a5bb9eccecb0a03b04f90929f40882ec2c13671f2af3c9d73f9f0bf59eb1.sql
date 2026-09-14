  SELECT
      toHour(EventTime) AS hour_of_day,
      -- Use sumForEach para contar visitas em desktop e em dispositivos móveis em uma única passagem
      sumForEach([
          IsMobile = 0, -- Visitas em desktop (IsMobile = 0)
          IsMobile = 1  -- Visitas em dispositivos móveis (IsMobile = 1)
      ]) AS device_counts
  FROM metrica.hits
  GROUP BY hour_of_day
  ORDER BY hour_of_day;
