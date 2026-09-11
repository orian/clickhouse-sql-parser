-- Without a default, a domain that is not listed is returned unchanged.
SELECT
transform(domain(Referer), ['yandex.ru', 'google.ru', 'vkontakte.ru'], ['www.yandex', 'example.com', 'vk.com']) AS s, count() AS c
FROM hits
GROUP BY domain(Referer)
ORDER BY count() DESC, s
LIMIT 10
