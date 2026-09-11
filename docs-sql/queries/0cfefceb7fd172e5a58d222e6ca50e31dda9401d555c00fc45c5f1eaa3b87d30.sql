-- delete all data from 2018 with a lightweight delete. Not recommended.
DELETE FROM posts WHERE toYear(CreationDate) = 2018
