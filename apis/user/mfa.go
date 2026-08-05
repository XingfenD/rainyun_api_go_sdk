package user

import (
	"github.com/XingfenD/rainyun_api_go_sdk/apis/common"
	"github.com/XingfenD/rainyun_api_go_sdk/constant"
)

// Request2FARequest 请求二次验证
type Request2FARequest struct {
	Type string `json:"type"`
}

// 请求二次验证
//
// verificationMethod: 验证方式(sms/totp/email)
func (s *UserService) Request2FA(verificationMethod string) (*common.BasicOperationResponse, error) {
	path := "/user/mfa/request"

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, Request2FARequest{Type: verificationMethod}, &resp)

	return &resp, err
}

// Verify2FAResultRequest 验证二次验证结果
type Verify2FAResultRequest struct {
	AuthCode int `json:"auth_code"`
}

// 验证二次验证结果
//
// authCode: 验证码
func (s *UserService) Verify2FAResult(authCode int) (*common.BasicOperationResponse, error) {
	path := "/user/mfa/verify"

	var resp common.BasicOperationResponse
	err := s.client.Do(constant.HTTPMethod_POST, path, nil, Verify2FAResultRequest{AuthCode: authCode}, &resp)

	return &resp, err
}
