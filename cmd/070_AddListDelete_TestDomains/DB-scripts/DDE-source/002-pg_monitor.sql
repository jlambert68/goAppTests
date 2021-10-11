create role pg_monitor;

grant pg_read_all_settings to pg_monitor;

grant pg_read_all_stats to pg_monitor;

grant pg_stat_scan_tables to pg_monitor;

