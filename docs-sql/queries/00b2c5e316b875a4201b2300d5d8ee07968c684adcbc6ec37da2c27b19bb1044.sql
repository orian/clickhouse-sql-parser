        CREATE USER `clickstack-ingest` IDENTIFIED WITH sha256_password BY '<password>';
        GRANT SELECT, INSERT ON logs.nginx_logs TO `clickstack-ingest`;
