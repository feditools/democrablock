package template

import (
	"html/template"
	"io/ioutil"
	"strings"

	"github.com/feditools/democrablock/internal/models"
	"github.com/feditools/democrablock/internal/token"
	"github.com/feditools/democrablock/web"
	"github.com/feditools/go-lib/language"
	libtemplate "github.com/feditools/go-lib/template"
)

const templateDir = "template"

// InitTemplate are the functions a template implementing Common will have.
type InitTemplate interface {
	AddHeadLink(l libtemplate.HeadLink)
	AddFooterScript(s libtemplate.Script)
	SetAccount(account *models.FediAccount)
	SetLanguage(l string)
	SetLocalizer(l *language.Localizer)
	SetLogoSrc(dark, light string)
	SetNavbar(nodes Navbar)
}

// New creates a new template.
func New(t *token.Tokenizer) (*template.Template, error) {
	tpl, err := libtemplate.New(template.FuncMap{
		"token": t.GetToken,
	})
	if err != nil {
		return nil, err
	}

	dir, err := web.Files.ReadDir(templateDir)
	if err != nil {
		panic(err)
	}
	for _, d := range dir {
		filePath := templateDir + "/" + d.Name()
		if d.IsDir() || !strings.HasSuffix(d.Name(), ".gohtml") {
			continue
		}

		// open it
		file, err := web.Files.Open(filePath)
		if err != nil {
			return nil, err
		}

		// read it
		tmplData, err := ioutil.ReadAll(file)
		if err != nil {
			return nil, err
		}

		// It can now be parsed as a string.
		_, err = tpl.Parse(string(tmplData))
		if err != nil {
			return nil, err
		}
	}

	return tpl, nil
}
