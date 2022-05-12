package path

const (
	// files.

	// FileDefaultCSS is the css document applies to all pages.
	FileDefaultCSS = StaticCSS + "/default.min.css"
	// FileErrorCSS is the css document applies to the error page.
	FileErrorCSS = StaticCSS + "/error.min.css"

	// parts.

	// PartStatic is used in a path for static files.
	PartStatic = "static"
	// PartList is used in a path for a list.
	PartList = "list"
	// PartLogin is used in a path for login.
	PartLogin = "login"

	// paths.

	// Home is the path for the home page.
	Home = "/"
	// List is the path for the block list.
	List = "/" + PartList
	// Login is the path for the login page.
	Login = "/" + PartLogin
	// Static is the path for static files.
	Static = "/" + PartStatic + "/"
	// StaticCSS is the path.
	StaticCSS = Static + "css"
)
