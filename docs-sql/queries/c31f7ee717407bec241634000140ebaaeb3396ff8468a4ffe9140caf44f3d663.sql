SELECT
    domainWithoutWWW(URL) AS domain,
    count(),
    any(Title) AS title -- obtendo o primeiro cabeçalho de página encontrado para cada domínio.
FROM hits
GROUP BY domain
