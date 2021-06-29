package mockapi

import (
	"fmt"
	"strconv"
	"time"
)

type Vdc struct {
	ID         string     `json:"id"`
	Specs      string     `json:"specs"`
	CreateTime time.Time  `json:"create_time"`
	UpdateTime time.Time  `json:"update_time"`
	DeleteTime *time.Time `json:"delete_time"`
}

type VdcSpecs struct {
	ID         string     `json:"id"`
	Spec       string     `json:"spec"`
	CreateTime time.Time  `json:"create_time"`
	UpdateTime time.Time  `json:"update_time"`
	DeleteTime *time.Time `json:"delete_time"`
}

var (
	vdcMap      = map[string]*Vdc{}
	vdcSpecsMap = map[string]*VdcSpecs{}
)

func CreateVdc(specs string) (*Vdc, error) {
	vdc := &Vdc{
		ID:         strconv.FormatInt(time.Now().Unix(), 10),
		Specs:      specs,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
		DeleteTime: nil,
	}
	if _, exist := vdcMap[vdc.ID]; exist {
		return nil, fmt.Errorf("vdc id %s is already exist", vdc.ID)
	}
	vdcMap[vdc.ID] = vdc
	return vdc, nil
}

func ReadVdc(id string) (*Vdc, error) {
	if vdc, exist := vdcMap[id]; exist {
		if vdc.DeleteTime == nil {
			return vdc, nil
		}
	}
	return nil, fmt.Errorf("vdc id %s not exist", id)
}

func UpdateVdc(id, specs string) (*Vdc, error) {
	if vdc, exist := vdcMap[id]; exist {
		vdc.Specs = specs
		vdc.UpdateTime = time.Now()
		return vdc, nil
	}
	return nil, fmt.Errorf("vdc id %s is not exist", id)
}

func DeleteVdc(id string) error {
	if vdc, exist := vdcMap[id]; exist {
		now := time.Now()
		vdc.DeleteTime = &now
		return nil
	}
	return nil
}

func CreateVdcSpecs(specs string) (*VdcSpecs, error) {
	vdcSpecs := &VdcSpecs{
		ID:         strconv.FormatInt(time.Now().Unix(), 10),
		Spec:       specs,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
		DeleteTime: nil,
	}
	if _, exist := vdcSpecsMap[vdcSpecs.ID]; exist {
		return nil, fmt.Errorf("vdc spec id %s is already exist", vdcSpecs.ID)
	}
	vdcSpecsMap[vdcSpecs.ID] = vdcSpecs
	return vdcSpecs, nil
}

func ReadVdcSpecs(id string) (*VdcSpecs, error) {
	if vdcSpec, exist := vdcSpecsMap[id]; exist {
		if vdcSpec.DeleteTime == nil {
			return vdcSpec, nil
		}
	}
	return nil, fmt.Errorf("vdc spec id %s is not exist", id)
}

func UpdateVdcSpecs(id, spec string) (*VdcSpecs, error) {
	if vdcSpec, exist := vdcSpecsMap[id]; exist {
		vdcSpec.Spec = spec
		vdcSpec.UpdateTime = time.Now()
		return vdcSpec, nil
	}
	return nil, fmt.Errorf("vdc spec id %s is not exist", id)
}

func DeleteVdcSpecs(id string) error {
	if vdcSpec, exist := vdcSpecsMap[id]; exist {
		now := time.Now()
		vdcSpec.DeleteTime = &now
		return nil
	}
	return fmt.Errorf("vdc spec id %s is not exist", id)
}

func GetVdcSepcList() []VdcSpecs {
	list := make([]VdcSpecs, 0)
	for _, vdcSpec := range vdcSpecsMap {
		list = append(list, *vdcSpec)
	}
	return list
}

func GetVdcList() []Vdc {
	list := make([]Vdc, 0)
	for _, vdc := range vdcMap {
		list = append(list, *vdc)
	}

	return list
}
