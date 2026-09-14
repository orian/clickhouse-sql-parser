  SELECT
      URLDomain,
      groupArrayDistinct(OS) AS distinct_os_codes
  FROM metrica.hits_v1
  WHERE URLDomain != '' -- Ne prendre en compte que les hits avec un domaine enregistré
  GROUP BY URLDomain
  ORDER BY URLDomain ASC
  LIMIT 20;
