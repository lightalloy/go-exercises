package chessboard

import "testing"

func TestChessboard(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			got, err := Get(tc.width, tc.length)
			switch {
			case tc.expectError:
				if err == nil {
					t.Fatalf("Get(%d, %d) expected error, got: %q", tc.width, tc.length, got)
				}
			case err != nil:
				t.Fatalf("Get(%d, %d) returned error: %v, want: %q", tc.width, tc.length, err, tc.want)
			case got != tc.want:
				t.Fatalf("Get(%d, %d) = %q, want %q", tc.width, tc.length, got, tc.want)
			}
		})
	}
}

func BenchmarkChessboard(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			_, _ = Get(tc.width, tc.length)
		}
	}
}
