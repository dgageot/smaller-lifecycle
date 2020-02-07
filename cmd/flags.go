package cmd

const (
	DefaultLayersDir   = "/layers"
	DefaultAppDir      = "/workspace"
	DefaultProcessType = "web"

	EnvLayersDir         = "CNB_LAYERS_DIR"
	EnvAppDir            = "CNB_APP_DIR"
	EnvProcessType       = "CNB_PROCESS_TYPE"
	EnvProcessTypeLegacy = "PACK_PROCESS_TYPE" // deprecated
)
