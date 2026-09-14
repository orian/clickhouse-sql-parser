  SELECT
      URLDomain,
      groupArrayDistinct(OS) AS distinct_os_codes
  FROM metrica.hits_v1
  WHERE URLDomain != '' -- Considera solo los hits con un dominio registrado
  GROUP BY URLDomain
  ORDER BY URLDomain ASC
  LIMIT 20;
