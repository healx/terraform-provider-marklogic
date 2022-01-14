package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	dac "github.com/judgegregg/go-http-digest-auth-client"
)

func (a *apiClient) GetUser(ctx context.Context, username string) (map[string]interface{}, diag.Diagnostics) {

	req := dac.NewRequest(a.username, a.password, "GET", fmt.Sprintf("%s/manage/v2/users/%s/properties", a.baseUrl, username), "")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", a.userAgent)

	tflog.Trace(ctx, fmt.Sprintf("Calling api %s", fmt.Sprintf("%s/manage/v2/users/%s/properties", a.baseUrl, username)))
	tflog.Trace(ctx, fmt.Sprintf("HTTP Headers %s", req.Header))

	resp, err := req.Execute()
	if err != nil {
		return nil, diag.FromErr(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	var result map[string]interface{}

	tflog.Trace(ctx, string(body))

	err = json.Unmarshal([]byte(body), &result)

	if err != nil {
		return nil, diag.FromErr(err)
	}

	return result, nil
}

func (a *apiClient) CreateUser(ctx context.Context, user map[string]interface{}) diag.Diagnostics {
	payload, err := json.Marshal(user)

	if err != nil {
		return diag.FromErr(err)
	}

	req := dac.NewRequest(a.username, a.password, "POST", fmt.Sprintf("%s/manage/v2/users", a.baseUrl), string(payload))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", a.userAgent)

	tflog.Trace(ctx, fmt.Sprintf("Calling api %s", fmt.Sprintf("%s/manage/v2/users", a.baseUrl)))
	tflog.Trace(ctx, fmt.Sprintf("HTTP Headers %s", req.Header))

	_, err = req.Execute()
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func (a *apiClient) DeleteUser(ctx context.Context, username string) diag.Diagnostics {
	req := dac.NewRequest(a.username, a.password, "DELETE", fmt.Sprintf("%s/manage/v2/users/%s", a.baseUrl, username), "")
	req.Header.Set("User-Agent", a.userAgent)

	tflog.Trace(ctx, fmt.Sprintf("Calling api %s", fmt.Sprintf("%s/manage/v2/users/%s", a.baseUrl, username)))
	tflog.Trace(ctx, fmt.Sprintf("HTTP Headers %s", req.Header))

	_, err := req.Execute()
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func (a *apiClient) UpdateUser(ctx context.Context, user map[string]interface{}) diag.Diagnostics {
	payload, err := json.Marshal(user)

	if err != nil {
		return diag.FromErr(err)
	}

	req := dac.NewRequest(a.username, a.password, "PUT", fmt.Sprintf("%s/manage/v2/users/%s/properties", a.baseUrl, user["user-name"]), string(payload))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", a.userAgent)

	tflog.Trace(ctx, fmt.Sprintf("Calling api %s", fmt.Sprintf("%s/manage/v2/users/%s/properties", a.baseUrl, user["user-name"])))
	tflog.Trace(ctx, fmt.Sprintf("HTTP Headers %s", req.Header))

	_, err = req.Execute()
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}
