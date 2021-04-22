module github.com/Confialink/wallet-settings

go 1.13

replace github.com/Confialink/wallet-settings/rpc/proto/settings => ./rpc/proto/settings

require (
	github.com/Confialink/wallet-permissions/rpc/permissions v0.0.0-20210218064621-7b7ddad868c8
	github.com/Confialink/wallet-pkg-discovery/v2 v2.0.0-20210217105157-30e31661c1d1
	github.com/Confialink/wallet-pkg-env_config v0.0.0-20210217112253-9483d21626ce
	github.com/Confialink/wallet-pkg-env_mods v0.0.0-20210217112432-4bda6de1ee2c
	github.com/Confialink/wallet-pkg-errors v1.0.2
	github.com/Confialink/wallet-pkg-service_names v0.0.0-20210217112604-179d69540dea
	github.com/Confialink/wallet-settings/rpc/proto/settings v0.0.0-00010101000000-000000000000
	github.com/Confialink/wallet-users/rpc/proto/users v0.0.0-20210218063041-d8de466c67cb
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.6.3
	github.com/go-playground/universal-translator v0.17.0
	github.com/go-playground/validator/v10 v10.2.0
	github.com/inconshreveable/log15 v0.0.0-20200109203555-b30bc20e4fd1
	github.com/jinzhu/gorm v1.9.15
	github.com/onsi/ginkgo v1.14.0
	github.com/onsi/gomega v1.10.1
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.6.1
	github.com/twitchtv/twirp v5.12.0+incompatible
)
