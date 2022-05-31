package clients

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	v1 "github.com/nndergunov/deliveryApp/app/pkg/api/v1"
)

var ErrStatusNotUp = errors.New("service status is not up")

type MultiServiceClient struct {
	serviceAddresses map[string]string
}

func NewMultiServiceClient(serviceAddresses map[string]string) *MultiServiceClient {
	return &MultiServiceClient{serviceAddresses: serviceAddresses}
}

func (m MultiServiceClient) CheckStatuses() error {
	for service, addr := range m.serviceAddresses {
		res, err := http.Get(addr + "/v1/status")
		if err != nil {
			return fmt.Errorf("sending request to %s service: %w", service, err)
		}

		resData, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return fmt.Errorf("receiving response from %s service: %w", service, err)
		}

		err = res.Body.Close()
		if err != nil {
			return fmt.Errorf("closing response body: %w", err)
		}

		resp := new(v1.Status)

		err = v1.Decode(resData, resp)
		if err != nil {
			return fmt.Errorf("sending request to %s service: %w", service, err)
		}

		if resp.IsUp != "up" {
			return fmt.Errorf("%s %w", service, ErrStatusNotUp)
		}
	}

	return nil
}
