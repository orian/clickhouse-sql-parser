ALTER TABLE logs ADD INDEX msg_text msg TYPE text(tokenizer = splitByNonAlpha);
ALTER TABLE logs MATERIALIZE INDEX msg_text;

SELECT count() FROM logs WHERE hasAllTokens(msg, 'exception');
