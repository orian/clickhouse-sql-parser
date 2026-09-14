-- Создать роль
CREATE ROLE user_role

-- Предоставить права только на чтение для таблицы events.
GRANT SELECT ON default.events TO user_role
GRANT user_role TO user_1
GRANT user_role TO user_2
