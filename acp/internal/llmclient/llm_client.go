package llmclient

import (
	"context"
	"fmt"

	acp "github.com/humanlayer/agentcontrolplane/acp/api/v1alpha1"
)

// LLMClient defines the interface for interacting with LLM providers
type LLMClient interface {
	// SendRequest sends a request to the LLM and returns the response
	SendRequest(ctx context.Context, messages []acp.Message, tools []Tool) (*acp.Message, error)
}

// LLMRequestError represents an error that occurred during an LLM request
// and includes HTTP status code information
type LLMRequestError struct {
	StatusCode int
	Message    string
	Err        error
}

func (e *LLMRequestError) Error() string {
	return fmt.Sprintf("LLM request failed with status %d: %s", e.StatusCode, e.Message)
}

func (e *LLMRequestError) Unwrap() error {
	return e.Err
}

// Tool represents a function that can be called by the LLM
type Tool struct {
	Type     string       `json:"type"`
	Function ToolFunction `json:"function"`
	// ACPToolType represents the ACP-specific type of tool (MCP, HumanContact)
	// This field is not sent to the LLM API but is used internally for tool identification
	ACPToolType acp.ToolType `json:"-"`
}

// ToolFunction contains the function details
type ToolFunction struct {
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Parameters  ToolFunctionParameters `json:"parameters"`
}

// ToolFunctionParameters defines the schema for the function parameters
// It's a map to accommodate any valid JSON Schema structure
type ToolFunctionParameters map[string]interface{}

// FromContactChannel creates a Tool from a ContactChannel resource
func ToolFromContactChannel(channel acp.ContactChannel) *Tool {
	// Create base parameters structure for human contact tools
	params := ToolFunctionParameters{
		"type": "object",
		"properties": map[string]interface{}{
			"message": map[string]interface{}{
				"type": "string",
			},
		},
		"required": []string{"message"},
	}

	var description string
	var name string

	// Customize based on channel type
	switch channel.Spec.Type {
	case acp.ContactChannelTypeEmail:

		name = fmt.Sprintf("%s__human_contact_email", channel.Name)
		description = channel.Spec.Email.ContextAboutUser
		if description == "" {
			description = "Contact a human via email"
		}

	case acp.ContactChannelTypeSlack:
		name = fmt.Sprintf("%s__human_contact_slack", channel.Name)
		description = channel.Spec.Slack.ContextAboutChannelOrUser
		if description == "" {
			description = "Contact a human via Slack"
		}

	default:
		name = fmt.Sprintf("%s__human_contact", channel.Name)
		description = fmt.Sprintf("Contact a human via %s channel", channel.Spec.Type)
	}
	// Create the Tool
	return &Tool{
		Type: "function",
		Function: ToolFunction{
			Name:        name,
			Description: description,
			Parameters:  params,
		},
		ACPToolType: acp.ToolTypeHumanContact, // Set as HumanContact type
	}
}
