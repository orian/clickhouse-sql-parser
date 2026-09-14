SELECT
    CAST('5', 'UInt8') AS original,
    bin(original) AS original_binary,
    bitNot(original) AS result,
    bin(bitNot(original)) AS result_binary;
