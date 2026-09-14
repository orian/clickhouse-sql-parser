-- 論理削除を使用して2018年のすべてのデータを削除する。非推奨。
DELETE FROM posts WHERE toYear(CreationDate) = 2018
