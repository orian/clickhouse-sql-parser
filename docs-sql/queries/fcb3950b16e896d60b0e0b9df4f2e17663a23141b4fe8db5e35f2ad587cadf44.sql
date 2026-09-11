SELECT
    domainWithoutWWW('http://user:password@www.example.com:8080/path?query=value#fragment'),
    domainWithoutWWWRFC('http://user:password@www.example.com:8080/path?query=value#fragment');
