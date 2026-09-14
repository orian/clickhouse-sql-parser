-- возвращаем вложенные объекты с помощью нотации ^
SELECT json.^company.labels
FROM people
