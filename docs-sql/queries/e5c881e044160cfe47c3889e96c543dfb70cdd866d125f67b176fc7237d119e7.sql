-- A expansão de intervalo pode ser útil quando você precisa de uma distribuição semelhante para
-- argumentos com intervalos (ou cardinalidade) muito diferentes.
-- Por exemplo: 'Endereço IP' (0...FFFFFFFF) e 'Código de país' (0...FF).
-- Nota: o tamanho da tupla deve ser igual ao número dos demais argumentos.
SELECT hilbertEncode((10, 6), 1024, 16)
