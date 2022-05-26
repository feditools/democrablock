package fedihelper

import (
	"context"
)

type Instance interface {
	IsOAuthSet() bool

	GetActorURI() (actorURI string)
	GetClientID() (clientID string)
	GetClientSecret() (clientID string, err error)
	GetDomain() (domain string)
	GetServerHostname() (hostname string)
	GetSoftware() (software string)

	SetActorURI(actorURI string)
	SetClientID(clientID string)
	SetClientSecret(clientID string) (err error)
	SetDomain(domain string)
	SetServerHostname(hostname string)
	SetSoftware(software string)
}

// GenerateFediInstanceFromDomain created a Instance object by querying the apis of the federated instance.
func (f *FediHelper) GenerateFediInstanceFromDomain(ctx context.Context, domain string, instance Instance) error {
	l := logger.WithField("func", "GenerateFediInstanceFromDomain")

	// get nodeinfo endpoints from well-known location
	wkni, err := f.GetWellknownNodeInfo(ctx, domain)
	if err != nil {
		l.Errorf("get nodeinfo: %s", err.Error())

		return err
	}

	// check for nodeinfo 2.0 schema
	nodeinfoURI, err := findNodeInfo20URI(wkni)
	if err != nil {
		return err
	}
	if nodeinfoURI == nil {
		return NewError("missing nodeinfo 2.0 uri")
	}

	// get nodeinfo from
	nodeinfo, err := f.GetNodeInfo20(ctx, domain, nodeinfoURI)
	if err != nil {
		fhErr := NewErrorf("get nodeinfo 2.0: %s", err.Error())
		l.Error(fhErr.Error())

		return fhErr
	}

	// get actor uri
	webfinger, err := f.GetWellknownWebFinger(ctx, domain, domain)
	if err != nil {
		fhErr := NewErrorf("get wellknown webfinger: %s", err.Error())
		l.Error(fhErr.Error())

		return fhErr
	}
	actorURI, err := FindActorURI(webfinger)
	if err != nil {
		fhErr := NewErrorf("find actor url: %s", err.Error())
		l.Error(fhErr.Error())

		return fhErr
	}
	if actorURI == nil {
		return NewError("missing actor uri")
	}

	instance.SetActorURI(actorURI.String())
	instance.SetDomain(domain)
	instance.SetServerHostname(nodeinfoURI.Host)
	instance.SetSoftware(nodeinfo.Software.Name)

	return nil
}
