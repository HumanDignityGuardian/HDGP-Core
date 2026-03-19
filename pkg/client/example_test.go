package client_test

import (
	"context"
	"fmt"

	"github.com/HumanDignityGuardian/HDGP-Core/internal/engine"
	"github.com/HumanDignityGuardian/HDGP-Core/pkg/client"
)

func ExampleClient_Evaluate() {
	c := client.New("http://localhost:8080")
	req := engine.EvaluateRequest{
		Meta: engine.Meta{},
		Subject: engine.Subject{
			Type: "output_text",
		},
		Candidate: engine.Candidate{
			Text: "Hello, this is a safe message.",
		},
	}
	resp, err := c.Evaluate(context.Background(), req, false)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	fmt.Printf("verdict: %s\n", resp.Verdict)
}
