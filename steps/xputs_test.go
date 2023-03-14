package steps

import (
	"reflect"
	"testing"

	"github.com/bitrise-io/envman/models"
)

func TestCollectXPuts(t *testing.T) {
	activateSsh, _ := ParseYamlFile("testdata/activate_ssh.yaml")

	type args struct {
		models []models.EnvironmentItemModel
	}
	tests := []struct {
		name string
		args args
		want map[string]XPut
	}{
		{
			"Capture common input details",
			args{activateSsh.Inputs},
			map[string]XPut{
				"ssh_rsa_private_key": {
					Identifier: "ssh_rsa_private_key",
					Options: XPutOptions{
						Description: "",
						Summary:     "",
						Title:       "SSH private key in RSA format",
						IsRequired:  false,
					},
					Value: "$SSH_RSA_PRIVATE_KEY",
				},
				"ssh_key_save_path": {
					Identifier: "ssh_key_save_path",
					Options: XPutOptions{
						Description: "",
						Summary:     "",
						Title:       "(Optional) path to save the private key.",
						IsRequired:  false,
					},
					Value: "$HOME/.ssh/bitrise_step_activate_ssh_key",
				},
				"is_remove_other_identities": {
					Identifier: "is_remove_other_identities",
					Options: XPutOptions{
						Description: `(Optional) Remove other or previously loaded keys and restart ssh-agent?

Options:

* "true"
* "false"`,
						Summary:    "",
						Title:      "Remove other identities?",
						IsRequired: false,
					},
					Value: "true",
				},
				"verbose": {
					Identifier: "verbose",
					Options: XPutOptions{
						Description: "Enable verbose log option for better debug",
						Summary:     "Enable verbose log option for better debug",
						Title:       "Enable verbose logging",
						IsRequired:  true,
					},
					Value: false,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CollectXPuts(tt.args.models)
			for _, value := range got {
				other := tt.want[value.Identifier]
				if !reflect.DeepEqual(value, other) {
					t.Errorf("CollectXPuts() = %v, want %v", value, other)
				}
			}
		})
	}
}
