// Code generated by qtc from "TableEdit.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/components/edit/TableEdit.html:1
package edit

//line views/components/edit/TableEdit.html:1
import (
	"github.com/kyleu/rituals/app/util"
)

//line views/components/edit/TableEdit.html:5
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/components/edit/TableEdit.html:5
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/components/edit/TableEdit.html:5
func StreamTableEditor(qw422016 *qt422016.Writer, key string, columns []*util.FieldDesc, values util.ValueMap, action string, method string, title string) {
//line views/components/edit/TableEdit.html:5
	qw422016.N().S(`
  <form action="`)
//line views/components/edit/TableEdit.html:6
	qw422016.E().S(action)
//line views/components/edit/TableEdit.html:6
	qw422016.N().S(`" method="`)
//line views/components/edit/TableEdit.html:6
	qw422016.E().S(method)
//line views/components/edit/TableEdit.html:6
	qw422016.N().S(`">
    `)
//line views/components/edit/TableEdit.html:7
	StreamTableEditorNoForm(qw422016, key, columns, values, "", "", title)
//line views/components/edit/TableEdit.html:7
	qw422016.N().S(`
  </form>
`)
//line views/components/edit/TableEdit.html:9
}

//line views/components/edit/TableEdit.html:9
func WriteTableEditor(qq422016 qtio422016.Writer, key string, columns []*util.FieldDesc, values util.ValueMap, action string, method string, title string) {
//line views/components/edit/TableEdit.html:9
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/edit/TableEdit.html:9
	StreamTableEditor(qw422016, key, columns, values, action, method, title)
//line views/components/edit/TableEdit.html:9
	qt422016.ReleaseWriter(qw422016)
//line views/components/edit/TableEdit.html:9
}

//line views/components/edit/TableEdit.html:9
func TableEditor(key string, columns []*util.FieldDesc, values util.ValueMap, action string, method string, title string) string {
//line views/components/edit/TableEdit.html:9
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/edit/TableEdit.html:9
	WriteTableEditor(qb422016, key, columns, values, action, method, title)
//line views/components/edit/TableEdit.html:9
	qs422016 := string(qb422016.B)
//line views/components/edit/TableEdit.html:9
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/edit/TableEdit.html:9
	return qs422016
//line views/components/edit/TableEdit.html:9
}

//line views/components/edit/TableEdit.html:11
func StreamTableEditorNoForm(qw422016 *qt422016.Writer, key string, columns []*util.FieldDesc, values util.ValueMap, name string, value string, title string) {
//line views/components/edit/TableEdit.html:11
	qw422016.N().S(`
  <div class="overflow full-width">
    <table class="mt min-200 expanded">
      <tbody>
        `)
//line views/components/edit/TableEdit.html:15
	StreamTableEditorNoTable(qw422016, key, columns, values)
//line views/components/edit/TableEdit.html:15
	qw422016.N().S(`
        <tr>
`)
//line views/components/edit/TableEdit.html:17
	if name == "" {
//line views/components/edit/TableEdit.html:17
		qw422016.N().S(`          <td colspan="2"><button type="submit">`)
//line views/components/edit/TableEdit.html:18
		qw422016.E().S(title)
//line views/components/edit/TableEdit.html:18
		qw422016.N().S(`</button></td>
`)
//line views/components/edit/TableEdit.html:19
	} else {
//line views/components/edit/TableEdit.html:19
		qw422016.N().S(`          <td colspan="2"><button name="`)
//line views/components/edit/TableEdit.html:20
		qw422016.E().S(name)
//line views/components/edit/TableEdit.html:20
		qw422016.N().S(`" value="`)
//line views/components/edit/TableEdit.html:20
		qw422016.E().S(value)
//line views/components/edit/TableEdit.html:20
		qw422016.N().S(`" type="submit">`)
//line views/components/edit/TableEdit.html:20
		qw422016.E().S(title)
//line views/components/edit/TableEdit.html:20
		qw422016.N().S(`</button></td>
`)
//line views/components/edit/TableEdit.html:21
	}
//line views/components/edit/TableEdit.html:21
	qw422016.N().S(`        </tr>
      </tbody>
    </table>
  </div>
`)
//line views/components/edit/TableEdit.html:26
}

