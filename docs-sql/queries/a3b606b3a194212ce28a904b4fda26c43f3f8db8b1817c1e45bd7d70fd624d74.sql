CREATE SUBSCRIPTION demo_rds_subscription
CONNECTION 'postgresql://<user>:<password>@<host>:<port>/<database>'
PUBLICATION <pub_name_you_entered_above>;
