package webapp

import (
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"github.com/feditools/democrablock/web"
	"io/ioutil"
)

// signature caching

func getSignature(path string) (string, error) {
	l := logger.WithField("func", "getSignature")

	file, err := web.Files.Open(path)
	if err != nil {
		l.Errorf("opening file: %s", err.Error())

		return "", err
	}

	// read it
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}

	// hash it
	h := sha512.New384()
	_, err = h.Write(data)
	if err != nil {
		return "", err
	}
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))

	return fmt.Sprintf("sha384-%s", signature), nil
}

func (m *Module) getSignatureCached(path string) (string, error) {
	if sig, ok := m.readCachedSignature(path); ok {
		return sig, nil
	}
	sig, err := getSignature(path)
	if err != nil {
		return "", err
	}
	m.writeCachedSignature(path, sig)

	return sig, nil
}

func (m *Module) readCachedSignature(path string) (string, bool) {
	m.sigCacheLock.RLock()
	val, ok := m.sigCache[path]
	m.sigCacheLock.RUnlock()

	return val, ok
}

func (m *Module) writeCachedSignature(path string, sig string) {
	m.sigCacheLock.Lock()
	m.sigCache[path] = sig
	m.sigCacheLock.Unlock()
}
