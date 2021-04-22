<?php

use Illuminate\Support\Facades\Schema;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Database\Migrations\Migration;
use Illuminate\Support\Facades\DB;

class InitTables extends Migration
{
    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
    }

    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        // skip the migration if there are another migrations
        // It means this migration was already applied
        $migrations = DB::select('SELECT * FROM migrations LIMIT 1');
        if (!empty($migrations)) {
            return;
        }
        $oldMigrationTable = DB::select("SHOW TABLES LIKE 'schema_migrations'");
        if (!empty($oldMigrationTable)) {
            return;
        }

        DB::beginTransaction();

        try {
            app("db")->getPdo()->exec($this->getSql());
        } catch (\Throwable $e) {
            DB::rollBack();
            throw $e;
        }

        DB::commit();
    }

    private function getSql()
    {
        return <<<SQL
            CREATE TABLE `configs` (
              `id` int(11) NOT NULL COMMENT 'The config id',
              `path` varchar(255) NOT NULL COMMENT 'The path is constructed as: section/group/field.',
              `value` text NOT NULL COMMENT 'The value.',
              `scope` enum('private','public') NOT NULL DEFAULT 'private',
              `root_only` tinyint(1) NOT NULL DEFAULT '0'
            ) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='Config Data';

            INSERT INTO `configs` (`id`, `path`, `value`, `scope`, `root_only`) VALUES
            (1, 'regional/general/default_timezone', 'Europe/Madrid', 'public', 0),
            (2, 'regional/general/default_date_format', 'DD/MM/YYYY', 'public', 0),
            (3, 'regional/general/site_name', 'Velmie Wallet', 'public', 0),
            (4, 'regional/general/user_register', 'enable', 'public', 0),
            (5, 'regional/modules/velmie_wallet_cards', 'enable', 'private', 0),
            (7, 'regional/login/failed_login_username_use', 'yes', 'private', 0),
            (8, 'regional/login/failed_login_username_limit', '5', 'private', 0),
            (9, 'regional/login/failed_login_username_cleanup', '5', 'private', 0),
            (10, 'regional/login/failed_login_user_use', '100', 'private', 0),
            (11, 'regional/login/failed_login_user_window', '5', 'private', 0),
            (13, 'profile/user-options/dormant', '3', 'private', 0),
            (14, 'profile/user-options/field_user_first_name', 'yes', 'private', 0),
            (15, 'profile/user-options/field_user_last_name', 'yes', 'private', 0),
            (16, 'profile/user-options/field_user_company_name', 'yes', 'private', 0),
            (17, 'profile/user-options/field_user_date_of_birth', 'yes', 'private', 0),
            (18, 'profile/user-options/field_user_document_personal_id', 'yes', 'private', 0),
            (19, 'profile/user-options/field_user_country_of_residence_iso2', 'yes', 'private', 0),
            (20, 'profile/user-options/field_user_country_of_citizenship_iso2', 'yes', 'private', 0),
            (21, 'profile/user-options/field_user_office_phone_number', 'yes', 'private', 0),
            (22, 'profile/user-options/field_user_home_phone_number', 'yes', 'private', 0),
            (23, 'profile/user-options/field_user_phone_number', 'yes', 'private', 0),
            (24, 'profile/user-options/field_user_fax', 'yes', 'private', 0),
            (25, 'profile/user-options/field_user_pa_address', 'yes', 'private', 0),
            (26, 'profile/user-options/field_user_pa_address_2nd_line', 'yes', 'private', 0),
            (27, 'profile/user-options/field_user_pa_city', 'yes', 'private', 0),
            (28, 'profile/user-options/field_user_pa_country_iso2', 'yes', 'private', 0),
            (29, 'profile/user-options/field_user_pa_state_prov_region', 'yes', 'private', 0),
            (30, 'profile/user-options/field_user_pa_zip_postal_code', 'yes', 'private', 0),
            (31, 'profile/user-options/field_user_ma_as_physical', 'yes', 'private', 0),
            (32, 'profile/user-options/field_user_ma_name', 'yes', 'private', 0),
            (33, 'profile/user-options/field_user_ma_address', 'yes', 'private', 0),
            (34, 'profile/user-options/field_user_ma_city', 'yes', 'private', 0),
            (35, 'profile/user-options/field_user_ma_address_2nd_line', 'yes', 'private', 0),
            (36, 'profile/user-options/field_user_ma_state_prov_region', 'yes', 'private', 0),
            (37, 'profile/user-options/field_user_ma_zip_postal_code', 'yes', 'private', 0),
            (38, 'profile/user-options/field_user_ma_phone_number', 'yes', 'private', 0),
            (39, 'profile/user-options/field_user_ma_country_iso2', 'yes', 'private', 0),
            (51, 'profile/autologout/status', 'no', 'public', 0),
            (52, 'profile/autologout/timeout', '10', 'public', 0),
            (53, 'profile/autologout/message', 'Your session is about to expire. Do you want to stay logged in?', 'public', 0),
            (54, 'profile/autologout/inactivity_message', 'You have been logged out', 'public', 0),
            (55, 'profile/autologout/padding', '50', 'public', 0),
            (56, 'regional/general/site_url', 'https://example.com', 'public', 0),
            (57, 'regional/modules/velmie_wallet_gdpr', 'enable', 'public', 0),
            (62, 'regional/general/site_incoming_message_path', '/messages/incoming/{id}', 'private', 0),
            (63, 'regional/general/site_my_profile_settings_path', '/my-profile/settings', 'private', 0),
            (64, 'profile/user-options/field_user_email', 'no', 'private', 0),
            (65, 'regional/general/maintenance', 'disable', 'public', 1),
            (66, 'regional/general/maintenance_text', 'This system is temporarily unavailable.', 'public', 1),
            (67, 'profile/user-options/field_user_sms_phone_number', 'yes', 'private', 0),
            (68, 'regional/general/no-replay-email', 'no-replay@example.com', 'public', 0),
            (69, 'regional/general/total_user_files_storage_limit_mb', '20', 'private', 1),
            (70, 'regional/general/user_file_size_limit_mb', '1', 'private', 1),
            (71, 'regional/general/default_time_format', 'hh:mm A', 'public', 0);

            CREATE TABLE `schema_migrations` (
              `version` bigint(20) NOT NULL,
              `dirty` tinyint(1) NOT NULL
            ) ENGINE=InnoDB DEFAULT CHARSET=latin1;

            INSERT INTO `schema_migrations` (`version`, `dirty`) VALUES
            (20190615132546, 0);

            ALTER TABLE `configs`
              ADD PRIMARY KEY (`id`),
              ADD UNIQUE KEY `path_UNIQUE` (`path`);

            ALTER TABLE `schema_migrations`
              ADD PRIMARY KEY (`version`);

            ALTER TABLE `configs`
              MODIFY `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'The config id', AUTO_INCREMENT=73;
SQL;
    }
}
