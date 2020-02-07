package lifecycle

var POSIXLaunchEnv = map[string][]string{
	"bin": {"PATH"},
	"lib": {"LD_LIBRARY_PATH"},
}
