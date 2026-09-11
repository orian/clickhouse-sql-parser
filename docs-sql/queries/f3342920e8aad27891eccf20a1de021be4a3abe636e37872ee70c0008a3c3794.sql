-- Prioridad menor: enmascaramiento básico para todos los datos sensibles
CREATE MASKING POLICY basic_masking ON orders
    UPDATE
        name = '****',
        email = '****@****.com'
    TO masked_data_viewer
    PRIORITY 0;

-- Prioridad mayor: enmascaramiento más refinado (sobrescribe basic_masking)
CREATE MASKING POLICY refined_masking ON orders
    UPDATE
        name = replaceRegexpOne(name, '^([A-Za-z]+)\\s+(.*)$', '\\1 ****')
    WHERE total_amount > 100
    TO masked_data_viewer
    PRIORITY 10;
