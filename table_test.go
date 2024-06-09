package loopfmt

import "testing"

func TestTable_Render(t *testing.T) {
	tests := []struct {
		name  string
		items [][]string
		want  string
	}{
		{
			name: "happy path",
			items: [][]string{
				{"column 1", "column 2", "column 3"},
				{"data 1-1", "data 1-2", "data 1-3"},
				{"data 2-1", "data 2-2", "data 2-3"},
				{"data 3-1", "data 3-2", "data 3-3"},
			},
			want: ` | column 1 | column 2 | column 3 | 
 | - | - | - | 
 | data 1-1 | data 1-2 | data 1-3 | 
 | data 2-1 | data 2-2 | data 2-3 | 
 | data 3-1 | data 3-2 | data 3-3 | 
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			table := NewTable()
			for i := range tt.items {
				table.Append(tt.items[i]...)
			}
			if got := table.Render(); got != tt.want {
				t.Errorf("Render() = %v, want %v", got, tt.want)
			}
		})
	}
}
