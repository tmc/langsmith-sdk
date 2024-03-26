package langsmith_test

import (
	"context"
	"fmt"
	"net/http/httputil"
	"os"

	"github.com/google/uuid"
	"github.com/langchain-ai/langsmith-sdk/go/langsmith"
)

func ExampleClient_CreateRunRunsPost() {
	client, err := langsmith.NewClientWithResponses(langsmith.DefaultBaseURL, langsmith.WithAPIKey(os.Getenv("LANGCHAIN_API_KEY")))
	if err != nil {
		panic(err)
	}

	runType := langsmith.Llm
	inputs := map[string]interface{}{
		"model":  "text-davinci-002",
		"prompt": "Hello, world!",
	}
	body := langsmith.CreateRunRunsPostJSONRequestBody{
		Name:    "Example Run",
		RunType: runType,
		Inputs:  &langsmith.RunCreateSchemaExtended_Inputs{},
	}
	body.Inputs.FromRunCreateSchemaExtendedInputs0(inputs)

	resp, err := client.CreateRunRunsPost(context.Background(), body)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// print the response using httputil.DumpResponse
	rbody, _ := httputil.DumpResponse(resp, true)
	fmt.Println(string(rbody))

	fmt.Println("Run created successfully")
	// output: Run created successfully
}

func ExampleClient_UpdateRunRunsRunIdPatch() {
	client, err := langsmith.NewClientWithResponses(langsmith.DefaultBaseURL, langsmith.WithAPIKey(os.Getenv("LANGCHAIN_API_KEY")))
	if err != nil {
		panic(err)
	}

	runID, err := uuid.Parse("123e4567-e89b-12d3-a456-426655440000")
	if err != nil {
		panic(err)
	}
	outputs := map[string]interface{}{
		"text": "The response from the language model.",
	}
	body := langsmith.UpdateRunRunsRunIdPatchJSONRequestBody{
		Outputs: &langsmith.RunUpdateSchemaExtended_Outputs{},
	}
	body.Outputs.FromRunUpdateSchemaExtendedOutputs0(outputs)

	resp, err := client.UpdateRunRunsRunIdPatch(context.Background(), runID, body)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	rbody, _ := httputil.DumpResponse(resp, true)
	fmt.Println(string(rbody))
	fmt.Println("Run updated successfully")
	// output: Run updated successfully
}
