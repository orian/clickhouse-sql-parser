CREATE TABLE words (class_id UInt32, text String) ENGINE = MergeTree ORDER BY tuple();
INSERT INTO words VALUES
    (0,'dog'),(0,'cat'),(0,'fish'),(0,'bird'),(0,'book'),(0,'hand'),(0,'tree'),(0,'milk'),(0,'duck'),(0,'frog'),(0,'lamp'),(0,'desk'),
    (1,'gato'),(1,'casa'),(1,'perro'),(1,'libro'),(1,'mano'),(1,'leche'),(1,'arbol'),(1,'agua'),(1,'queso'),(1,'fuego'),(1,'mesa'),(1,'silla');

CREATE TABLE word_ngrams (ngram String, class_id UInt32, count UInt64) ENGINE = MergeTree ORDER BY (class_id, ngram);
INSERT INTO word_ngrams
SELECT ngram, class_id, count()
FROM words
ARRAY JOIN naiveBayesNgrams(text, 2, 'codepoint', '0x10FFFE', '0x10FFFF') AS ngram
GROUP BY ngram, class_id;

CREATE DICTIONARY lang (ngram String, class_id UInt32, count UInt64)
PRIMARY KEY ngram SOURCE(CLICKHOUSE(TABLE 'word_ngrams'))
LAYOUT(NAIVE_BAYES(class_attribute 'class_id' n 2 mode 'codepoint' start_token '0x10FFFE' end_token '0x10FFFF')) LIFETIME(0);

SELECT naiveBayesClassifier('lang', 'window') AS window, naiveBayesClassifier('lang', 'fiesta') AS fiesta;
