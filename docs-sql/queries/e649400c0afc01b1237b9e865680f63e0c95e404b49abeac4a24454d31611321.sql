CREATE USER mysql4foobar_readonly
IDENTIFIED WITH double_sha1_password BY 'YourPassword42$'
SETTINGS PROFILE 'readonly_profile';
