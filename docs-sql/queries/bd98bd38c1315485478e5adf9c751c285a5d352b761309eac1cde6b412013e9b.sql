  SELECT
      toHour(EventTime) AS hour_of_day,
      -- استخدم sumForEach لعدّ زيارات أجهزة سطح المكتب والأجهزة المحمولة في خطوة واحدة
      sumForEach([
          IsMobile = 0, -- زيارات أجهزة سطح المكتب (IsMobile = 0)
          IsMobile = 1  -- زيارات الأجهزة المحمولة (IsMobile = 1)
      ]) AS device_counts
  FROM metrica.hits
  GROUP BY hour_of_day
  ORDER BY hour_of_day;
