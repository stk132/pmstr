package pmstr

import "testing"

func TestPmstr_PutString(t *testing.T) {
	mockOutput := &mockPutParameterOutput{
		output: nil,
		err:    nil,
	}
	mock := &mockSSMClient{mockPutParameterOutput:mockOutput}
	client := NewFromSsmiface(mock)
	_, err := client.PutString("test", "testvalue").Do()
	if err != nil {
		t.Fatal(err)
	}

	input := client.PutString("test", "testvalue")
	if *input.PutParameterInput.Name != "test" {
		t.Errorf("wont: %s, got: %s", "test", *input.PutParameterInput.Name)
	}

	if *input.PutParameterInput.Value != "testvalue" {
		t.Errorf("wont: %s, got: %s", "testvalue", *input.PutParameterInput.Value)
	}

}
