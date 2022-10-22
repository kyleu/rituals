// Code generated by qtc from "Package.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/components/fieldview/Package.html:2
package fieldview

//line views/components/fieldview/Package.html:2
import (
	"github.com/kyleu/rituals/app/util"
)

//line views/components/fieldview/Package.html:6
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/components/fieldview/Package.html:6
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/components/fieldview/Package.html:6
func StreamPackage(qw422016 *qt422016.Writer, v util.Pkg) {
//line views/components/fieldview/Package.html:7
	for idx, x := range v {
//line views/components/fieldview/Package.html:8
		qw422016.E().S(x)
//line views/components/fieldview/Package.html:8
		if len(v) < idx {
//line views/components/fieldview/Package.html:8
			qw422016.N().S(`/`)
//line views/components/fieldview/Package.html:8
		}
//line views/components/fieldview/Package.html:9
	}
//line views/components/fieldview/Package.html:10
}

//line views/components/fieldview/Package.html:10
func WritePackage(qq422016 qtio422016.Writer, v util.Pkg) {
//line views/components/fieldview/Package.html:10
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/fieldview/Package.html:10
	StreamPackage(qw422016, v)
//line views/components/fieldview/Package.html:10
	qt422016.ReleaseWriter(qw422016)
//line views/components/fieldview/Package.html:10
}

//line views/components/fieldview/Package.html:10
func Package(v util.Pkg) string {
//line views/components/fieldview/Package.html:10
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/fieldview/Package.html:10
	WritePackage(qb422016, v)
//line views/components/fieldview/Package.html:10
	qs422016 := string(qb422016.B)
//line views/components/fieldview/Package.html:10
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/fieldview/Package.html:10
	return qs422016
//line views/components/fieldview/Package.html:10
}
