package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinuxParser_ParseVolume(t *testing.T) {
	testCases := map[string]struct {
		volumeSpec    string
		expectedParts *Volume
		expectedError error
	}{
		"empty": {
			volumeSpec:    "",
			expectedError: NewInvalidVolumeSpecErr(""),
		},
		"destination only": {
			volumeSpec:    "/destination",
			expectedParts: &Volume{Destination: "/destination"},
		},
		"source and destination": {
			volumeSpec:    "/source:/destination",
			expectedParts: &Volume{Source: "/source", Destination: "/destination"},
		},
		"destination and mode": {
			volumeSpec:    "/destination:rw",
			expectedParts: &Volume{Destination: "/destination", Mode: "rw"},
		},
		"all values": {
			volumeSpec:    "/source:/destination:rw",
			expectedParts: &Volume{Source: "/source", Destination: "/destination", Mode: "rw"},
		},
		"read only": {
			volumeSpec:    "/source:/destination:ro",
			expectedParts: &Volume{Source: "/source", Destination: "/destination", Mode: "ro"},
		},
		"volume case sensitive": {
			volumeSpec:    "/Source:/Destination:rw",
			expectedParts: &Volume{Source: "/Source", Destination: "/Destination", Mode: "rw"},
		},
		"support SELinux label bind mount content is shared among multiple containers": {
			volumeSpec:    "/source:/destination:z",
			expectedParts: &Volume{Source: "/source", Destination: "/destination", Mode: "z"},
		},
		"support SELinux label bind mount content is private and unshare": {
			volumeSpec:    "/source:/destination:Z",
			expectedParts: &Volume{Source: "/source", Destination: "/destination", Mode: "Z"},
		},
		"unsupported mode": {
			volumeSpec:    "/source:/destination:T",
			expectedError: NewInvalidVolumeSpecErr("/source:/destination:T"),
		},
		"too much colons": {
			volumeSpec:    "/source:/destination:rw:something",
			expectedError: NewInvalidVolumeSpecErr("/source:/destination:rw:something"),
		},
		"invalid source": {
			volumeSpec:    ":/destination",
			expectedError: NewInvalidVolumeSpecErr(":/destination"),
		},
		"named source": {
			volumeSpec:    "volume_name:/destination",
			expectedParts: &Volume{Source: "volume_name", Destination: "/destination"},
		},
	}

	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			parser := NewLinuxParser()
			parts, err := parser.ParseVolume(testCase.volumeSpec)

			if testCase.expectedError == nil {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, testCase.expectedError.Error())
			}

			assert.Equal(t, testCase.expectedParts, parts)
		})
	}
}