//line views/components/edit/TableEdit.html:26
func WriteTableEditorNoForm(qq422016 qtio422016.Writer, key string, columns []*util.FieldDesc, values util.ValueMap, name string, value string, title string) {
//line views/components/edit/TableEdit.html:26
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/edit/TableEdit.html:26
	StreamTableEditorNoForm(qw422016, key, columns, values, name, value, title)
//line views/components/edit/TableEdit.html:26
	qt422016.ReleaseWriter(qw422016)
//line views/components/edit/TableEdit.html:26
}

//line views/components/edit/TableEdit.html:26
func TableEditorNoForm(key string, columns []*util.FieldDesc, values util.ValueMap, name string, value string, title string) string {
//line views/components/edit/TableEdit.html:26
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/edit/TableEdit.html:26
	WriteTableEditorNoForm(qb422016, key, columns, values, name, value, title)
//line views/components/edit/TableEdit.html:26
	qs422016 := string(qb422016.B)
//line views/components/edit/TableEdit.html:26
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/edit/TableEdit.html:26
	return qs422016
//line views/components/edit/TableEdit.html:26
}

//line views/components/edit/TableEdit.html:28
func StreamTableEditorNoTable(qw422016 *qt422016.Writer, key string, columns []*util.FieldDesc, values util.ValueMap) {
//line views/components/edit/TableEdit.html:29
	for _, arg := range columns {
//line views/components/edit/TableEdit.html:30
		switch arg.Type {
//line views/components/edit/TableEdit.html:31
		case "bool":
//line views/components/edit/TableEdit.html:32
			StreamBoolTable(qw422016, arg.Key, arg.Title, values.GetBoolOpt(arg.Key), 3, arg.Description)
//line views/components/edit/TableEdit.html:33
		case "textarea":
//line views/components/edit/TableEdit.html:34
			StreamTextareaTable(qw422016, arg.Key, "", arg.Title, 12, values.GetStringOpt(arg.Key), 3, arg.Description)
//line views/components/edit/TableEdit.html:35
		case "number", "int":
//line views/components/edit/TableEdit.html:36
			StreamIntTable(qw422016, arg.Key, "", arg.Title, values.GetIntOpt(arg.Key), 3, arg.Description)
//line views/components/edit/TableEdit.html:37
		case "float":
//line views/components/edit/TableEdit.html:38
			StreamFloatTable(qw422016, arg.Key, "", arg.Title, values.GetFloatOpt(arg.Key), 3, arg.Description)
//line views/components/edit/TableEdit.html:39
		default:
//line views/components/edit/TableEdit.html:40
			StreamDatalistTable(qw422016, arg.Key, "", arg.Title, values.GetStringOpt(arg.Key), arg.Choices, nil, 3, arg.Description)
//line views/components/edit/TableEdit.html:41
		}
//line views/components/edit/TableEdit.html:42
	}
//line views/components/edit/TableEdit.html:43
}

//line views/components/edit/TableEdit.html:43
func WriteTableEditorNoTable(qq422016 qtio422016.Writer, key string, columns []*util.FieldDesc, values util.ValueMap) {
//line views/components/edit/TableEdit.html:43
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/edit/TableEdit.html:43
	StreamTableEditorNoTable(qw422016, key, columns, values)
//line views/components/edit/TableEdit.html:43
	qt422016.ReleaseWriter(qw422016)
//line views/components/edit/TableEdit.html:43
}

//line views/components/edit/TableEdit.html:43
func TableEditorNoTable(key string, columns []*util.FieldDesc, values util.ValueMap) string {
//line views/components/edit/TableEdit.html:43
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/edit/TableEdit.html:43
	WriteTableEditorNoTable(qb422016, key, columns, values)
//line views/components/edit/TableEdit.html:43
	qs422016 := string(qb422016.B)
//line views/components/edit/TableEdit.html:43
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/edit/TableEdit.html:43
	return qs422016
//line views/components/edit/TableEdit.html:43
}
