package bootstrap

import (
	"chi-rest/lib/utils"

	"github.com/go-chi/chi"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	"github.com/jmoiron/sqlx"
	validator "gopkg.in/go-playground/validator.v9"
	enTranslations "gopkg.in/go-playground/validator.v9/translations/en"
	idTranslations "gopkg.in/go-playground/validator.v9/translations/id"
)

// App instance of the skeleton
type App struct {
	Debug     bool
	R         *chi.Mux
	DB        *sqlx.DB
	Config    utils.Config
	Validator *Validator
}

// Validator set validator instance
type Validator struct {
	Driver     *validator.Validate
	Uni        *ut.UniversalTranslator
	Translator ut.Translator
}

// SetupValidator create new instance of validator driver
func SetupValidator(config utils.Config) *Validator {
	en := en.New()
	id := id.New()
	uni := ut.New(en, id)

	transEN, _ := uni.GetTranslator("en")
	transID, _ := uni.GetTranslator("id")

	validatorDriver := validator.New()

	enTranslations.RegisterDefaultTranslations(validatorDriver, transEN)
	idTranslations.RegisterDefaultTranslations(validatorDriver, transID)

	var trans ut.Translator
	switch config.GetString("app.locale") {
	case "en":
		trans = transEN
	case "id":
		trans = transID
	}

	return &Validator{Driver: validatorDriver, Uni: uni, Translator: trans}
}
