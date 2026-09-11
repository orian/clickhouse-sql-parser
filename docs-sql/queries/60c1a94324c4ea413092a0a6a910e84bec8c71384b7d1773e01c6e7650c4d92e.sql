SET allow_experimental_nlp_functions = 1;

SELECT
    detectTonality('Шарик - хороший пёс'),
    detectTonality('Шарик - пёс'),
    detectTonality('Шарик - плохой пёс')
