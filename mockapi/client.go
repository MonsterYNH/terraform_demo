package mockapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type MockapiClient struct {
	BaseURL  string `json:"base_url"`
	Account  string `json:"account"`
	Password string `json:"password"`
}

func (client *MockapiClient) CreateVdc(specs string) (*Vdc, error) {
	response, err := http.Get(fmt.Sprintf("%s/api/v1/vdc/create?specs=%s", client.BaseURL, specs))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, errors.New("not 200")
	}

	bytes, _ := ioutil.ReadAll(response.Body)
	var vdc = &Vdc{}
	if err := json.Unmarshal(bytes, vdc); err != nil {
		return nil, err
	}

	return vdc, nil
}

func (client *MockapiClient) UpdateVdc(id, specs string) (*Vdc, error) {
	response, err := http.Get(fmt.Sprintf("%s/api/v1/vdc/update?id=%s&specs=%s", client.BaseURL, id, specs))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, errors.New("not 200")
	}

	bytes, _ := ioutil.ReadAll(response.Body)
	log.Println(string(bytes))
	var vdc = &Vdc{}
	return vdc, json.Unmarshal(bytes, vdc)
}

func (client *MockapiClient) DeleteVdc(id string) error {
	response, err := http.Get(fmt.Sprintf("%s/api/v1/vdc/delete?id=%s", client.BaseURL, id))
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return errors.New("not 200")
	}

	return err
}

func (client *MockapiClient) ReadVdc(id string) (*Vdc, error) {
	response, err := http.Get(fmt.Sprintf("%s/api/v1/vdc/instance?id=%s", client.BaseURL, id))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, errors.New("not 200")
	}

	bytes, _ := ioutil.ReadAll(response.Body)
	var vdc = &Vdc{}
	return vdc, json.Unmarshal(bytes, vdc)
}

func (client *MockapiClient) GetVdcList() ([]Vdc, error) {
	response, err := http.Get(fmt.Sprintf("%s/api/v1/vdc/list", client.BaseURL))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, errors.New("not 200")
	}

	bytes, _ := ioutil.ReadAll(response.Body)
	var list = make([]Vdc, 0)
	return list, json.Unmarshal(bytes, &list)
}

func (client *MockapiClient) CreateVdcSpecs(specs string) (*VdcSpecs, error) {
	response, err := http.Get(fmt.Sprintf("%s/api/v1/vdcspec/create?specs=%s", client.BaseURL, specs))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, errors.New("not 200")
	}

	bytes, _ := ioutil.ReadAll(response.Body)
	var vdcSpec = &VdcSpecs{}
	return vdcSpec, json.Unmarshal(bytes, vdcSpec)
}

func (client *MockapiClient) UpdateVdcSpecs(id, specs string) (*VdcSpecs, error) {
	response, err := http.Get(fmt.Sprintf("%s/api/v1/vdcspec/update?id=%s&specs=%s", client.BaseURL, id, specs))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, errors.New("not 200")
	}

	bytes, _ := ioutil.ReadAll(response.Body)
	var vdcSpec = &VdcSpecs{}
	return vdcSpec, json.Unmarshal(bytes, vdcSpec)
}

func (client *MockapiClient) DeleteVdcSpecs(id string) error {
	response, err := http.Get(fmt.Sprintf("%s/api/v1/vdcspec/delete?id=%s", client.BaseURL, id))
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return errors.New("not 200")
	}

	return err
}

func (client *MockapiClient) ReadVdcSpec(id string) (*VdcSpecs, error) {
	response, err := http.Get(fmt.Sprintf("%s/api/v1/vdcspec/instance?id=%s", client.BaseURL, id))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, errors.New("not 200")
	}

	bytes, _ := ioutil.ReadAll(response.Body)
	var vdcSpec = &VdcSpecs{}
	return vdcSpec, json.Unmarshal(bytes, vdcSpec)
}

func (client *MockapiClient) GetVdcSpecList() ([]VdcSpecs, error) {
	response, err := http.Get(fmt.Sprintf("%s/api/v1/vdcspec/list", client.BaseURL))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, errors.New("not 200")
	}

	bytes, _ := ioutil.ReadAll(response.Body)
	var list = make([]VdcSpecs, 0)
	return list, json.Unmarshal(bytes, &list)
}
