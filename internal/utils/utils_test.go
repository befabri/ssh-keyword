package utils

import (
	"os"
	"ssh-keyword/internal/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsIP(t *testing.T) {
	testCases := []struct {
		name string
		ip   string
		want bool
	}{
		{"Valid IP v4", "192.168.1.1", true},
		{"Valid IP v6", "2001:db8::68", true},
		{"Invalid IP", "999.999.999.999", false},
		{"Empty String", "", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := IsIP(tc.ip)
			assert.Equal(t, tc.want, got, "IsIP(%q) should return %v", tc.ip, tc.want)
		})
	}
}

func TestPromptInput(t *testing.T) {
	cases := []struct {
		name      string
		input     string
		expected  string
		expectErr bool
	}{
		{"Normal Input", "test input\n", "test input", false},
		{"Leading Space", "  test input\n", "test input", false},
		{"Trailing Space", "test input  \n", "test input", false},
		{"Upper Case Quit", "QUIT\n", "", true},
		{"Lower Case Quit", "quit\n", "", true},
		{"Mixed Case Quit", "Quit\n", "", true},
		{"Spaces Around Quit", "  quit  \n", "", true},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a pipe to simulate stdin.
			r, w, _ := os.Pipe()

			// Set os.Stdin to our input source.
			origStdin := os.Stdin
			os.Stdin = r
			defer func() { os.Stdin = origStdin }() // Restore original Stdin after the test

			// Write input to pipe in a goroutine
			go func() {
				w.Write([]byte(tc.input))
				w.Close()
			}()

			result, err := PromptInput("Please enter some input:")
			if tc.expectErr {
				assert.Error(t, err, "Expected an error but got none")
			} else {
				assert.NoError(t, err, "Expected no error but got one")
				assert.Equal(t, tc.expected, result, "Expected output does not match actual output")
			}
		})
	}
}

func TestFindConnectionByIP(t *testing.T) {
	connections := []config.Connection{
		{
			IP:       "192.168.1.1",
			Default:  false,
			Keywords: []string{"home", "wifi"},
			User:     "admin",
			Port:     "22",
		},
		{
			IP:       "10.0.0.1",
			Default:  true,
			Keywords: []string{"office", "lan"},
			User:     "guest",
			Port:     "22",
		},
	}
	want := config.Connection{
		IP:       "192.168.1.1",
		Default:  false,
		Keywords: []string{"home", "wifi"},
		User:     "admin",
		Port:     "22",
	}
	got, found := FindConnectionByIP(connections, "192.168.1.1")
	assert.True(t, found)
	assert.Equal(t, want, got)
}

func TestFindConnectionIndex(t *testing.T) {
	connections := []config.Connection{
		{IP: "192.168.1.1", Default: false, Keywords: []string{"home", "wifi"}, User: "admin", Port: "22"},
		{IP: "10.0.0.1", Default: true, Keywords: []string{"office", "lan"}, User: "guest", Port: "22"},
	}
	testCases := []struct {
		name string
		ip   string
		want int
	}{
		{"Found at index 0", "192.168.1.1", 0},
		{"Found at index 1", "10.0.0.1", 1},
		{"Not found", "127.0.0.1", -1},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := FindConnectionIndex(connections, tc.ip)
			assert.Equal(t, tc.want, got, "FindConnectionIndex(%s)", tc.ip)
		})
	}
}

func TestFindConnectionByKeyword(t *testing.T) {
	connections := []config.Connection{
		{IP: "192.168.1.1", Default: false, Keywords: []string{"home", "wifi"}, User: "admin", Port: "22"},
		{IP: "10.0.0.1", Default: true, Keywords: []string{"office", "lan"}, User: "guest", Port: "22"},
	}
	testCases := []struct {
		name    string
		keyword string
		want    config.Connection
		found   bool
	}{
		{"Found by keyword home", "home", connections[0], true},
		{"Found by keyword lan", "LAN", connections[1], true},
		{"Not found by keyword", "server", config.Connection{}, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, found := FindConnectionByKeyword(connections, tc.keyword)
			assert.Equal(t, tc.found, found, "Expected found status to match for keyword: %s", tc.keyword)
			if found {
				assert.Equal(t, tc.want.IP, got.IP, "Expected IP to match for keyword: %s", tc.keyword)
				assert.Equal(t, tc.want, got, "Expected connection object to match for keyword: %s", tc.keyword)
			}
		})
	}
}

func TestFindConnectionDefault(t *testing.T) {
	connectionsOne := []config.Connection{
		{IP: "192.168.1.1", Default: false, Keywords: []string{"home", "wifi"}, User: "admin", Port: "22"},
		{IP: "10.0.0.1", Default: false, Keywords: []string{"office", "lan"}, User: "guest", Port: "22"},
	}
	connectionsTwo := []config.Connection{
		{IP: "192.168.50.50", Default: false, Keywords: []string{"home", "wifi"}, User: "admin", Port: "22"},
		{IP: "172.16.23.202", Default: true, Keywords: []string{"office", "lan"}, User: "guest", Port: "2223"},
	}
	testCases := []struct {
		name        string
		connections []config.Connection
		want        config.Connection
		found       bool
	}{
		{"No default connection", connectionsOne, config.Connection{}, false},
		{"Default connection found", connectionsTwo, connectionsTwo[1], true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, found := FindConnectionDefault(tc.connections)
			assert.Equal(t, tc.found, found, "Expected found status to match")
			if found {
				assert.Equal(t, tc.want, got, "Expected default connection to match")
			}
		})
	}
}

func TestContains(t *testing.T) {
	testCases := []struct {
		name     string
		slice    []string
		str      string
		expected bool
	}{
		{"String present", []string{"apple", "banana", "cherry"}, "banana", true},
		{"String absent", []string{"apple", "banana", "cherry"}, "orange", false},
		{"Empty slice", []string{}, "banana", false},
		{"Empty string search", []string{"apple", "banana", ""}, "", true},
		{"All empty", []string{}, "", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := Contains(tc.slice, tc.str)
			assert.Equal(t, tc.expected, got, "Expect result to match expected for Contains")
		})
	}
}
