package main

import "errors"

type SecurityCode struct {
	code int
}

func (s *SecurityCode) checkCode(code int) error {
	if code != s.code {
		return errors.New("wrong code")
	}
	return nil
}

func newSecurityCode(code int) *SecurityCode {
	return &SecurityCode{code: code}
}
