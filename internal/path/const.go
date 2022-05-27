package path

const (
	// files.

	// FileDefaultCSS is the css document applies to all pages.
	FileDefaultCSS = StaticCSS + "/default.min.css"
	// FileErrorCSS is the css document applies to the error page.
	FileErrorCSS = StaticCSS + "/error.min.css"
	// FileLoginCSS is the css document applies to the login page.
	FileLoginCSS = StaticCSS + "/login.min.css"

	// parts.

	// PartCallback is used in a path for callback.
	PartCallback = "callback"
	// PartList is used in a path for a list.
	PartList = "list"
	// PartLogin is used in a path for login.
	PartLogin = "login"
	// PartLogout is used in a path for logout.
	PartLogout = "logout"
	// PartOauth is used in a path for oauth.
	PartOauth = "oauth"
	// PartStatic is used in a path for static files.
	PartStatic = "static"

	// paths.

	// CallbackOauth is the path for an oauth callback.
	CallbackOauth = "/" + PartCallback + "/" + PartOauth + "/" + VarInstance
	// Home is the path for the home page.
	Home = "/"
	// List is the path for the block list.
	List = "/" + PartList
	// Login is the path for the login page.
	Login = "/" + PartLogin
	// Logout is the path for the logout page.
	Logout = "/" + PartLogout
	// Static is the path for static files.
	Static = "/" + PartStatic + "/"
	// StaticCSS is the path.
	StaticCSS = Static + "css"

	// vars.

	// VarInstanceID is the id of the instance variable.
	VarInstanceID = "instance"
	// VarInstance is the var path of the instance variable.
	VarInstance = "{" + VarInstanceID + ":" + reToken + "}"
)
