CREATE NAMED COLLECTION gcs_creds AS
access_key_id = '<access_key>',
secret_access_key = '<secret>';

CREATE TABLE gcpDeltaLake
ENGINE = DeltaLake(gcs_creds, url = 'https://storage.googleapis.com/<bucket>/<path>')
