SELECT
    domainWithoutWWW(URL) AS domain,
    count(),
    any(Title) AS title -- 各ドメインで最初に出現したページヘッダーを取得する。
FROM hits
GROUP BY domain
