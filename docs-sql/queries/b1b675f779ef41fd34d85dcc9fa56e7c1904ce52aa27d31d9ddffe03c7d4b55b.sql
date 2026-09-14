SELECT clientip, status, size, `request.method` FROM http WHERE has(request.method, 'GET');
