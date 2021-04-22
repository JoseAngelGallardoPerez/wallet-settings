print("Wallet Settings")

load("ext://restart_process", "docker_build_with_restart")

cfg = read_yaml(
    "tilt.yaml",
    default = read_yaml("tilt.yaml.sample"),
)

local_resource(
    "settings-build-binary",
    "make fast_build",
    deps = ["./cmd", "./internal", "./rpc/cmd", "./rpc/internal"],
)
local_resource(
    "settings-generate-protpbuf",
    "make gen-protobuf",
    deps = ["./rpc/proto/settings/settings.proto"],
)

docker_build(
    "velmie/wallet-settings-db-migration",
    ".",
    dockerfile = "Dockerfile.migrations",
    only = "migrations",
)
k8s_resource(
    "wallet-settings-db-migration",
    trigger_mode = TRIGGER_MODE_MANUAL,
    resource_deps = ["wallet-settings-db-init"],
)

wallet_settings_options = dict(
    entrypoint = "/app/service_settings",
    dockerfile = "Dockerfile.prebuild",
    port_forwards = [],
    helm_set = [],
)

if cfg["debug"]:
    wallet_settings_options["entrypoint"] = "$GOPATH/bin/dlv --continue --listen :%s --accept-multiclient --api-version=2 --headless=true exec /app/service_settings" % cfg["debug_port"]
    wallet_settings_options["dockerfile"] = "Dockerfile.debug"
    wallet_settings_options["port_forwards"] = cfg["debug_port"]
    wallet_settings_options["helm_set"] = ["containerLivenessProbe.enabled=false", "containerPorts[0].containerPort=%s" % cfg["debug_port"]]

docker_build_with_restart(
    "velmie/wallet-settings",
    ".",
    dockerfile = wallet_settings_options["dockerfile"],
    entrypoint = wallet_settings_options["entrypoint"],
    only = [
        "./build",
        "zoneinfo.zip",
    ],
    live_update = [
        sync("./build", "/app/"),
    ],
)
k8s_resource(
    "wallet-settings",
    resource_deps = ["wallet-settings-db-migration"],
    port_forwards = wallet_settings_options["port_forwards"],
)

yaml = helm(
    "./helm/wallet-settings",
    # The release name, equivalent to helm --name
    name = "wallet-settings",
    # The values file to substitute into the chart.
    values = ["./helm/values-dev.yaml"],
    set = wallet_settings_options["helm_set"],
)

k8s_yaml(yaml)
