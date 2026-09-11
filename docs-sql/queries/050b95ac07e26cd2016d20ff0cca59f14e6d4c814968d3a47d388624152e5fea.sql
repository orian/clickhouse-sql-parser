-- Primeiro crie a tabela e insira alguns dados
CREATE TABLE hilbert_numbers(
    n1 UInt32,
    n2 UInt32
)
ENGINE=MergeTree()
ORDER BY n1;
insert into hilbert_numbers (*) values(1, 2);

-- Use nomes de colunas em vez de constantes como argumentos da função
SELECT hilbertEncode(n1, n2) FROM hilbert_numbers;
