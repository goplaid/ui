package vuetify

import (
	"github.com/qor5/web"
)

func (b *VCheckboxBuilder) FieldName(v string) (r *VCheckboxBuilder) {
	b.tag.Attr(web.VFieldName(v)...)
	return b
}

func (b *VCheckboxBuilder) ErrorMessages(v ...string) (r *VCheckboxBuilder) {
	SetErrorMessages(b.tag, v)
	return b
}
