// Code generated by qtc from "TableHeader.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/components/TableHeader.html:1
package components

//line views/components/TableHeader.html:1
import (
	"net/url"

	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/lib/filter"
)

//line views/components/TableHeader.html:8
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/components/TableHeader.html:8
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/components/TableHeader.html:8
func StreamTableHeader(qw422016 *qt422016.Writer, section string, key string, title string, params *filter.Params, icon string, u *url.URL, tooltip string, sortable bool, cls string, resizable bool, ps *cutil.PageState) {
//line views/components/TableHeader.html:8
	qw422016.N().S(`<th class="`)
//line views/components/TableHeader.html:9
	if cls != `` {
//line views/components/TableHeader.html:9
		qw422016.E().S(cls)
//line views/components/TableHeader.html:9
		qw422016.N().S(` `)
//line views/components/TableHeader.html:9
	}
//line views/components/TableHeader.html:9
	qw422016.N().S(`no-padding" scope="col"><div class="`)
//line views/components/TableHeader.html:10
	if resizable {
//line views/components/TableHeader.html:10
		qw422016.N().S(`resize`)
//line views/components/TableHeader.html:10
	} else {
//line views/components/TableHeader.html:10
		qw422016.N().S(`noresize`)
//line views/components/TableHeader.html:10
	}
//line views/components/TableHeader.html:10
	qw422016.N().S(`">`)
//line views/components/TableHeader.html:11
	if !sortable {
//line views/components/TableHeader.html:11
		qw422016.N().S(`<div title="`)
//line views/components/TableHeader.html:12
		qw422016.E().S(tooltip)
//line views/components/TableHeader.html:12
		qw422016.N().S(`">`)
//line views/components/TableHeader.html:13
		if icon != "" {
//line views/components/TableHeader.html:14
			qw422016.N().S(` `)
//line views/components/TableHeader.html:15
			StreamSVGRef(qw422016, icon, 16, 16, "icon-block", ps)
//line views/components/TableHeader.html:16
		}
//line views/components/TableHeader.html:17
		qw422016.E().S(title)
//line views/components/TableHeader.html:17
		qw422016.N().S(`</div>`)
//line views/components/TableHeader.html:19
	} else if params == nil {
//line views/components/TableHeader.html:20
		streamthNormal(qw422016, section, key, title, params, icon, u, tooltip, ps)
//line views/components/TableHeader.html:21
	} else {
//line views/components/TableHeader.html:22
		o := params.GetOrdering(key)

//line views/components/TableHeader.html:23
		if o == nil {
//line views/components/TableHeader.html:24
			streamthNormal(qw422016, section, key, title, params, icon, u, tooltip, ps)
//line views/components/TableHeader.html:25
		} else {
//line views/components/TableHeader.html:26
			streamthSorted(qw422016, o.Asc, section, key, title, params, icon, u, tooltip, ps)
//line views/components/TableHeader.html:27
		}
//line views/components/TableHeader.html:28
	}
//line views/components/TableHeader.html:28
	qw422016.N().S(`</div></th>`)
//line views/components/TableHeader.html:31
}

//line views/components/TableHeader.html:31
func WriteTableHeader(qq422016 qtio422016.Writer, section string, key string, title string, params *filter.Params, icon string, u *url.URL, tooltip string, sortable bool, cls string, resizable bool, ps *cutil.PageState) {
//line views/components/TableHeader.html:31
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/TableHeader.html:31
	StreamTableHeader(qw422016, section, key, title, params, icon, u, tooltip, sortable, cls, resizable, ps)
//line views/components/TableHeader.html:31
	qt422016.ReleaseWriter(qw422016)
//line views/components/TableHeader.html:31
}

//line views/components/TableHeader.html:31
func TableHeader(section string, key string, title string, params *filter.Params, icon string, u *url.URL, tooltip string, sortable bool, cls string, resizable bool, ps *cutil.PageState) string {
//line views/components/TableHeader.html:31
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/TableHeader.html:31
	WriteTableHeader(qb422016, section, key, title, params, icon, u, tooltip, sortable, cls, resizable, ps)
//line views/components/TableHeader.html:31
	qs422016 := string(qb422016.B)
//line views/components/TableHeader.html:31
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/TableHeader.html:31
	return qs422016
//line views/components/TableHeader.html:31
}

//line views/components/TableHeader.html:33
func StreamTableHeaderSimple(qw422016 *qt422016.Writer, section string, key string, title string, tooltip string, params *filter.Params, u *url.URL, ps *cutil.PageState) {
//line views/components/TableHeader.html:34
	StreamTableHeader(qw422016, section, key, title, params, "", u, tooltip, u != nil, "", false, ps)
//line views/components/TableHeader.html:35
}

//line views/components/TableHeader.html:35
func WriteTableHeaderSimple(qq422016 qtio422016.Writer, section string, key string, title string, tooltip string, params *filter.Params, u *url.URL, ps *cutil.PageState) {
//line views/components/TableHeader.html:35
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/TableHeader.html:35
	StreamTableHeaderSimple(qw422016, section, key, title, tooltip, params, u, ps)
//line views/components/TableHeader.html:35
	qt422016.ReleaseWriter(qw422016)
//line views/components/TableHeader.html:35
}

//line views/components/TableHeader.html:35
func TableHeaderSimple(section string, key string, title string, tooltip string, params *filter.Params, u *url.URL, ps *cutil.PageState) string {
//line views/components/TableHeader.html:35
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/TableHeader.html:35
	WriteTableHeaderSimple(qb422016, section, key, title, tooltip, params, u, ps)
//line views/components/TableHeader.html:35
	qs422016 := string(qb422016.B)
//line views/components/TableHeader.html:35
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/TableHeader.html:35
	return qs422016
//line views/components/TableHeader.html:35
}

