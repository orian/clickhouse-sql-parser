-- Si se proporciona un único argumento con una tupla que especifica los desplazamientos de bits, la función
-- desplaza el argumento hacia la izquierda el número de bits especificado.
SELECT hilbertEncode(tuple(2), 128)
