SELECT count()
FROM bigquery('my-project', 'my_dataset', 'my_table',
              service_account_key = '{"type": "service_account", "private_key": "...", "client_email": "...", ...}');
