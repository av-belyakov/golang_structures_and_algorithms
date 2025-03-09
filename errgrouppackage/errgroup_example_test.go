package errgrouppackage

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"golang.org/x/sync/errgroup"
)

func TestErrGroup(t *testing.T) {
	var gOne, gTwo string

	g := new(errgroup.Group)
	g.Go(func() error {
		time.Sleep(3 * time.Second)

		t.Log("Goroutin ONE")

		gOne = "one"

		return nil
	})
	g.Go(func() error {
		t.Log("Goroutin TWO")

		gTwo = "two"

		return errors.New("somethere error")
	})
	err := g.Wait()

	assert.Error(t, err)
	assert.Equal(t, gOne, "one")
	assert.Equal(t, gTwo, "two")
}
