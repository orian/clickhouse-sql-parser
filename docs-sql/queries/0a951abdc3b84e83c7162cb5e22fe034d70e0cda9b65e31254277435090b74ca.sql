CREATE TABLE users () WITH (
    copy_from     = 's3://my-bucket/users.csv',
    access_key    = 'AKIAIOSFODNN7EXAMPLE',
    access_secret = 'wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY',
    format        = 'CSVWithNames',
    fillfactor    = 90
);
