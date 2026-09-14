SELECT
    regexpPosition('hello world', 'world'),
    regexpPosition('aXbXcXd', 'X', 1, 2),
    regexpPosition('aXbXcXd', 'X', 1, 2, 1),
    regexpPosition('Hello WORLD', 'world', 1, 1, 0, 'i'),
    regexpPosition('foo123bar456', '([a-z]+)([0-9]+)', 1, 2, 0, '', 2);
