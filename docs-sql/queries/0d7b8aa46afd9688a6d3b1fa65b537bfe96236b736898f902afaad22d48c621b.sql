SELECT
    domainWithoutWWW(URL) AS domain,
    count(),
    any(Title) AS title -- получение первого встреченного заголовка страницы для каждого домена.
FROM hits
GROUP BY domain
