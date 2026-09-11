let MultiplyByN = (val: long, n: long = 2) { val * n };
let RecentErrors = (since: timespan) { Logs | where Level == 'Error' and Timestamp > ago(since) };

RecentErrors(1h) | summarize Count = count() by Component
