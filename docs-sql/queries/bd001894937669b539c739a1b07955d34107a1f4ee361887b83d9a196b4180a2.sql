  SELECT
      toHour(EventTime) AS hour_of_day,
      -- Utiliser sumForEach pour compter les visites sur ordinateur de bureau et sur mobile en une seule passe
      sumForEach([
          IsMobile = 0, -- Visites sur ordinateur de bureau (IsMobile = 0)
          IsMobile = 1  -- Visites sur mobile (IsMobile = 1)
      ]) AS device_counts
  FROM metrica.hits
  GROUP BY hour_of_day
  ORDER BY hour_of_day;
