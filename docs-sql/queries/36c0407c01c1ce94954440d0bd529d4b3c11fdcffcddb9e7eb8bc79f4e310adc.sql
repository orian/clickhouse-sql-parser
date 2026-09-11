-- delete all data from 2018 with a mutation. Not recommended.
ALTER TABLE posts DELETE WHERE toYear(CreationDate) = 2018
