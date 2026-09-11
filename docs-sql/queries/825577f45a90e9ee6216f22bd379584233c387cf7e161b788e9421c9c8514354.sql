-- eliminar todos los datos de 2018 con una eliminación ligera. No recomendado.
DELETE FROM posts WHERE toYear(CreationDate) = 2018
