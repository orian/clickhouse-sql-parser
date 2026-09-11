CREATE NAMED COLLECTION mongo_creds AS
       uri='mongodb://test_user:password@127.0.0.1:27017/test?connectionTimeoutMS=10000',
       collection='default_collection';

SELECT * FROM mongodb(
        mongo_creds,
        collection = 'my_collection',
        structure = 'log_type String, host String, command String'
)
