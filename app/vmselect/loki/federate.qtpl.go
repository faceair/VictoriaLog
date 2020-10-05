// Code generated by qtc from "federate.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line app/vmselect/loki/federate.qtpl:1
package loki

//line app/vmselect/loki/federate.qtpl:1
import (
	"github.com/VictoriaMetrics/VictoriaMetrics/app/vmselect/netstorage"
)

// Federate writes rs in /federate format.// See https://prometheus.io/docs/prometheus/latest/federation/

//line app/vmselect/loki/federate.qtpl:9
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line app/vmselect/loki/federate.qtpl:9
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line app/vmselect/loki/federate.qtpl:9
func StreamFederate(qw422016 *qt422016.Writer, rs *netstorage.Result) {
//line app/vmselect/loki/federate.qtpl:10
	if len(rs.Timestamps) == 0 || len(rs.Datas) == 0 {
//line app/vmselect/loki/federate.qtpl:10
		return
//line app/vmselect/loki/federate.qtpl:10
	}
//line app/vmselect/loki/federate.qtpl:11
	streamprometheusMetricName(qw422016, &rs.MetricName)
//line app/vmselect/loki/federate.qtpl:11
	qw422016.N().S(` `)
//line app/vmselect/loki/federate.qtpl:12
	qw422016.N().Z(rs.Datas[len(rs.Datas)-1])
//line app/vmselect/loki/federate.qtpl:12
	qw422016.N().S(` `)
//line app/vmselect/loki/federate.qtpl:13
	qw422016.N().DL(rs.Timestamps[len(rs.Timestamps)-1])
//line app/vmselect/loki/federate.qtpl:13
	qw422016.N().S(`
`)
//line app/vmselect/loki/federate.qtpl:14
}

//line app/vmselect/loki/federate.qtpl:14
func WriteFederate(qq422016 qtio422016.Writer, rs *netstorage.Result) {
//line app/vmselect/loki/federate.qtpl:14
	qw422016 := qt422016.AcquireWriter(qq422016)
//line app/vmselect/loki/federate.qtpl:14
	StreamFederate(qw422016, rs)
//line app/vmselect/loki/federate.qtpl:14
	qt422016.ReleaseWriter(qw422016)
//line app/vmselect/loki/federate.qtpl:14
}

//line app/vmselect/loki/federate.qtpl:14
func Federate(rs *netstorage.Result) string {
//line app/vmselect/loki/federate.qtpl:14
	qb422016 := qt422016.AcquireByteBuffer()
//line app/vmselect/loki/federate.qtpl:14
	WriteFederate(qb422016, rs)
//line app/vmselect/loki/federate.qtpl:14
	qs422016 := string(qb422016.B)
//line app/vmselect/loki/federate.qtpl:14
	qt422016.ReleaseByteBuffer(qb422016)
//line app/vmselect/loki/federate.qtpl:14
	return qs422016
//line app/vmselect/loki/federate.qtpl:14
}
