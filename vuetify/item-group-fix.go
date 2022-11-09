package vuetify

import (
	"github.com/qor5/web"
)

func (b *VItemGroupBuilder) FieldName(v string) (r *VItemGroupBuilder) {
	b.tag.Attr(web.VFieldName(v)...)
	return b
}
