package dbmocks

import (
	"errors"
	"testing"
)

func TestMock(t *testing.T) {
	tbl := []struct {
		query string
		resp  Response
		err   error
	}{
		{"success", Response{"Success", 200}, nil},
		{"error", Response{"", 500}, errors.New("something is wrong")},
		{"empty", Response{"", 404}, nil},
	}

	for i, item := range tbl {
		m := &Mock{}
		m.SetResponse(item.resp, item.err)
		resp, err := m.GetData(item.query)
		if resp.Text != item.resp.Text {
			t.Errorf("ITEM: %d TEXT ERROR: got %q, want %q", i+1, resp.Text, item.resp.Text)
		}

		if resp.StatusCode != item.resp.StatusCode {
			t.Errorf("ITEM: %d STATUS ERROR: got %d, want %d", i+1, resp.StatusCode, item.resp.StatusCode)
		}

		if err != nil {
			t.Errorf("ITEM: %d ERROR: got %v, want nil", i+1, err)
		}
	}
}
