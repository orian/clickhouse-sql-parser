 INDEX idx_text Body TYPE text(tokenizer = splitByNonAlpha,
                               preprocessor = JSONExtract(Body, 'msg', 'String'))
