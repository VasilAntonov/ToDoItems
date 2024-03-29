// Code generated by qtc from "create.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line C:/Users/vasil/go/src/ToDoTask/templates/create.qtpl:1
package templates

//line C:/Users/vasil/go/src/ToDoTask/templates/create.qtpl:1
import "net/http"

//line C:/Users/vasil/go/src/ToDoTask/templates/create.qtpl:3
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line C:/Users/vasil/go/src/ToDoTask/templates/create.qtpl:3
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line C:/Users/vasil/go/src/ToDoTask/templates/create.qtpl:4
type CreatePage struct {
	R *http.Request
}

//line C:/Users/vasil/go/src/ToDoTask/templates/create.qtpl:9
func (p *CreatePage) StreamTitle(qw422016 *qt422016.Writer) {
//line C:/Users/vasil/go/src/ToDoTask/templates/create.qtpl:9
	qw422016.N().S(`
	Creating item
`)
//line C:/Users/vasil/go/src/ToDoTask/templates/create.qtpl:11
}

//line C:/Users/vasil/go/src/ToDoTask/templates/create.qtpl:11
func (p *CreatePage) WriteTitle(qq422016 qtio422016.Writer) {
//line C:/Users/vasil/go/src/ToDoTask/templates/create.qtpl:11
	qw422016 := qt422016.AcquireWriter(qq422016)
//line C:/Users/vasil/go/src/ToDoTask/templates/create.qtpl:11
	p.StreamTitle(qw422016)
//line C:/Users/vasil/go/src/ToDoTask/templates/create.qtpl:11
	qt422016.ReleaseWriter(qw422016)
//line C:/Users/vasil/go/src/ToDoTask/templates/create.qtpl:11
}

//line C:/Users/vasil/go/src/ToDoTask/templates/create.qtpl:11
func (p *CreatePage) Title() string {
//line C:/Users/vasil/go/src/ToDoTask/templates/create.qtpl:11
	qb422016 := qt422016.AcquireByteBuffer()
//line C:/Users/vasil/go/src/ToDoTask/templates/create.qtpl:11
	p.WriteTitle(qb422016)
//line C:/Users/vasil/go/src/ToDoTask/templates/create.qtpl:11
	qs422016 := string(qb422016.B)
//line C:/Users/vasil/go/src/ToDoTask/templates/create.qtpl:11
	qt422016.ReleaseByteBuffer(qb422016)
//line C:/Users/vasil/go/src/ToDoTask/templates/create.qtpl:11
	return qs422016
//line C:/Users/vasil/go/src/ToDoTask/templates/create.qtpl:11
}

//line C:/Users/vasil/go/src/ToDoTask/templates/create.qtpl:13
func (p *CreatePage) StreamBody(qw422016 *qt422016.Writer) {
//line C:/Users/vasil/go/src/ToDoTask/templates/create.qtpl:13
	qw422016.N().S(`
	<h1>Create item page</h1>
	<div>
		<form action="/item-create" method="post" enctype="multipart/form-data">
		 	Username:<input type="text" name="username">
		 	Category:<input type="category" name="category">
		 	Title:<input type="title" name="title">
		 	<input type="submit" value="Create">
 		</form>
	</div>
`)
//line C:/Users/vasil/go/src/ToDoTask/templates/create.qtpl:23
}

//line C:/Users/vasil/go/src/ToDoTask/templates/create.qtpl:23
func (p *CreatePage) WriteBody(qq422016 qtio422016.Writer) {
//line C:/Users/vasil/go/src/ToDoTask/templates/create.qtpl:23
	qw422016 := qt422016.AcquireWriter(qq422016)
//line C:/Users/vasil/go/src/ToDoTask/templates/create.qtpl:23
	p.StreamBody(qw422016)
//line C:/Users/vasil/go/src/ToDoTask/templates/create.qtpl:23
	qt422016.ReleaseWriter(qw422016)
//line C:/Users/vasil/go/src/ToDoTask/templates/create.qtpl:23
}

//line C:/Users/vasil/go/src/ToDoTask/templates/create.qtpl:23
func (p *CreatePage) Body() string {
//line C:/Users/vasil/go/src/ToDoTask/templates/create.qtpl:23
	qb422016 := qt422016.AcquireByteBuffer()
//line C:/Users/vasil/go/src/ToDoTask/templates/create.qtpl:23
	p.WriteBody(qb422016)
//line C:/Users/vasil/go/src/ToDoTask/templates/create.qtpl:23
	qs422016 := string(qb422016.B)
//line C:/Users/vasil/go/src/ToDoTask/templates/create.qtpl:23
	qt422016.ReleaseByteBuffer(qb422016)
//line C:/Users/vasil/go/src/ToDoTask/templates/create.qtpl:23
	return qs422016
//line C:/Users/vasil/go/src/ToDoTask/templates/create.qtpl:23
}
