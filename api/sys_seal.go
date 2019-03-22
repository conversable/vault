package api

import "context"

func (c *Sys) SealStatus(ctx context.Context) (*SealStatusResponse, error) {
	r := c.c.NewRequest("GET", "/v1/sys/seal-status")
	return sealStatusRequest(ctx, c, r)
}

func (c *Sys) Seal(ctx context.Context) error {
	r := c.c.NewRequest("PUT", "/v1/sys/seal")

	ctx, cancelFunc := context.WithCancel(ctx)
	defer cancelFunc()
	resp, err := c.c.RawRequestWithContext(ctx, r)
	if err == nil {
		defer resp.Body.Close()
	}
	return err
}

func (c *Sys) ResetUnsealProcess(ctx context.Context) (*SealStatusResponse, error) {
	body := map[string]interface{}{"reset": true}

	r := c.c.NewRequest("PUT", "/v1/sys/unseal")
	if err := r.SetJSONBody(body); err != nil {
		return nil, err
	}

	return sealStatusRequest(ctx, c, r)
}

func (c *Sys) Unseal(ctx context.Context, shard string) (*SealStatusResponse, error) {
	body := map[string]interface{}{"key": shard}

	r := c.c.NewRequest("PUT", "/v1/sys/unseal")
	if err := r.SetJSONBody(body); err != nil {
		return nil, err
	}

	return sealStatusRequest(ctx, c, r)
}

func (c *Sys) UnsealWithOptions(ctx context.Context, opts *UnsealOpts) (*SealStatusResponse, error) {
	r := c.c.NewRequest("PUT", "/v1/sys/unseal")
	if err := r.SetJSONBody(opts); err != nil {
		return nil, err
	}

	return sealStatusRequest(ctx, c, r)
}

func sealStatusRequest(ctx context.Context, c *Sys, r *Request) (*SealStatusResponse, error) {
	ctx, cancelFunc := context.WithCancel(ctx)
	defer cancelFunc()
	resp, err := c.c.RawRequestWithContext(ctx, r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result SealStatusResponse
	err = resp.DecodeJSON(&result)
	return &result, err
}

type SealStatusResponse struct {
	Type         string `json:"type"`
	Initialized  bool   `json:"initialized"`
	Sealed       bool   `json:"sealed"`
	T            int    `json:"t"`
	N            int    `json:"n"`
	Progress     int    `json:"progress"`
	Nonce        string `json:"nonce"`
	Version      string `json:"version"`
	Migration    bool   `json:"migration"`
	ClusterName  string `json:"cluster_name,omitempty"`
	ClusterID    string `json:"cluster_id,omitempty"`
	RecoverySeal bool   `json:"recovery_seal"`
}

type UnsealOpts struct {
	Key     string `json:"key"`
	Reset   bool   `json:"reset"`
	Migrate bool   `json:"migrate"`
}
