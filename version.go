package goexprfinder

var (
	version    = "0.0.0"
	preRelease = "alpha"
	build      = ""
)

func Version() string {
	ver := version
	if preRelease != "" {
		ver += "-" + preRelease
	}
	if build != "" {
		ver += "+" + build
	}

	return ver
}
