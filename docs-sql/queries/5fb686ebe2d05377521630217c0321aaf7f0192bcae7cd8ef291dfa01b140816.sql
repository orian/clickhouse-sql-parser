CREATE TABLE hits (SearchEngineID UInt8, Referer String) ENGINE = Memory;

INSERT INTO hits VALUES
    (2, 'http://yandex.ru/search'),
    (2, 'http://yandex.ru/news'),
    (2, 'http://mail.yandex.ru/'),
    (3, 'http://google.ru/search'),
    (4, 'http://duckduckgo.com/'),
    (0, 'http://vkontakte.ru/feed'),
    (0, '');

SELECT
transform(SearchEngineID, [2, 3], ['Yandex', 'Google'], 'Other') AS title,
count() AS c
FROM hits
WHERE SearchEngineID != 0
GROUP BY title
ORDER BY c DESC, title
