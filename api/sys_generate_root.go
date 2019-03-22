package api

import "context"

func (c *Sys) GenerateRootStatus(ctx context.Context) (*GenerateRootStatusResponse, error) {
	return c.generateRootStatusCommon(ctx, "/v1/sys/generate-root/attempt")
}

func (c *Sys) GenerateDROperationTokenStatus(ctx context.Context) (*GenerateRootStatusResponse, error) {
	return c.generateRootStatusCommon(ctx, "/v1/sys/replication/dr/secondary/generate-operation-token/attempt")
}

func (c *Sys) generateRootStatusCommon(ctx context.Context, path string) (*GenerateRootStatusResponse, error) {
	r := c.c.NewRequest("GET", path)

	ctx, cancelFunc := context.WithCancel(ctx)
	defer cancelFunc()
	resp, err := c.c.RawRequestWithContext(ctx, r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result GenerateRootStatusResponse
	err = resp.DecodeJSON(&result)
	return &result, err
}

func (c *Sys) GenerateRootInit(ctx context.Context, otp, pgpKey string) (*GenerateRootStatusResponse, error) {
	return c.generateRootInitCommon(ctx, "/v1/sys/generate-root/attempt", otp, pgpKey)
}

func (c *Sys) GenerateDROperationTokenInit(ctx context.Context, otp, pgpKey string) (*GenerateRootStatusResponse, error) {
	return c.generateRootInitCommon(ctx, "/v1/sys/replication/dr/secondary/generate-operation-token/attempt", otp, pgpKey)
}

func (c *Sys) generateRootInitCommon(ctx context.Context, path, otp, pgpKey string) (*GenerateRootStatusResponse, error) {
	body := map[string]interface{}{
		"otp":     otp,
		"pgp_key": pgpKey,
	}

	r := c.c.NewRequest("PUT", path)
	if err := r.SetJSONBody(body); err != nil {
		return nil, err
	}

	ctx, cancelFunc := context.WithCancel(ctx)
	defer cancelFunc()
	resp, err := c.c.RawRequestWithContext(ctx, r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result GenerateRootStatusResponse
	err = resp.DecodeJSON(&result)
	return &result, err
}

func (c *Sys) GenerateRootCancel(ctx context.Context) error {
	return c.generateRootCancelCommon(ctx, "/v1/sys/generate-root/attempt")
}

func (c *Sys) GenerateDROperationTokenCancel(ctx context.Context) error {
	return c.generateRootCancelCommon(ctx, "/v1/sys/replication/dr/secondary/generate-operation-token/attempt")
}

func (c *Sys) generateRootCancelCommon(ctx context.Context, path string) error {
	r := c.c.NewRequest("DELETE", path)

	ctx, cancelFunc := context.WithCancel(ctx)
	defer cancelFunc()
	resp, err := c.c.RawRequestWithContext(ctx, r)
	if err == nil {
		defer resp.Body.Close()
	}
	return err
}

func (c *Sys) GenerateRootUpdate(ctx context.Context, shard, nonce string) (*GenerateRootStatusResponse, error) {
	return c.generateRootUpdateCommon(ctx, "/v1/sys/generate-root/update", shard, nonce)
}

func (c *Sys) GenerateDROperationTokenUpdate(ctx context.Context, shard, nonce string) (*GenerateRootStatusResponse, error) {
	return c.generateRootUpdateCommon(ctx, "/v1/sys/replication/dr/secondary/generate-operation-token/update", shard, nonce)
}

func (c *Sys) generateRootUpdateCommon(ctx context.Context, path, shard, nonce string) (*GenerateRootStatusResponse, error) {
	body := map[string]interface{}{
		"key":   shard,
		"nonce": nonce,
	}

	r := c.c.NewRequest("PUT", path)
	if err := r.SetJSONBody(body); err != nil {
		return nil, err
	}

	ctx, cancelFunc := context.WithCancel(ctx)
	defer cancelFunc()
	resp, err := c.c.RawRequestWithContext(ctx, r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result GenerateRootStatusResponse
	err = resp.DecodeJSON(&result)
	return &result, err
}

type GenerateRootStatusResponse struct {
	Nonce            string `json:"nonce"`
	Started          bool   `json:"started"`
	Progress         int    `json:"progress"`
	Required         int    `json:"required"`
	Complete         bool   `json:"complete"`
	EncodedToken     string `json:"encoded_token"`
	EncodedRootToken string `json:"encoded_root_token"`
	PGPFingerprint   string `json:"pgp_fingerprint"`
	OTP              string `json:"otp"`
	OTPLength        int    `json:"otp_length"`
}
