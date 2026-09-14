    RESTORE DATABASE test_backups
    FROM AzureBlobStorage(
      '<cadena de conexión del endpoint de AzureBlobStorage>',
      '<contenedor>',
      '<blob>/<uuid>'
    )
