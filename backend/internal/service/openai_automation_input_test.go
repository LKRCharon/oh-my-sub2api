package service

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSanitizeOpenAIResponsesPreservesScheduledAutomationInput(t *testing.T) {
	const heartbeat = `<heartbeat>
  <automation_id>mail-review</automation_id>
  <current_time_iso>2026-09-15T10:01:58.531Z</current_time_iso>
  <instructions>
Read the local queue. Record the source as mail:<message-hash>; preserve A && B.
  </instructions>
</heartbeat>
`
	for _, withHistory := range []bool{false, true} {
		name := "without history"
		if withHistory {
			name = "with completed tool call history"
		}
		t.Run(name, func(t *testing.T) {
			input := []any{}
			if withHistory {
				input = append(input,
					map[string]any{"type": "function_call", "name": "exec_command", "call_id": "call_history", "arguments": "{}"},
					map[string]any{"type": "function_call_output", "call_id": "call_history", "output": "previous check complete"},
					map[string]any{"type": "message", "role": "assistant", "content": "Previous check finished."},
				)
			}
			wake := map[string]any{
				"type":      "function_call_output",
				"id":        "fco_scheduled_wake",
				"namespace": "codex_app",
				"name":      "automation_update",
				"output":    heartbeat,
			}
			input = append(input, wake)
			// This is the OAuth forwarding pipeline: ID normalization runs before
			// orphan filtering. A client message ID must not become a call ID.
			input = filterCodexInputWithOptions(input, codexInputFilterOptions{
				PreserveReferences: true,
				PreserveCallIDs:    true,
			})
			request := map[string]any{"input": input}

			require.False(t, sanitizeOpenAIResponsesOrphanToolOutputs(request, input, false))
			require.Equal(t, input, request["input"])
			got, ok := request["input"].([]any)
			require.True(t, ok)
			require.NotEmpty(t, got)
			lastItem, ok := got[len(got)-1].(map[string]any)
			require.True(t, ok)
			require.Equal(t, heartbeat, lastItem["output"])
			require.NotContains(t, lastItem, "call_id")
			require.Equal(t, "automation_update", lastItem["name"])
			require.Equal(t, "codex_app", lastItem["namespace"])
		})
	}
}

func TestNamedStandaloneInputDoesNotRequireToolContinuation(t *testing.T) {
	wake := map[string]any{
		"type": "function_call_output", "id": "fco_scheduled_wake",
		"name": "automation_update", "namespace": "codex_app",
		"output": "Preserve qq-mail:<message-hash> and A && B.",
	}
	request := map[string]any{"input": []any{wake}}
	body, err := json.Marshal(request)
	require.NoError(t, err)
	require.Equal(t, FunctionCallOutputValidation{}, ValidateFunctionCallOutputContext(request))
	require.Equal(t, FunctionCallOutputValidation{}, ValidateFunctionCallOutputContextBytes(body))
	require.Equal(t, ToolContinuationSignals{}, AnalyzeToolContinuationSignals(request))
	require.Equal(t, ToolCallOutputContextCoverage{}, AnalyzeToolCallOutputContextCoverageBytes(body))
	require.False(t, NeedsToolContinuation(request))

	// A standalone wake must not exempt a genuine, unpaired tool result.
	request["input"] = []any{wake, map[string]any{"type": "function_call_output", "output": "unpaired result"}}
	body, err = json.Marshal(request)
	require.NoError(t, err)
	require.True(t, ValidateFunctionCallOutputContext(request).HasFunctionCallOutputMissingCallID)
	require.True(t, ValidateFunctionCallOutputContextBytes(body).HasFunctionCallOutputMissingCallID)
	require.True(t, AnalyzeToolContinuationSignals(request).HasFunctionCallOutputMissingCallID)

	// Normal paired history still carries its continuation context.
	request["input"] = []any{
		map[string]any{"type": "custom_tool_call", "name": "exec", "call_id": "call_history"},
		map[string]any{"type": "custom_tool_call_output", "call_id": "call_history", "output": "done"},
		wake,
	}
	body, err = json.Marshal(request)
	require.NoError(t, err)
	require.Equal(t, ToolCallOutputContextCoverage{HasFunctionCallOutput: true, ContextCoversAllCallIDs: true}, AnalyzeToolCallOutputContextCoverageBytes(body))
}
