package bpdf

import (
	"github.com/pchchv/bpdf/core"
	"github.com/pchchv/bpdf/metrics"
)

type MetricsDecorator struct {
	addRowsTime    []*metrics.Time
	addRowTime     []*metrics.Time
	addAutoRowTime []*metrics.Time
	addPageTime    []*metrics.Time
	headerTime     *metrics.Time
	footerTime     *metrics.Time
	generateTime   *metrics.Time
	structureTime  *metrics.Time
	inner          core.BPDF
}
