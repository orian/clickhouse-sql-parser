BACKUP TABLE mydb.events
TO AzureBlobStorage('DefaultEndpointsProtocol=https;AccountName=myaccount;AccountKey=...', 'my-container', 'snapshots/events/')
SETTINGS experimental_lightweight_snapshot = true, id = 'events_snapshot_1'
