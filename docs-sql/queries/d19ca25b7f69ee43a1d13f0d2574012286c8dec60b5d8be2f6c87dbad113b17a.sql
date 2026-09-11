ALTER TABLE hackernews
    ADD INDEX comment_idx lower(comment)
    TYPE text(tokenizer = splitByNonAlpha);

ALTER TABLE hackernews
    MATERIALIZE INDEX comment_idx
    SETTINGS mutations_sync = 2;
