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

	// PartAccounts is used in a path for accounts.
	PartAccounts = "accounts"
	// PartAdmin is used in a path for administrative tasks.
	PartAdmin = "admin"
	// PartCallback is used in a path for callback.
	PartCallback = "callback"
	// PartFediverse is used in a path for federated things.
	PartFediverse = "fedi"
	// PartFilestore is used in a path for the filestore.
	PartFilestore = "filestore"
	// PartInstances is used in a path for instances.
	PartInstances = "instances"
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
	// PartSystem is used in a path for system things.
	PartSystem = "system"

	// paths.

	// Admin is the path for the admin page.
	Admin = "/" + PartAdmin

	// AdminFediverse is the path for the fediverse admin page.
	AdminFediverse = Admin + AdminSubFediverse
	// AdminFediverseAccounts is the path for the fediverse admin page.
	AdminFediverseAccounts = Admin + AdminSubFediverseAccounts
	// AdminFediverseInstances is the path for the fediverse instances page.
	AdminFediverseInstances = Admin + AdminSubFediverseInstances

	// AdminSubFediverse is the sub path for the fediverse admin page.
	AdminSubFediverse = "/" + PartFediverse
	// AdminSubFediverseAccounts is the sub path for the fediverse admin accounts page.
	AdminSubFediverseAccounts = AdminSubFediverse + "/" + PartAccounts
	// AdminSubFediverseInstances is the sub path for the fediverse admin instances page.
	AdminSubFediverseInstances = AdminSubFediverse + "/" + PartInstances

	// AdminSystem is the path for the system admin page.
	AdminSystem = Admin + AdminSubSystem

	// AdminSubSystem is the sub path for the system admin page.
	AdminSubSystem = "/" + PartSystem

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
