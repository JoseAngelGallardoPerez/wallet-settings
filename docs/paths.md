# Configuration paths
This topic lists all configuration paths

### General
These configuration values are available in the Settings -> System Config -> General

| Name | Path | Default |
| ------ | ------ | ------ | 
| Default time zone | regional/general/default_timezone | Europe/Madrid |
| Date format | regional/general/default_date_format | DD/MM/YYYY |
| Site name | regional/general/site_name | Velmie Wallet |
| Online sign-up form | regional/general/user_register | enable |

### Modules
These configuration values are available in the Settings -> System Config -> Modules

| Name | Path | Default |
| ------ | ------ | ------ | 
| Card management module | regional/modules/velmie_wallet_cards | enable |
| GDPR Module | regional/modules/velmie_wallet_gdpr | enable |

### Login Security
These configuration values are available in the Settings -> System Config -> Login Security

| Name | Path | Default |
| ------ | ------ | ------ | 
| Active | regional/login/failed_login_username_use | yes |
| Number of failed attempts | regional/login/failed_login_username_limit | 5 |
| Attempts reset | regional/login/failed_login_username_cleanup | 10 |
| Number of failed attempts | regional/login/failed_login_user_use | 6 |
| Blocking duraction | regional/login/failed_login_user_window | 1 |


### Autologout
These configuration values are available in the Settings -> Profile -> Auto logout

| Name | Path | Default |
| ------ | ------ | ------ | 
| Autologout for regular user | profile/autologout/status | enable |
| Autologout timeout | profile/autologout/timeout | 30 |
| Autologout message | profile/autologout/message | |
| Message after autologout | profile/autologout/inactivity_message | You have been auto logout |
| Timeout padding | profile/autologout/padding | 13 |


### User Options
These configuration values are available in the Settings -> Profile -> User Options

| Name | Path | Default |
| ------ | ------ | ------ | 
| Beneficial Owner Fields | profile/user-options/field_user_beneficial_owner | enable |
| Dormant Status After ( months ) | profile/user-options/dormant | 1 |

Fields that should be available for editing:

| Name | Path | Default |
| ------ | ------ | ------ | 
| First Name | profile/user-options/field_user_first_name | yes |
| Last Name | profile/user-options/field_user_last_name | yes |
| Date of Birth | profile/user-options/field_user_date_of_birth | yes |
| Company Name | profile/user-options/field_user_company_name | yes |
| | profile/user-options/field_user_document_personal_id | yes |
| | profile/user-options/field_user_country_of_residence_iso2 | yes |
| | profile/user-options/field_user_country_of_citizenship_iso2 | yes |
| | profile/user-options/field_user_office_phone_number | yes |
| | profile/user-options/field_user_home_phone_number | yes |
| | profile/user-options/field_user_phone_number | yes |
| | profile/user-options/field_user_fax | yes |
| | profile/user-options/field_user_pa_address | yes |
| | profile/user-options/field_user_pa_address_2nd_line | yes |
| | profile/user-options/field_user_pa_city | yes |
| | profile/user-options/field_user_pa_country_iso2 | yes |
| | profile/user-options/field_user_pa_state_prov_region | yes |
| | profile/user-options/field_user_pa_zip_postal_code | yes |
| | profile/user-options/field_user_ma_as_physical | yes |
| | profile/user-options/field_user_ma_name | yes |
| | profile/user-options/field_user_ma_address | yes |
| | profile/user-options/field_user_ma_city | yes |
| | profile/user-options/field_user_ma_address_2nd_line | yes |
| | profile/user-options/field_user_ma_state_prov_region | yes |
| | profile/user-options/field_user_ma_zip_postal_code | yes |
| | profile/user-options/field_user_ma_phone_number | yes |
| | profile/user-options/field_user_ma_country_iso2 | yes |
| | profile/user-options/field_user_bo_full_name | yes |
| | profile/user-options/field_user_bo_date_of_birth | yes |
| | profile/user-options/field_user_bo_document_type | yes |
| | profile/user-options/field_user_bo_document_personal_id | yes |
| | profile/user-options/field_user_bo_relationship | yes |
| | profile/user-options/field_user_bo_address | yes |
| | profile/user-options/field_user_bo_phone_number | yes |
| | profile/user-options/field_user_email | yes |