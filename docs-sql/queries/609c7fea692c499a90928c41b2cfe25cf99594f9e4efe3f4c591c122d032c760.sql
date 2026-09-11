    BACKUP DATABASE test_backups 
    TO AzureBlobStorage(
      '<cadena de conexión del endpoint de AzureBlobStorage>',
      '<contenedor>',
      '<blob>/<>'
    )
