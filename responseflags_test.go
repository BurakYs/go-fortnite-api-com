package fortniteapi

import "testing"

func TestCombineFlags(t *testing.T) {
	tests := []struct {
		name     string
		input    []ResponseFlag
		expected ResponseFlag
	}{
		{
			name:     "no flags",
			input:    []ResponseFlag{},
			expected: 0,
		},
		{
			name:     "single flag",
			input:    []ResponseFlag{FlagIncludePaths},
			expected: FlagIncludePaths,
		},
		{
			name:     "multiple flags",
			input:    []ResponseFlag{FlagIncludePaths, FlagIncludeGameplayTags},
			expected: FlagIncludePaths | FlagIncludeGameplayTags,
		},
		{
			name:     "duplicate flags",
			input:    []ResponseFlag{FlagIncludePaths, FlagIncludePaths},
			expected: FlagIncludePaths,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CombineFlags(tt.input...)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestFlagAllConstant(t *testing.T) {
	expected := FlagIncludePaths | FlagIncludeGameplayTags | FlagIncludeShopHistory
	if FlagAll != expected {
		t.Errorf("FlagAll constant is incorrect. Expected %v, got %v", expected, FlagAll)
	}
}
