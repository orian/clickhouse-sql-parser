SELECT
    domainWithoutWWW(URL) AS domain,
    count(),
    any(Title) AS title -- 获取每个域名中第一个出现的页面标题。
FROM hits
GROUP BY domain
