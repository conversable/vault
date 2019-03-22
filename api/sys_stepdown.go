package api

import "context"

func (c *Sys) StepDown(ctx context.Context) error {
	r := c.c.NewRequest("PUT", "/v1/sys/step-down")

	ctx, cancelFunc := context.WithCancel(ctx)
	defer cancelFunc()
	resp, err := c.c.RawRequestWithContext(ctx, r)
	if resp != nil && resp.Body != nil {
		resp.Body.Close()
	}
	return err
}
