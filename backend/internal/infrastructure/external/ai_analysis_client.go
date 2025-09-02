package external

import (
	"context"
	"google.golang.org/genai"
)

type GeminiClient struct {
	Client  *genai.Client
	Context context.Context
	Model   string
}

func NewGeminiClient(ctx context.Context, model string) (*GeminiClient, error) {
	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		return nil, err
	}
	return &GeminiClient{
		Client:  client,
		Context: ctx,
		Model:   model,
	}, nil
}

func (c *GeminiClient) GenerateResponse(prompt string) (*genai.GenerateContentResponse, error) {
	result, err := c.Client.Models.GenerateContent(
		c.Context,
		c.Model,
		genai.Text(prompt),
		nil,
	)
	if err != nil {
		return nil, err
	}
	return result, nil
}
