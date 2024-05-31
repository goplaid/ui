package vuetifyx

import (
	"context"

	"github.com/qor5/ui/v3/vuetify"
	h "github.com/theplant/htmlgo"
)

type VXAutocompleteBuilder struct {
	tag           *h.HTMLTagBuilder
	selectedItems interface{}
	items         interface{}
}

func VXAutocomplete(children ...h.HTMLComponent) (r *VXAutocompleteBuilder) {
	r = &VXAutocompleteBuilder{
		tag: h.Tag("vx-autocomplete").Children(children...),
	}
	r.Multiple(true).Chips(true).DeletableChips(true).Clearable(true)

	return
}

func (b *VXAutocompleteBuilder) ErrorMessages(v ...string) (r *VXAutocompleteBuilder) {
	vuetify.SetErrorMessages(b.tag, v)
	return b
}

func (b *VXAutocompleteBuilder) SelectedItems(v interface{}) (r *VXAutocompleteBuilder) {
	b.selectedItems = v
	return b
}

func (b *VXAutocompleteBuilder) HasIcon(v bool) (r *VXAutocompleteBuilder) {
	b.tag.Attr("has-icon", v)
	return b
}

func (b *VXAutocompleteBuilder) Sorting(v bool) (r *VXAutocompleteBuilder) {
	b.tag.Attr("sorting", v)
	return b
}

func (b *VXAutocompleteBuilder) Variant(v string) (r *VXAutocompleteBuilder) {
	b.tag.Attr("variant", v)
	return b
}

func (b *VXAutocompleteBuilder) Density(v string) (r *VXAutocompleteBuilder) {
	b.tag.Attr("density", v)
	return b
}

func (b *VXAutocompleteBuilder) Items(v interface{}) (r *VXAutocompleteBuilder) {
	b.items = v
	return b
}

func (b *VXAutocompleteBuilder) ChipColor(v string) (r *VXAutocompleteBuilder) {
	b.tag.Attr("chip-color", v)
	return b
}

func (b *VXAutocompleteBuilder) ChipTextColor(v string) (r *VXAutocompleteBuilder) {
	b.tag.Attr("chip-text-color", v)
	return b
}

func (b *VXAutocompleteBuilder) SetDataSource(ds *AutocompleteDataSource) (r *VXAutocompleteBuilder) {
	b.tag.Attr("remote-url", ds.RemoteURL)
	b.tag.Attr("event-name", ds.EventName)
	b.tag.Attr("is-paging", ds.IsPaging)
	b.tag.Attr("has-icon", ds.HasIcon)
	return b
}

func (b *VXAutocompleteBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	if b.items == nil {
		b.items = b.selectedItems
	}
	b.tag.Attr(":items", b.items)
	b.tag.Attr(":selected-items", b.selectedItems)
	return b.tag.MarshalHTML(ctx)
}

type AutocompleteDataSource struct {
	RemoteURL string `json:"remote-url"`
	EventName string `json:"event-name"`
	IsPaging  bool   `json:"is-paging"`
	HasIcon   bool   `json:"has-icon"`
}
