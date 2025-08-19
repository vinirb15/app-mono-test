package services

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"

	"github.com/leandro-andrade-candido/auth-service/helpers"
	"github.com/leandro-andrade-candido/auth-service/models"
)

func FindUserByID(ctx context.Context, cfg models.Config, id uuid.UUID) (models.User, error) {
	var u models.User

	url := fmt.Sprintf("%s/user/%s", cfg.ProfileServiceURL, id.String())

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return u, err
	}

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return u, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return u, fmt.Errorf("erro na API: %s", resp.Status)
	}

	if err := json.NewDecoder(resp.Body).Decode(&u); err != nil {
		return u, err
	}

	return u, nil
}

func FindUserByEmail(ctx context.Context, cfg models.Config, email string) (models.User, error) {
	var u models.User

	url := fmt.Sprintf("%s/user/%s", cfg.ProfileServiceURL, email)

	access, _, err := helpers.GenerateAccessToken(cfg, u)
	if err != nil {
		return u, fmt.Errorf("erro ao gerar access token: %w", err)
	}

	fmt.Println("Access Token:", access)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return u, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", access))

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return u, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return u, fmt.Errorf("erro na API: %s", resp.Status)
	}

	if err := json.NewDecoder(resp.Body).Decode(&u); err != nil {
		return u, err
	}

	return u, nil
}
