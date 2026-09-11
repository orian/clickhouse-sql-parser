-- Сначала создайте таблицу и вставьте данные
CREATE TABLE hilbert_numbers(
    n1 UInt32,
    n2 UInt32
)
ENGINE=MergeTree()
ORDER BY n1;
insert into hilbert_numbers (*) values(1, 2);

-- Используйте имена столбцов вместо констант в качестве аргументов функции
SELECT hilbertEncode(n1, n2) FROM hilbert_numbers;
