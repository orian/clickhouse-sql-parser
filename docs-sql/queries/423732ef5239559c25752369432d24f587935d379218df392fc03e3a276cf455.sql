SELECT arrayJoin(['1 B', '1 KiB', '3 MB', '5.314 KiB']) AS readable_sizes, parseReadableSize(readable_sizes) AS sizes;
