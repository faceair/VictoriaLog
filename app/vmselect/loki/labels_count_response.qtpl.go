// Code generated by qtc from "labels_count_response.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line app/vmselect/loki/labels_count_response.qtpl:1
package loki

//line app/vmselect/loki/labels_count_response.qtpl:1
import "github.com/VictoriaMetrics/VictoriaMetrics/lib/storage"

// LabelsCountResponse generates response for /api/v1/labels/count .

//line app/vmselect/loki/labels_count_response.qtpl:5
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line app/vmselect/loki/labels_count_response.qtpl:5
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line app/vmselect/loki/labels_count_response.qtpl:5
func StreamLabelsCountResponse(qw422016 *qt422016.Writer, labelEntries []storage.TagEntry) {
//line app/vmselect/loki/labels_count_response.qtpl:5
	qw422016.N().S(`{"status":"success","data":{`)
//line app/vmselect/loki/labels_count_response.qtpl:9
	for i, e := range labelEntries {
//line app/vmselect/loki/labels_count_response.qtpl:10
		qw422016.N().Q(e.Key)
//line app/vmselect/loki/labels_count_response.qtpl:10
		qw422016.N().S(`:`)
//line app/vmselect/loki/labels_count_response.qtpl:10
		qw422016.N().D(len(e.Values))
//line app/vmselect/loki/labels_count_response.qtpl:11
		if i+1 < len(labelEntries) {
//line app/vmselect/loki/labels_count_response.qtpl:11
			qw422016.N().S(`,`)
//line app/vmselect/loki/labels_count_response.qtpl:11
		}
//line app/vmselect/loki/labels_count_response.qtpl:12
	}
//line app/vmselect/loki/labels_count_response.qtpl:12
	qw422016.N().S(`}}`)
//line app/vmselect/loki/labels_count_response.qtpl:15
}

//line app/vmselect/loki/labels_count_response.qtpl:15
func WriteLabelsCountResponse(qq422016 qtio422016.Writer, labelEntries []storage.TagEntry) {
//line app/vmselect/loki/labels_count_response.qtpl:15
	qw422016 := qt422016.AcquireWriter(qq422016)
//line app/vmselect/loki/labels_count_response.qtpl:15
	StreamLabelsCountResponse(qw422016, labelEntries)
//line app/vmselect/loki/labels_count_response.qtpl:15
	qt422016.ReleaseWriter(qw422016)
//line app/vmselect/loki/labels_count_response.qtpl:15
}

//line app/vmselect/loki/labels_count_response.qtpl:15
func LabelsCountResponse(labelEntries []storage.TagEntry) string {
//line app/vmselect/loki/labels_count_response.qtpl:15
	qb422016 := qt422016.AcquireByteBuffer()
//line app/vmselect/loki/labels_count_response.qtpl:15
	WriteLabelsCountResponse(qb422016, labelEntries)
//line app/vmselect/loki/labels_count_response.qtpl:15
	qs422016 := string(qb422016.B)
//line app/vmselect/loki/labels_count_response.qtpl:15
	qt422016.ReleaseByteBuffer(qb422016)
//line app/vmselect/loki/labels_count_response.qtpl:15
	return qs422016
//line app/vmselect/loki/labels_count_response.qtpl:15
}
