SELECT
    hex(HMAC('md5', 'message', 'key')) AS hmac_md5,
    hex(HMAC('sha1', 'message', 'key')) AS hmac_sha1,
    hex(HMAC('sha256', 'message', 'key')) AS hmac_sha256;
