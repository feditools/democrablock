package path

const (
	// files.

	// FileDefaultCSS is the css document applies to all pages.
	FileDefaultCSS = StaticCSS + "/default.min.css"

	// parts.

	// PartStatic is used in a path for static files.
	PartStatic = "static"

	// paths.

	// Static is the path for static files.
	Static = "/" + PartStatic + "/"
	// StaticCSS is the path.
	StaticCSS = Static + "css"
)
