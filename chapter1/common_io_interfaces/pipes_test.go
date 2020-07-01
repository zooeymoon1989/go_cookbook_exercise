package common_io_interfaces

import "testing"

func TestPipeExample(t *testing.T) {
	if err := PipeExample(); err != nil {
		t.Error(err)
	}
}
