SELECT
    hmac('SHA256', 'message', 'key') = HMAC('sha256', 'message', 'key') AS same_result,
    HMAC('SHA256', 'message', 'key') = Hmac('Sha256', 'message', 'key') AS also_same;
