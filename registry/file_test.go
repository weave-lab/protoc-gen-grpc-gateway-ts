package registry

import "testing"

func Test_getAPIGatewayPathPrefix(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"not in the schema directory",
			args{"shared/shared_something.proto"},
			"",
		},
		{
			"in the schema directory - one level",
			args{"schemas/feature-flags/feature_flags_service_api.proto"},
			"/feature-flags",
		},
		{
			"in the schema directory - multiple levels",
			args{"schemas/feature-flags/v2/feature_flags_service_api.proto"},
			"/feature-flags/v2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getAPIGatewayPathPrefix(tt.args.fileName); got != tt.want {
				t.Errorf("getAPIGatewayPathPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
