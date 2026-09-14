  SELECT
      URLDomain,
      groupArrayDistinct(OS) AS distinct_os_codes
  FROM metrica.hits_v1
  WHERE URLDomain != '' -- Considere apenas hits com um domínio registrado
  GROUP BY URLDomain
  ORDER BY URLDomain ASC
  LIMIT 20;
