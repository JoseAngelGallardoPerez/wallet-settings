<?php

use Illuminate\Support\Facades\Schema;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Database\Migrations\Migration;

class AddDefaultPermissions extends Migration
{
    const TABLE_CONFIGS = "configs";

    const DEFAULT_PERMISSIONS_PATHS = [
        "profile/default-user-classes/client" => "2", # default ID for role "Client"
    ];

    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        foreach (self::DEFAULT_PERMISSIONS_PATHS as $path => $value) {
            DB::table(self::TABLE_CONFIGS)->insert([
                'path' => $path,
                'value' => $value,
                'scope' => 'private',
                'root_only' => 0
            ]);
        }
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        foreach (self::DEFAULT_PERMISSIONS_PATHS as $path => $value) {
            DB::delete("DELETE FROM configs WHERE `path` = ?", [$path]);
        }
    }
}
