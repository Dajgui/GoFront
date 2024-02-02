// Code generated by qtc from "mainpage.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// Main page template. Implements BasePage methods.
//

//line templates/mainpage.qtpl:3
package templates

//line templates/mainpage.qtpl:3
import "github.com/valyala/fasthttp"

//line templates/mainpage.qtpl:5
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line templates/mainpage.qtpl:5
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line templates/mainpage.qtpl:6
type MainPage struct {
	CTX *fasthttp.RequestCtx
}

//line templates/mainpage.qtpl:12
func (p *MainPage) StreamTitle(qw422016 *qt422016.Writer) {
//line templates/mainpage.qtpl:12
	qw422016.N().S(`
	This is the main page
`)
//line templates/mainpage.qtpl:14
}

//line templates/mainpage.qtpl:14
func (p *MainPage) WriteTitle(qq422016 qtio422016.Writer) {
//line templates/mainpage.qtpl:14
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/mainpage.qtpl:14
	p.StreamTitle(qw422016)
//line templates/mainpage.qtpl:14
	qt422016.ReleaseWriter(qw422016)
//line templates/mainpage.qtpl:14
}

//line templates/mainpage.qtpl:14
func (p *MainPage) Title() string {
//line templates/mainpage.qtpl:14
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/mainpage.qtpl:14
	p.WriteTitle(qb422016)
//line templates/mainpage.qtpl:14
	qs422016 := string(qb422016.B)
//line templates/mainpage.qtpl:14
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/mainpage.qtpl:14
	return qs422016
//line templates/mainpage.qtpl:14
}

//line templates/mainpage.qtpl:17
func (p *MainPage) StreamBody(qw422016 *qt422016.Writer) {
//line templates/mainpage.qtpl:17
	qw422016.N().S(`
	<h1>Main page</h1>
	<div>
		Click links below:
		<ul>
			<li><a href="/table?rowsCount=42">Table page</a></li>
			<li><a href="/unknown-page">Error page</a></li>
		</ul>
	</div>
	<div>
		Some info about you:<br/>
		IP: <b>`)
//line templates/mainpage.qtpl:28
	qw422016.E().S(p.CTX.RemoteIP().String())
//line templates/mainpage.qtpl:28
	qw422016.N().S(`</b><br/>
		User-Agent: <b>`)
//line templates/mainpage.qtpl:29
	qw422016.E().Z(p.CTX.UserAgent())
//line templates/mainpage.qtpl:29
	qw422016.N().S(`</b><br/>
	</div>
`)
//line templates/mainpage.qtpl:31
}

//line templates/mainpage.qtpl:31
func (p *MainPage) WriteBody(qq422016 qtio422016.Writer) {
//line templates/mainpage.qtpl:31
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/mainpage.qtpl:31
	p.StreamBody(qw422016)
//line templates/mainpage.qtpl:31
	qt422016.ReleaseWriter(qw422016)
//line templates/mainpage.qtpl:31
}

//line templates/mainpage.qtpl:31
func (p *MainPage) Body() string {
//line templates/mainpage.qtpl:31
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/mainpage.qtpl:31
	p.WriteBody(qb422016)
//line templates/mainpage.qtpl:31
	qs422016 := string(qb422016.B)
//line templates/mainpage.qtpl:31
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/mainpage.qtpl:31
	return qs422016
//line templates/mainpage.qtpl:31
}
