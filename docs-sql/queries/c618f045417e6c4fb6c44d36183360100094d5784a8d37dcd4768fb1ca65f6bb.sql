SELECT formatQueryFromJSON(parseQueryToJSON('SELECT a FROM t'), 'SELECT /* comment */ a FROM t');
