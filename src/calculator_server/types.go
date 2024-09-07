package main

import (
	"errors"
)

type genericNumberInput interface {
	validate() error
}

type numberPair struct {
	Number1 *int `json:"number1"`
	Number2 *int `json:"number2"`
}

func (p numberPair) validate() error {
	if p.Number1 == nil {
		return errors.New("number1 field is required but missing")
	}
	if p.Number2 == nil {
		return errors.New("number2 field is required but missing")
	}
	return nil
}

type dividePair struct {
	Dividend *int `json:"dividend"`
	Divisor  *int `json:"divisor"`
}

func (p dividePair) validate() error {
	if p.Dividend == nil {
		return errors.New("dividend field is required but missing")
	}
	if p.Divisor == nil {
		return errors.New("divisor field is required but missing")
	}
	if *p.Divisor == 0 {
		return errors.New("divisor cannot be zero")
	}
	return nil
}

type sumArray struct {
	Items []int `json:"items"`
}

func (s sumArray) validate() error {
	if s.Items == nil {
		return errors.New("items field is required but missing")
	}
	return nil
}
