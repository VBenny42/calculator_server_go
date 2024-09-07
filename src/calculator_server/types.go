package main

import (
	"encoding/json"
	"errors"
	"strconv"
)

type genericNumberInput interface {
	validate() error
}

type numberPair struct {
	Number1 *int `json:"number1"`
	Number2 *int `json:"number2"`
}

func (pair *numberPair) UnmarshalJSON(data []byte) error {
	type Alias numberPair
	aux := &struct {
		Number1 interface{} `json:"number1"`
		Number2 interface{} `json:"number2"`
		*Alias
	}{
		Alias: (*Alias)(pair),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	if aux.Number1 != nil {
		switch v := aux.Number1.(type) {
		case string:
			num, err := strconv.Atoi(aux.Number1.(string))
			if err != nil {
				return errors.New("number1 field must be a valid integer")
			}
			pair.Number1 = &num
		case float64:
			num := int(v)
			pair.Number1 = &num
		default:
			return errors.New("number1 field must be a valid number")
		}
	}

	if aux.Number2 != nil {
		switch v := aux.Number2.(type) {
		case string:
			num, err := strconv.Atoi(aux.Number2.(string))
			if err != nil {
				return errors.New("number2 field must be a valid integer")
			}
			pair.Number2 = &num
		case float64:
			num := int(v)
			pair.Number2 = &num
		case int:
			num := v
			pair.Number2 = &num
		default:
			return errors.New("number2 field must be a valid number")
		}
	}

	return nil
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

func (pair *dividePair) UnmarshalJSON(data []byte) error {
	type Alias dividePair
	aux := &struct {
		Divisor  interface{} `json:"divisor"`
		Dividend interface{} `json:"Dividend"`
		*Alias
	}{
		Alias: (*Alias)(pair),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	if aux.Divisor != nil {
		switch v := aux.Divisor.(type) {
		case string:
			num, err := strconv.Atoi(aux.Divisor.(string))
			if err != nil {
				return errors.New("Divisor field must be a valid integer")
			}
			pair.Divisor = &num
		case float64:
			num := int(v)
			pair.Divisor = &num
		default:
			return errors.New("Divisor field must be a valid number")
		}
	}

	if aux.Dividend != nil {
		switch v := aux.Dividend.(type) {
		case string:
			num, err := strconv.Atoi(aux.Dividend.(string))
			if err != nil {
				return errors.New("Dividend field must be a valid integer")
			}
			pair.Dividend = &num
		case float64:
			num := int(v)
			pair.Dividend = &num
		case int:
			num := v
			pair.Dividend = &num
		default:
			return errors.New("Dividend field must be a valid number")
		}
	}

	return nil
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

func (s *sumArray) UnmarshalJSON(data []byte) error {
	type Alias sumArray
	aux := &struct {
		Items interface{} `json:"items"`
		*Alias
	}{
		Alias: (*Alias)(s),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	if aux.Items != nil {
		switch v := aux.Items.(type) {
		case []interface{}:
			for _, item := range v {
				switch i := item.(type) {
				case string:
					num, err := strconv.Atoi(i)
					if err != nil {
						return errors.New("items field must be a list of valid integers")
					}
					s.Items = append(s.Items, num)
				case float64:
					num := int(i)
					s.Items = append(s.Items, num)
				default:
					return errors.New("items field must be a list of valid numbers")
				}
			}
		default:
			return errors.New("items field must be a list of numbers")
		}
	}

	return nil
}

func (s sumArray) validate() error {
	if len(s.Items) == 0 {
		return nil
	}
	if s.Items == nil {
		return errors.New("items field is required but missing")
	}
	return nil
}
