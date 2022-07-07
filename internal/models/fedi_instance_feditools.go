package models

func (f *FediInstance) GetActorURI() (actorURI string) {
	return f.ActorURI
}

func (f *FediInstance) GetDomain() (domain string) {
	return f.Domain
}

func (f *FediInstance) GetID() (instanceID int64) {
	return f.ID
}

func (f *FediInstance) GetServerHostname() (hostname string) {
	return f.ServerHostname
}

func (f *FediInstance) GetSoftware() (software string) {
	return f.Software
}

func (f *FediInstance) SetActorURI(actorURI string) {
	f.ActorURI = actorURI
}

func (f *FediInstance) SetDomain(domain string) {
	f.Domain = domain
}

func (f *FediInstance) SetServerHostname(hostname string) {
	f.ServerHostname = hostname
}

func (f *FediInstance) SetSoftware(software string) {
	f.Software = software
}
