package locale

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type Locale struct {
	Tag     language.Tag
	printer *message.Printer
}

func (l *Locale) Translate(key message.Reference, args ...interface{}) string {
	return l.printer.Sprintf(key, args...)
}
