package steps

import (
	"reflect"
	"testing"

	"github.com/bitrise-io/stepman/models"
)

func TestParseYamlFile(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name          string
		args          args
		expectedTitle string
		wantErr       bool
	}{
		{
			"Valid yaml file",
			args{"testdata/activate_ssh.yaml"},
			"Activate SSH key (RSA private key)",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseYamlFile(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseYamlFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if *got.Title != tt.expectedTitle {
				t.Errorf("ParseYamlFile() = %v, want %v", *got.Title, tt.expectedTitle)
			}
		})
	}
}

func TestParseYaml(t *testing.T) {
	type args struct {
		bytes []byte
	}
	tests := []struct {
		name    string
		args    args
		want    *models.StepModel
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseYaml(tt.args.bytes)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseYaml() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseYaml() = %v, want %v", got, tt.want)
			}
		})
	}
}
