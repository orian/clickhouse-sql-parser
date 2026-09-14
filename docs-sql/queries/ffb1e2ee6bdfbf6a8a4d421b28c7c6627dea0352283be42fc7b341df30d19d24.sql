-- Exclui todas as linhas da tabela `hits` onde a coluna `Title` contém o texto `hello`
DELETE FROM hits WHERE Title LIKE '%hello%';
