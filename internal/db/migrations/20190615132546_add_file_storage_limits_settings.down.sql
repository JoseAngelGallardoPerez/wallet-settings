DELETE FROM `configs` WHERE `path` IN (
  'regional/general/total_user_files_storage_limit_mb',
  'regional/general/user_file_size_limit_mb'
);