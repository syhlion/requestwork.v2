package requestwork

import (
	"context"
	"net/http"
	"testing"
	"time"
)

func TestExecute(t *testing.T) {
	req, err := http.NewRequest("GET", "http://tw.yahoo.com", nil)
	if err != nil {
		t.Error("request error: ", err)
	}
	a := New(5, nil)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	req.WithContext(ctx)
	defer cancel()
	err = a.Execute(req, func(resp *http.Response, err error) error {

		if err != nil {
			return err
		}
		defer resp.Body.Close()
		return nil

	})
	if err != nil {
		t.Error(err)
		return
	}
	err = a.Execute(req, func(resp *http.Response, err error) error {

		if err != nil {
			return err
		}
		defer resp.Body.Close()
		return nil

	})
	if err != nil {
		t.Error(err)
	}

}
