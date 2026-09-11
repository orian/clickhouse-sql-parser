BEGIN TRANSACTION;
SELECT transactionLatestSnapshot();
ROLLBACK;
