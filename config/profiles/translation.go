package profiles

import "github.com/melisource/fury_c-go-toolkit/pkg/translator"

type Translation struct {
	Fallback         translator.Locale   `json:"fallback"`
	SupportedLocales []translator.Locale `json:"supported_locales"`
	Dirname          string              `json:"dirname"`
}
