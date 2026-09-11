SELECT
    domainWithoutWWW(URL) AS domain,
    count(),
    any(Title) AS title -- obtiene el primer encabezado de página encontrado para cada dominio.
FROM hits
GROUP BY domain
