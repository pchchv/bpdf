package test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/pchchv/bpdf/internal/fixture"
	"github.com/stretchr/testify/assert"
)

const file = "bpdf_test.json"

func TestNew(t *testing.T) {
	t.Run("when called first, should setup singleton and set t", func(t *testing.T) {
		sut := New(t)

		assert.Equal(t, t, sut.t)
	})

	t.Run("when called not first, should use singleton and set t", func(t *testing.T) {
		_ = New(t)

		sut := New(t)

		assert.Equal(t, t, sut.t)
	})
}

func TestBPDFTest_Assert(t *testing.T) {
	t.Run("when call assert, should set node", func(t *testing.T) {
		n := fixture.Node("bpdf")
		sut := New(t)

		sut.Assert(n)

		assert.Equal(t, n, sut.node)
	})
}

func TestBPDFTest_Save(t *testing.T) {
	t.Run("when cannot save, should not create file", func(t *testing.T) {
		file := ""
		n := fixture.Node("bpdf")
		innerT := &testing.T{}
		sut := New(innerT).Assert(n)

		sut.Equals(file)

		path := configSingleton.getAbsoluteFilePath(file)
		_, err := os.ReadFile(path)
		assert.NotNil(t, err)
		assert.True(t, innerT.Failed())
	})

	t.Run("when can save, should create file", func(t *testing.T) {
		n := fixture.Node("bpdf")
		sut := New(t).Assert(n)

		sut.Equals(file)

		path := configSingleton.getAbsoluteFilePath(file)
		bytes, err := os.ReadFile(path)
		assert.Nil(t, err)
		testNode := &Node{}
		_ = json.Unmarshal(bytes, testNode)
		assert.Equal(t, "bpdf", testNode.Type)
		assert.Equal(t, "page", testNode.Nodes[0].Type)
	})
}

func TestBPDFTest_Equals(t *testing.T) {
	t.Run("when file saved is not equals to current, should fail", func(t *testing.T) {
		n := fixture.Node("not_bpdf")
		innerT := &testing.T{}
		sut := New(innerT).Assert(n)

		sut.Equals(file)

		assert.True(t, innerT.Failed())
	})

	t.Run("when file saved is equals to current, should be success", func(t *testing.T) {
		n := fixture.Node("bpdf")
		innerT := &testing.T{}
		sut := New(innerT).Assert(n)

		sut.Equals(file)

		assert.False(t, innerT.Failed())
	})
}