//line views/components/TableHeader.html:37
func streamthNormal(qw422016 *qt422016.Writer, section string, key string, title string, params *filter.Params, icon string, u *url.URL, tooltip string, ps *cutil.PageState) {
//line views/components/TableHeader.html:37
	qw422016.N().S(`<a class="sort-hover" href="?`)
//line views/components/TableHeader.html:38
	qw422016.N().S(params.CloneOrdering(&filter.Ordering{Column: key, Asc: true}).ToQueryString(u))
//line views/components/TableHeader.html:38
	qw422016.N().S(`" title="`)
//line views/components/TableHeader.html:38
	qw422016.E().S(tooltip)
//line views/components/TableHeader.html:38
	qw422016.N().S(`"><div class="sort-icon" title="click to sort by this column, ascending">`)
//line views/components/TableHeader.html:39
	StreamSVGRef(qw422016, `down`, 0, 0, ``, ps)
//line views/components/TableHeader.html:39
	qw422016.N().S(`</div><div class="sort-title">`)
//line views/components/TableHeader.html:41
	if icon != "" {
//line views/components/TableHeader.html:42
		qw422016.N().S(` `)
//line views/components/TableHeader.html:43
		StreamSVGRef(qw422016, icon, 16, 16, "icon-block", ps)
//line views/components/TableHeader.html:44
	}
//line views/components/TableHeader.html:45
	qw422016.E().S(title)
//line views/components/TableHeader.html:45
	qw422016.N().S(`</div></a>`)
//line views/components/TableHeader.html:48
}

//line views/components/TableHeader.html:48
func writethNormal(qq422016 qtio422016.Writer, section string, key string, title string, params *filter.Params, icon string, u *url.URL, tooltip string, ps *cutil.PageState) {
//line views/components/TableHeader.html:48
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/TableHeader.html:48
	streamthNormal(qw422016, section, key, title, params, icon, u, tooltip, ps)
//line views/components/TableHeader.html:48
	qt422016.ReleaseWriter(qw422016)
//line views/components/TableHeader.html:48
}

//line views/components/TableHeader.html:48
func thNormal(section string, key string, title string, params *filter.Params, icon string, u *url.URL, tooltip string, ps *cutil.PageState) string {
//line views/components/TableHeader.html:48
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/TableHeader.html:48
	writethNormal(qb422016, section, key, title, params, icon, u, tooltip, ps)
//line views/components/TableHeader.html:48
	qs422016 := string(qb422016.B)
//line views/components/TableHeader.html:48
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/TableHeader.html:48
	return qs422016
//line views/components/TableHeader.html:48
}

//line views/components/TableHeader.html:50
func streamthSorted(qw422016 *qt422016.Writer, asc bool, section string, key string, title string, params *filter.Params, icon string, u *url.URL, tooltip string, ps *cutil.PageState) {
//line views/components/TableHeader.html:52
	ascStr := "ascending"
	dirStr := "up"
	if asc {
		ascStr = "descending"
		dirStr = "down"
	}

//line views/components/TableHeader.html:58
	qw422016.N().S(`<a href="?`)
//line views/components/TableHeader.html:59
	qw422016.N().S(params.CloneOrdering(&filter.Ordering{Column: key, Asc: !asc}).ToQueryString(u))
//line views/components/TableHeader.html:59
	qw422016.N().S(`" title="`)
//line views/components/TableHeader.html:59
	qw422016.E().S(tooltip)
//line views/components/TableHeader.html:59
	qw422016.N().S(`"><div class="sort-icon" title="click to sort by this column,`)
//line views/components/TableHeader.html:60
	qw422016.N().S(` `)
//line views/components/TableHeader.html:60
	qw422016.E().S(ascStr)
//line views/components/TableHeader.html:60
	qw422016.N().S(`">`)
//line views/components/TableHeader.html:60
	StreamSVGRef(qw422016, dirStr, 0, 0, ``, ps)
//line views/components/TableHeader.html:60
	qw422016.N().S(`</div><div class="sort-title">`)
//line views/components/TableHeader.html:62
	if icon != "" {
//line views/components/TableHeader.html:63
		qw422016.N().S(` `)
//line views/components/TableHeader.html:64
		StreamSVGRef(qw422016, icon, 16, 16, "icon-block", ps)
//line views/components/TableHeader.html:65
	}
//line views/components/TableHeader.html:66
	qw422016.E().S(title)
//line views/components/TableHeader.html:66
	qw422016.N().S(`</div></a>`)
//line views/components/TableHeader.html:69
}

//line views/components/TableHeader.html:69
func writethSorted(qq422016 qtio422016.Writer, asc bool, section string, key string, title string, params *filter.Params, icon string, u *url.URL, tooltip string, ps *cutil.PageState) {
//line views/components/TableHeader.html:69
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/TableHeader.html:69
	streamthSorted(qw422016, asc, section, key, title, params, icon, u, tooltip, ps)
//line views/components/TableHeader.html:69
	qt422016.ReleaseWriter(qw422016)
//line views/components/TableHeader.html:69
}

//line views/components/TableHeader.html:69
func thSorted(asc bool, section string, key string, title string, params *filter.Params, icon string, u *url.URL, tooltip string, ps *cutil.PageState) string {
//line views/components/TableHeader.html:69
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/TableHeader.html:69
	writethSorted(qb422016, asc, section, key, title, params, icon, u, tooltip, ps)
//line views/components/TableHeader.html:69
	qs422016 := string(qb422016.B)
//line views/components/TableHeader.html:69
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/TableHeader.html:69
	return qs422016
//line views/components/TableHeader.html:69
}
