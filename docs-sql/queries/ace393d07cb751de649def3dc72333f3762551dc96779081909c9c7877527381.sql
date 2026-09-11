SELECT * FROM azureBlobStorage(
    '<YOUR CONNECTION STRING>',
    'data-container',
    '*.json',
    'JSONEachRow');
