package vuetify

import (
	"github.com/qor5/web"
)

func (b *VSwitchBuilder) ErrorMessages(v ...string) (r *VSwitchBuilder) {
	SetErrorMessages(b.tag, v)
	return b
}

func (b *VSwitchBuilder) FieldName(v string) (r *VSwitchBuilder) {
	b.tag.Attr(web.VFieldName(v)...)
	return b
}
