-- удалить все данные за 2018 год с помощью легковесного удаления. Не рекомендуется.
DELETE FROM posts WHERE toYear(CreationDate) = 2018
