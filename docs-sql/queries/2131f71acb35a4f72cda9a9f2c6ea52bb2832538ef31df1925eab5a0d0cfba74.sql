-- Pero encrypt falla cuando key o iv son más largos de lo esperado:
SELECT encrypt('aes-256-ofb', 'Secret', '123456789101213141516171819202122', 'iviviviviviviviv123');
