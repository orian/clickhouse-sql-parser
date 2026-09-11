INSERT INTO stackoverflow.posts SELECT * FROM postgresql('<host>:<port>', 'postgres', 'posts', '<username>', '<password>')
