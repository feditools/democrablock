package fedihelper

import (
	"context"
	"net/url"

	"github.com/feditools/go-lib"
)

// GetLoginURL retrieves an oauth url for a federated instance.
func (f *FediHelper) GetLoginURL(ctx context.Context, act string) (*url.URL, error) {
	l := logger.WithField("func", "GetLoginURL")
	_, domain, err := lib.SplitAccount(act)
	if err != nil {
		l.Errorf("split account: %s", err.Error())

		return nil, err
	}

	// try to get instance from the database
	instance, err := f.GetInstanceHandler(ctx, domain)
	if err != nil {
		l.Errorf("db read: %s", err.Error())

		return nil, err
	}
	if instance != nil {
		u, err := f.loginURLForInstance(ctx, instance)
		if err != nil {
			l.Errorf("get login url: %s", err.Error())

			return nil, err
		}

		return u, nil
	}

	// get instance data from instance apis
	err = f.GenerateFediInstanceFromDomain(ctx, domain, instance)
	if err != nil {
		l.Errorf("get nodeinfo: %s", err.Error())

		return nil, err
	}
	err = f.CreateInstanceHandler(ctx, instance)
	if err != nil {
		l.Errorf("db create: %s", err.Error())

		return nil, err
	}

	u, err := f.loginURLForInstance(ctx, instance)
	if err != nil {
		l.Errorf("get login url: %s", err.Error())

		return nil, err
	}

	return u, nil
}

func (f *FediHelper) loginURLForInstance(ctx context.Context, instance Instance) (*url.URL, error) {
	l := logger.WithField("func", "loginURLForInstance")

	if _, ok := f.helpers[Software(instance.GetSoftware())]; !ok {
		return nil, NewErrorf("no helper for '%s'", instance.GetSoftware())
	}

	if !instance.IsOAuthSet() {
		clientID, clientSecret, err := f.helpers[SoftwareMastodon].RegisterApp(ctx, instance)
		if err != nil {
			fhErr := NewErrorf("registering app: %s", err.Error())
			l.Error(fhErr.Error())

			return nil, fhErr
		}
		l.Debugf("got app: %s, %s", clientID, clientSecret)
		instance.SetClientID(clientID)
		err = instance.SetClientSecret(clientSecret)
		if err != nil {
			fhErr := NewErrorf("setting secret: %s", err.Error())
			l.Error(fhErr.Error())

			return nil, fhErr
		}

		err = f.UpdateInstanceHandler(ctx, instance)
		if err != nil {
			fhErr := NewErrorf("db update: %s", err.Error())
			l.Error(fhErr.Error())

			return nil, fhErr
		}
	}

	return f.helpers[SoftwareMastodon].MakeLoginURI(ctx, instance)
}
