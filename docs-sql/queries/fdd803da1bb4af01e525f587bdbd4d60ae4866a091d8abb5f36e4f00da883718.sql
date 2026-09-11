-- exclui todos os dados de 2018 com uma exclusão leve. Não recomendado.
DELETE FROM posts WHERE toYear(CreationDate) = 2018
