package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VBreadcrumbsBuilder struct {
	tag *h.HTMLTagBuilder
}

func VBreadcrumbs(children ...h.HTMLComponent) (r *VBreadcrumbsBuilder) {
	r = &VBreadcrumbsBuilder{
		tag: h.Tag("v-breadcrumbs").Children(children...),
	}
	return
}

func (b *VBreadcrumbsBuilder) Dark(v bool) (r *VBreadcrumbsBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
	return b
}

func (b *VBreadcrumbsBuilder) Divider(v string) (r *VBreadcrumbsBuilder) {
	b.tag.Attr("divider", v)
	return b
}

func (b *VBreadcrumbsBuilder) Items(v interface{}) (r *VBreadcrumbsBuilder) {
	b.tag.Attr(":items", h.JSONString(v))
	return b
}

func (b *VBreadcrumbsBuilder) Large(v bool) (r *VBreadcrumbsBuilder) {
	b.tag.Attr(":large", fmt.Sprint(v))
	return b
}

func (b *VBreadcrumbsBuilder) Light(v bool) (r *VBreadcrumbsBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VBreadcrumbsBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VBreadcrumbsBuilder) Attr(vs ...interface{}) (r *VBreadcrumbsBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VBreadcrumbsBuilder) Children(children ...h.HTMLComponent) (r *VBreadcrumbsBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VBreadcrumbsBuilder) AppendChildren(children ...h.HTMLComponent) (r *VBreadcrumbsBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VBreadcrumbsBuilder) PrependChildren(children ...h.HTMLComponent) (r *VBreadcrumbsBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VBreadcrumbsBuilder) Class(names ...string) (r *VBreadcrumbsBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VBreadcrumbsBuilder) ClassIf(name string, add bool) (r *VBreadcrumbsBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VBreadcrumbsBuilder) On(name string, value string) (r *VBreadcrumbsBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VBreadcrumbsBuilder) Bind(name string, value string) (r *VBreadcrumbsBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VBreadcrumbsBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
