SELECT arrayJoin(['1 B', '1 KiB', '3 MB', '5.314 KiB', 'invalid']) AS readable_sizes, parseReadableSizeOrZero(readable_sizes) AS sizes;
