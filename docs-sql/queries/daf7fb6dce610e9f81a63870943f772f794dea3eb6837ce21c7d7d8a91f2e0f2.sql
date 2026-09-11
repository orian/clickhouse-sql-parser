BEGIN TRANSACTION;
SELECT transactionOldestSnapshot();
ROLLBACK;
