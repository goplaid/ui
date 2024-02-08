package vuetifyx

import (
	"context"
	"fmt"

	"github.com/qor5/ui/vuetify"
	"github.com/qor5/web"
	h "github.com/theplant/htmlgo"
)

type PickerBuilder struct {
	value     interface{}
	label     string
	fieldName string
	comp      h.MutableAttrHTMLComponent
}

func Picker(c h.MutableAttrHTMLComponent) (r *PickerBuilder) {
	r = &PickerBuilder{comp: c}
	return
}

func (b *PickerBuilder) Value(v interface{}) (r *PickerBuilder) {
	b.value = v
	return b
}

func (b *PickerBuilder) Label(v string) (r *PickerBuilder) {
	b.label = v
	return b
}

func (b *PickerBuilder) FieldName(v string) (r *PickerBuilder) {
	b.fieldName = v
	return b
}

func (b *PickerBuilder) MarshalHTML(ctx context.Context) ([]byte, error) {
	menuLocal := fmt.Sprintf("picker_%s_menu", b.fieldName)
	valueLocal := fmt.Sprintf("picker_%s_value", b.fieldName)

	b.comp.SetAttr("@change", fmt.Sprintf(`locals.%s = false; locals.%s = $event; $plaid().form(plaidForm).fieldValue(%s, $event)`, menuLocal, valueLocal, h.JSONString(b.fieldName)))

	return web.Scope(
		vuetify.VMenu(
			web.Slot(
				vuetify.VTextField().
					Label(b.label).
					Attr(web.VField(b.fieldName, b.value)...).
					Readonly(true).
					PrependIcon("edit_calendar").
					Attr("v-model", fmt.Sprintf("locals.%s", valueLocal)).
					Attr("v-bind", "attrs").
					Attr("v-on", "on"),
			).Name("activator").Scope("{ on, attrs }"),

			b.comp,
		).Attr("v-model", fmt.Sprintf("locals.%s", menuLocal)).
			CloseOnContentClick(false).
			MaxWidth(290),
	).Init(fmt.Sprintf(`{%s: %s, %s: false}`, valueLocal, h.JSONString(b.value), menuLocal)).
		VSlot("{ locals }").
		MarshalHTML(ctx)

}
