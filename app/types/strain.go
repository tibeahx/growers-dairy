package types

import (
	"errors"
)

type StrainType struct {
	Auto  bool `json:"auto"`
	Photo bool `json:"photo"`
}

type FeminizedType struct {
	Feminized bool `json:"feminized"`
	Regular   bool `json:"regular"`
}

type StrainAttrs struct {
	StrainType    StrainType    `json:"strain_type"`
	FeminizedType FeminizedType `json:"feminized_type"`
}

type Strain struct {
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	StrainAttrs StrainAttrs `json:"strain_attributes"`
	Description string      `json:"description"`
}

type CreateStrainReq struct {
	Name        string       `json:"name"`
	StrainAttrs *StrainAttrs `json:"strain_attributes"`
	Description string       `json:"description"`
}

type CreateStrainResp struct {
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	StrainAttrs StrainAttrs `json:"strain_attributes"`
	Description string      `json:"description"`
}

type UpdateStrainReq struct {
	ID          int          `json:"id"`
	Name        string       `json:"name"`
	StrainAttrs *StrainAttrs `json:"strain_attributes"`
	Description string       `json:"description"`
}

type UpdateStrainResp struct {
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	StrainAttrs StrainAttrs `json:"strain_attributes"`
	Description string      `json:"description"`
}

func ValidateStrainType(t StrainType) error {
	if t.Auto == t.Photo {
		return errors.New("invalid values for strain type")
	}
	return nil
}

func ValidateFeminizedType(f FeminizedType) error {
	if f.Feminized == f.Regular {
		return errors.New("invalid values for feminized parameter")
	}
	return nil
}

func AssertStrainAttrsExist(req interface{}) error {
	switch r := req.(type) {
	case CreateStrainReq:
		if r.StrainAttrs == nil {
			return errors.New("missing strain attributes")
		}
	case UpdateStrainReq:
		if r.StrainAttrs == nil {
			return errors.New("missing strain attributes")
		}
	}
	return nil
}
