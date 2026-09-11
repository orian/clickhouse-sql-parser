/*
 * Calcula o número de dias entre duas datas.
 * /* Retorna NULL se algum dos argumentos for NULL */
 */
SELECT
    dateDiff('day', toDate('2024-01-01'), toDate('2024-12-31')) AS days_in_year, -- 365
    dateDiff('day', toDate('2020-01-01'), today()) AS days_since  #! desde 2020
    ///////////////////////////////////////////////////////////////////
    # TODO: adicionar variantes de hora/minuto
