fields user.username
| stats count(*) as count by user.username
