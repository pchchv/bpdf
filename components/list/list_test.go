package list_test

import (
	"fmt"
	"testing"

	"github.com/pchchv/bpdf/components/list"
	"github.com/pchchv/bpdf/components/page"
	"github.com/pchchv/bpdf/components/row"
	"github.com/pchchv/bpdf/components/text"
	"github.com/pchchv/bpdf/core"
	"github.com/pchchv/bpdf/internal/fixture"
	"github.com/pchchv/bpdf/test"
	"github.com/stretchr/testify/assert"
)

type anyType struct {
	Key   string
	Value string
}

func (a anyType) GetHeader() core.Row {
	r := row.New(10).Add(
		text.NewCol(6, "Key"),
		text.NewCol(6, "Value"),
	)

	return r
}

func (a anyType) GetContent(i int) core.Row {
	r := row.New(10).Add(
		text.NewCol(6, a.Key),
		text.NewCol(6, a.Value),
	)

	if i%2 == 0 {
		cell := fixture.CellProp()
		r.WithStyle(&cell)
	}

	return r
}

func buildList(qtd int) (arr []anyType) {
	for i := 0; i < qtd; i++ {
		arr = append(arr, anyType{
			Key:   fmt.Sprintf("key(%d)", i),
			Value: fmt.Sprintf("value(%d)", i),
		})
	}

	return
}

func buildPointerList(qtd int) (arr []*anyType) {
	for i := 0; i < qtd; i++ {
		arr = append(arr, &anyType{
			Key:   fmt.Sprintf("key(%d)", i),
			Value: fmt.Sprintf("value(%d)", i),
		})
	}

	return
}

func TestBuild(t *testing.T) {
	t.Run("when arr is empty, should return error", func(t *testing.T) {
		r, err := list.Build[anyType](nil)

		assert.NotNil(t, err)
		assert.Nil(t, r)
	})

	t.Run("when arr is not empty, should return rows", func(t *testing.T) {
		arr := buildList(10)

		r, err := list.Build(arr)
		p := page.New().Add(r...)

		assert.Nil(t, err)
		test.New(t).Assert(p.GetStructure()).Equals("components/list/build.json")
	})
}

func TestBuildFromPointer(t *testing.T) {
	t.Run("when arr is empty, should return error", func(t *testing.T) {
		arr := buildPointerList(0)

		r, err := list.BuildFromPointer(arr)

		assert.NotNil(t, err)
		assert.Nil(t, r)
	})

	t.Run("when arr is not empty, should return rows", func(t *testing.T) {
		arr := buildPointerList(10)

		r, _ := list.BuildFromPointer(arr)
		p := page.New().Add(r...)

		test.New(t).Assert(p.GetStructure()).Equals("components/list/build_from_pointer.json")
	})

	t.Run("when arr is has a nil element, should return error", func(t *testing.T) {
		arr := buildPointerList(10)
		arr[5] = nil

		r, err := list.BuildFromPointer(arr)

		assert.NotNil(t, err)
		assert.Nil(t, r)
	})
}
