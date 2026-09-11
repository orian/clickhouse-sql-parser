  SELECT
      URLDomain,
      groupArrayDistinct(OS) AS distinct_os_codes
  FROM metrica.hits_v1
  WHERE URLDomain != '' -- Учитываем только hits с указанным доменом
  GROUP BY URLDomain
  ORDER BY URLDomain ASC
  LIMIT 20;
