SELECT name, is_complex, syntax
FROM system.dictionary_layouts
WHERE name IN ('flat', 'hashed', 'complex_key_hashed')
ORDER BY name
