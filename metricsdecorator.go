package bpdf

import (
	"github.com/pchchv/bpdf/core"
	"github.com/pchchv/bpdf/core/entity"
	"github.com/pchchv/bpdf/internal/time"
	"github.com/pchchv/bpdf/metrics"
	"github.com/pchchv/bpdf/node"
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

// GetStructure decorates the GetStructure method of bpdf instance.
func (m *MetricsDecorator) GetStructure() *node.Node[core.Structure] {
	var tree *node.Node[core.Structure]
	timeSpent := time.GetTimeSpent(func() {
		tree = m.inner.GetStructure()
	})
	m.structureTime = timeSpent

	return tree
}

// GetCurrentConfig decorates the GetCurrentConfig method of bpdf instance.
func (m *MetricsDecorator) GetCurrentConfig() *entity.Config {
	return m.inner.GetCurrentConfig()
}

// Generate decorates the Generate method of bpdf instance.
func (m *MetricsDecorator) Generate() (doc core.Document, err error) {
	timeSpent := time.GetTimeSpent(func() {
		doc, err = m.inner.Generate()
	})
	m.generateTime = timeSpent
	if err != nil {
		return nil, err
	}

	bytes := doc.GetBytes()
	report := m.buildMetrics(len(bytes)).Normalize()

	return core.NewPDF(bytes, report), nil
}

// RegisterHeader decorates the RegisterHeader method of bpdf instance.
func (m *MetricsDecorator) RegisterHeader(rows ...core.Row) (err error) {
	timeSpent := time.GetTimeSpent(func() {
		err = m.inner.RegisterHeader(rows...)
	})
	m.headerTime = timeSpent
	return
}

// RegisterFooter decorates the RegisterFooter method of bpdf instance.
func (m *MetricsDecorator) RegisterFooter(rows ...core.Row) (err error) {
	timeSpent := time.GetTimeSpent(func() {
		err = m.inner.RegisterFooter(rows...)
	})
	m.footerTime = timeSpent
	return
}

// AddRow decorates the AddRow method of bpdf instance.
func (m *MetricsDecorator) AddRow(rowHeight float64, cols ...core.Col) core.Row {
	var r core.Row
	timeSpent := time.GetTimeSpent(func() {
		r = m.inner.AddRow(rowHeight, cols...)
	})

	m.addRowTime = append(m.addRowTime, timeSpent)
	return r
}

// AddRows decorates the AddRows method of bpdf instance.
func (m *MetricsDecorator) AddRows(rows ...core.Row) {
	timeSpent := time.GetTimeSpent(func() {
		m.inner.AddRows(rows...)
	})

	m.addRowsTime = append(m.addRowsTime, timeSpent)
}

// AddAutoRow decorates the AddRow method of bpdf instance.
func (m *MetricsDecorator) AddAutoRow(cols ...core.Col) core.Row {
	var r core.Row
	timeSpent := time.GetTimeSpent(func() {
		r = m.inner.AddAutoRow(cols...)
	})

	m.addAutoRowTime = append(m.addAutoRowTime, timeSpent)
	return r
}

// AddPages decorates the AddPages method of bpdf instance.
func (m *MetricsDecorator) AddPages(pages ...core.Page) {
	timeSpent := time.GetTimeSpent(func() {
		m.inner.AddPages(pages...)
	})

	m.addPageTime = append(m.addPageTime, timeSpent)
}

// FitlnCurrentPage decoratess the FitlnCurrentPage method of bpdf instance.
func (m *MetricsDecorator) FitlnCurrentPage(heightNewLine float64) bool {
	return m.inner.FitlnCurrentPage(heightNewLine)
}

func (m *MetricsDecorator) getAVG(times []*metrics.Time) *metrics.Time {
	var sum float64
	for _, time := range times {
		sum += time.Value
	}

	return &metrics.Time{
		Value: sum / float64(len(times)),
		Scale: times[0].Scale,
	}
}

func (m *MetricsDecorator) buildMetrics(bytesSize int) *metrics.Report {
	var timeMetrics []metrics.TimeMetric
	if m.structureTime != nil {
		timeMetrics = append(timeMetrics, metrics.TimeMetric{
			Key:   "get_tree_structure",
			Times: []*metrics.Time{m.structureTime},
			Avg:   m.structureTime,
		})
	}

	if m.generateTime != nil {
		timeMetrics = append(timeMetrics, metrics.TimeMetric{
			Key:   "generate",
			Times: []*metrics.Time{m.generateTime},
			Avg:   m.generateTime,
		})
	}

	if m.headerTime != nil {
		timeMetrics = append(timeMetrics, metrics.TimeMetric{
			Key:   "header",
			Times: []*metrics.Time{m.headerTime},
			Avg:   m.headerTime,
		})
	}

	if m.footerTime != nil {
		timeMetrics = append(timeMetrics, metrics.TimeMetric{
			Key:   "footer",
			Times: []*metrics.Time{m.footerTime},
			Avg:   m.footerTime,
		})
	}

	if len(m.addPageTime) > 0 {
		timeMetrics = append(timeMetrics, metrics.TimeMetric{
			Key:   "add_page",
			Times: m.addPageTime,
			Avg:   m.getAVG(m.addPageTime),
		})
	}

	if len(m.addRowTime) > 0 {
		timeMetrics = append(timeMetrics, metrics.TimeMetric{
			Key:   "add_row",
			Times: m.addRowTime,
			Avg:   m.getAVG(m.addRowTime),
		})
	}

	if len(m.addRowsTime) > 0 {
		timeMetrics = append(timeMetrics, metrics.TimeMetric{
			Key:   "add_rows",
			Times: m.addRowsTime,
			Avg:   m.getAVG(m.addRowsTime),
		})
	}

	return &metrics.Report{
		TimeMetrics: timeMetrics,
		SizeMetric: metrics.SizeMetric{
			Key: "file_size",
			Size: metrics.Size{
				Value: float64(bytesSize),
				Scale: metrics.Byte,
			},
		},
	}
}
