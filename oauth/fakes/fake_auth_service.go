// This file was generated by counterfeiter
package fakes

import (
	"net/http"
	"sync"

	"github.com/cfmobile/oauth2-route-service/oauth"
)

type FakeAuthService struct {
	IsUaaRedirectUrlStub        func(*http.Request) bool
	isUaaRedirectUrlMutex       sync.RWMutex
	isUaaRedirectUrlArgsForCall []struct {
		arg1 *http.Request
	}
	isUaaRedirectUrlReturns struct {
		result1 bool
	}
	AddSessionCookieStub        func(uaaRedirectRequest *http.Request, res *http.Response) error
	addSessionCookieMutex       sync.RWMutex
	addSessionCookieArgsForCall []struct {
		uaaRedirectRequest *http.Request
		res                *http.Response
	}
	addSessionCookieReturns struct {
		result1 error
	}
	HasValidAuthHeadersStub        func(*http.Request) bool
	hasValidAuthHeadersMutex       sync.RWMutex
	hasValidAuthHeadersArgsForCall []struct {
		arg1 *http.Request
	}
	hasValidAuthHeadersReturns struct {
		result1 bool
	}
	CreateLoginRequiredResponseStub        func() (*http.Response, error)
	createLoginRequiredResponseMutex       sync.RWMutex
	createLoginRequiredResponseArgsForCall []struct{}
	createLoginRequiredResponseReturns     struct {
		result1 *http.Response
		result2 error
	}
}

func (fake *FakeAuthService) IsUaaRedirectUrl(arg1 *http.Request) bool {
	fake.isUaaRedirectUrlMutex.Lock()
	fake.isUaaRedirectUrlArgsForCall = append(fake.isUaaRedirectUrlArgsForCall, struct {
		arg1 *http.Request
	}{arg1})
	fake.isUaaRedirectUrlMutex.Unlock()
	if fake.IsUaaRedirectUrlStub != nil {
		return fake.IsUaaRedirectUrlStub(arg1)
	} else {
		return fake.isUaaRedirectUrlReturns.result1
	}
}

func (fake *FakeAuthService) IsUaaRedirectUrlCallCount() int {
	fake.isUaaRedirectUrlMutex.RLock()
	defer fake.isUaaRedirectUrlMutex.RUnlock()
	return len(fake.isUaaRedirectUrlArgsForCall)
}

func (fake *FakeAuthService) IsUaaRedirectUrlArgsForCall(i int) *http.Request {
	fake.isUaaRedirectUrlMutex.RLock()
	defer fake.isUaaRedirectUrlMutex.RUnlock()
	return fake.isUaaRedirectUrlArgsForCall[i].arg1
}

func (fake *FakeAuthService) IsUaaRedirectUrlReturns(result1 bool) {
	fake.IsUaaRedirectUrlStub = nil
	fake.isUaaRedirectUrlReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeAuthService) AddSessionCookie(uaaRedirectRequest *http.Request, res *http.Response) error {
	fake.addSessionCookieMutex.Lock()
	fake.addSessionCookieArgsForCall = append(fake.addSessionCookieArgsForCall, struct {
		uaaRedirectRequest *http.Request
		res                *http.Response
	}{uaaRedirectRequest, res})
	fake.addSessionCookieMutex.Unlock()
	if fake.AddSessionCookieStub != nil {
		return fake.AddSessionCookieStub(uaaRedirectRequest, res)
	} else {
		return fake.addSessionCookieReturns.result1
	}
}

func (fake *FakeAuthService) AddSessionCookieCallCount() int {
	fake.addSessionCookieMutex.RLock()
	defer fake.addSessionCookieMutex.RUnlock()
	return len(fake.addSessionCookieArgsForCall)
}

func (fake *FakeAuthService) AddSessionCookieArgsForCall(i int) (*http.Request, *http.Response) {
	fake.addSessionCookieMutex.RLock()
	defer fake.addSessionCookieMutex.RUnlock()
	return fake.addSessionCookieArgsForCall[i].uaaRedirectRequest, fake.addSessionCookieArgsForCall[i].res
}

func (fake *FakeAuthService) AddSessionCookieReturns(result1 error) {
	fake.AddSessionCookieStub = nil
	fake.addSessionCookieReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAuthService) HasValidAuthHeaders(arg1 *http.Request) bool {
	fake.hasValidAuthHeadersMutex.Lock()
	fake.hasValidAuthHeadersArgsForCall = append(fake.hasValidAuthHeadersArgsForCall, struct {
		arg1 *http.Request
	}{arg1})
	fake.hasValidAuthHeadersMutex.Unlock()
	if fake.HasValidAuthHeadersStub != nil {
		return fake.HasValidAuthHeadersStub(arg1)
	} else {
		return fake.hasValidAuthHeadersReturns.result1
	}
}

func (fake *FakeAuthService) HasValidAuthHeadersCallCount() int {
	fake.hasValidAuthHeadersMutex.RLock()
	defer fake.hasValidAuthHeadersMutex.RUnlock()
	return len(fake.hasValidAuthHeadersArgsForCall)
}

func (fake *FakeAuthService) HasValidAuthHeadersArgsForCall(i int) *http.Request {
	fake.hasValidAuthHeadersMutex.RLock()
	defer fake.hasValidAuthHeadersMutex.RUnlock()
	return fake.hasValidAuthHeadersArgsForCall[i].arg1
}

func (fake *FakeAuthService) HasValidAuthHeadersReturns(result1 bool) {
	fake.HasValidAuthHeadersStub = nil
	fake.hasValidAuthHeadersReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeAuthService) CreateLoginRequiredResponse() (*http.Response, error) {
	fake.createLoginRequiredResponseMutex.Lock()
	fake.createLoginRequiredResponseArgsForCall = append(fake.createLoginRequiredResponseArgsForCall, struct{}{})
	fake.createLoginRequiredResponseMutex.Unlock()
	if fake.CreateLoginRequiredResponseStub != nil {
		return fake.CreateLoginRequiredResponseStub()
	} else {
		return fake.createLoginRequiredResponseReturns.result1, fake.createLoginRequiredResponseReturns.result2
	}
}

func (fake *FakeAuthService) CreateLoginRequiredResponseCallCount() int {
	fake.createLoginRequiredResponseMutex.RLock()
	defer fake.createLoginRequiredResponseMutex.RUnlock()
	return len(fake.createLoginRequiredResponseArgsForCall)
}

func (fake *FakeAuthService) CreateLoginRequiredResponseReturns(result1 *http.Response, result2 error) {
	fake.CreateLoginRequiredResponseStub = nil
	fake.createLoginRequiredResponseReturns = struct {
		result1 *http.Response
		result2 error
	}{result1, result2}
}

var _ oauth.AuthService = new(FakeAuthService)
