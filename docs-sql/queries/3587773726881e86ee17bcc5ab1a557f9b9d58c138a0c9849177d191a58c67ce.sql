SELECT
    URLDomain,
    groupArrayDistinct(OS) AS distinct_os_codes
FROM metrica.hits_v1
WHERE URLDomain != '' -- Consider only hits with a recorded domain
GROUP BY URLDomain
ORDER BY URLDomain ASC
LIMIT 20;
