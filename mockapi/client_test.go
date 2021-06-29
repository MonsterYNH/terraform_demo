package mockapi

import "testing"

var client = &MockapiClient{
	BaseURL:  "http://localhost:8080",
	Account:  "account",
	Password: "password",
}

func TestClient(t *testing.T) {
	vdc, err := client.CreateVdc("1C2G")
	if err != nil {
		t.Fatal(err)
	}

	queryVdc, err := client.ReadVdc(vdc.ID)
	if err != nil {
		t.Fatal(err)
	}

	if queryVdc.Specs != "1C2G" {
		t.Fatal("failed create")
	}

	updateVdc, err := client.UpdateVdc(vdc.ID, "2C4G")
	if err != nil {
		t.Fatal(err)
	}

	if updateVdc.Specs != "2C4G" {
		t.Fatal("failed update")
	}
}
