package utils_test

import (
	"testing"

	"github.com/smartxworks/cloudtower-go-sdk/v2/utils"
)

func TestCompareBigIntStrings(t *testing.T) {
	tests := []struct {
		name        string
		string1     *string
		string2     *string
		expected    int
		expectError bool
	}{
		{
			name:        "both nil",
			string1:     nil,
			string2:     nil,
			expected:    0,
			expectError: false,
		},
		{
			name:        "first nil",
			string1:     nil,
			string2:     strPtr("123"),
			expected:    -1,
			expectError: false,
		},
		{
			name:        "second nil",
			string1:     strPtr("123"),
			string2:     nil,
			expected:    1,
			expectError: false,
		},
		{
			name:        "equal values",
			string1:     strPtr("123"),
			string2:     strPtr("123"),
			expected:    0,
			expectError: false,
		},
		{
			name:        "first greater",
			string1:     strPtr("456"),
			string2:     strPtr("123"),
			expected:    1,
			expectError: false,
		},
		{
			name:        "second greater",
			string1:     strPtr("123"),
			string2:     strPtr("456"),
			expected:    -1,
			expectError: false,
		},
		{
			name:        "large numbers",
			string1:     strPtr("999999999999999999999999999999"),
			string2:     strPtr("999999999999999999999999999998"),
			expected:    1,
			expectError: false,
		},
		{
			name:        "zero and positive",
			string1:     strPtr("0"),
			string2:     strPtr("1"),
			expected:    -1,
			expectError: false,
		},
		{
			name:        "same length different values",
			string1:     strPtr("1000"),
			string2:     strPtr("1001"),
			expected:    -1,
			expectError: false,
		},
		{
			name:        "different length values",
			string1:     strPtr("1000"),
			string2:     strPtr("999"),
			expected:    1,
			expectError: false,
		},
		// 非数字测试用例
		{
			name:        "first non-numeric",
			string1:     strPtr("abc"),
			string2:     strPtr("123"),
			expected:    0,
			expectError: true,
		},
		{
			name:        "second non-numeric",
			string1:     strPtr("123"),
			string2:     strPtr("def"),
			expected:    0,
			expectError: true,
		},
		{
			name:        "both non-numeric",
			string1:     strPtr("abc"),
			string2:     strPtr("def"),
			expected:    0,
			expectError: true,
		},
		{
			name:        "special characters",
			string1:     strPtr("123!@#"),
			string2:     strPtr("456"),
			expected:    0,
			expectError: true,
		},
		{
			name:        "empty strings",
			string1:     strPtr(""),
			string2:     strPtr(""),
			expected:    0,
			expectError: true,
		},
		{
			name:        "whitespace",
			string1:     strPtr(" 123 "),
			string2:     strPtr("123"),
			expected:    0,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := utils.CompareBigIntStrings(tt.string1, tt.string2)

			if tt.expectError && err == nil {
				t.Errorf("compareBigIntStrings(%v, %v) expected error but got nil",
					stringPtrValue(tt.string1),
					stringPtrValue(tt.string2))
				return
			}
			if !tt.expectError && err != nil {
				t.Errorf("compareBigIntStrings(%v, %v) unexpected error: %v",
					stringPtrValue(tt.string1),
					stringPtrValue(tt.string2),
					err)
				return
			}

			if tt.expectError {
				return
			}

			if result != tt.expected {
				t.Errorf("compareBigIntStrings(%v, %v) = %v, want %v",
					stringPtrValue(tt.string1),
					stringPtrValue(tt.string2),
					result,
					tt.expected)
			}
		})
	}
}

// 辅助函数：创建字符串指针
func strPtr(s string) *string {
	return &s
}

// 辅助函数：安全地获取字符串指针的值
func stringPtrValue(s *string) string {
	if s == nil {
		return "nil"
	}
	return *s
}
