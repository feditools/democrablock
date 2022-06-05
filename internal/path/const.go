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

	// filestore.

	// Filestore is the path for the filestore.
	Filestore = "/" + PartFilestore
	// FilestoreSubFile is the path for the filestore.
	FilestoreSubFile = "/" + VarGroup + "/" + VarHash1 + "/" + VarHash2 + "/" + VarHash3 + "/" + VarFileStoreHash + "." + VarFileStoreSuffix

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

	// VarFileStoreHashID is the id of the filestore hash variable.
	VarFileStoreHashID = "filestorehash"
	// VarFileStoreHash is the var path of the filestore hash variable.
	VarFileStoreHash = "{" + VarFileStoreHashID + ":" + reHexSHA256 + "}"
	// VarFileStoreSuffixID is the id of the filestore suffix variable.
	VarFileStoreSuffixID = "filestoresuffix"
	// VarFileStoreSuffix is the var path of the filestore suffix variable.
	VarFileStoreSuffix = "{" + VarFileStoreSuffixID + ":[a-z0-9]{3,4}}"
	// VarGroupID is the id of the group variable.
	VarGroupID = "group"
	// VarGroup is the var path of the group variable.
	VarGroup = "{" + VarGroupID + ":[a-zA-Z0-9_-]+}"
	// VarHash1ID is the id of the hash1 1 variable.
	VarHash1ID = "hash1"
	// VarHash1 is the var path of the hash 1 variable.
	VarHash1 = "{" + VarHash1ID + ":" + reHexByte + "}"
	// VarHash2ID is the id of the hash1 2 variable.
	VarHash2ID = "hash2"
	// VarHash2 is the var path of the hash 2 variable.
	VarHash2 = "{" + VarHash2ID + ":" + reHexByte + "}"
	// VarHash3ID is the id of the hash1 3 variable.
	VarHash3ID = "hash3"
	// VarHash3 is the var path of the hash 3 variable.
	VarHash3 = "{" + VarHash3ID + ":" + reHexByte + "}"
	// VarInstanceID is the id of the instance variable.
	VarInstanceID = "instance"
	// VarInstance is the var path of the instance variable.
	VarInstance = "{" + VarInstanceID + ":" + reToken + "}"
)
