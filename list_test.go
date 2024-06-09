package loopfmt

import "testing"

func TestOrderedList_Render(t *testing.T) {
	tests := []struct {
		name  string
		items []string
		want  string
	}{
		{
			name:  "happy path",
			items: []string{"item1", "item2", "item3"},
			want: ` 1. item1
 1. item2
 1. item3`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ol := NewOrderedList()
			for _, item := range tt.items {
				ol.Append(item)
			}
			if got := ol.Render(); got != tt.want {
				t.Errorf("Render() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnorderedList_Render(t *testing.T) {
	tests := []struct {
		name  string
		items []string
		want  string
	}{
		{
			name:  "happy path",
			items: []string{"item1", "item2", "item3"},
			want: ` - item1
 - item2
 - item3`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ol := NewUnorderedList()
			for _, item := range tt.items {
				ol.Append(item)
			}
			if got := ol.Render(); got != tt.want {
				t.Errorf("Render() = %v, want %v", got, tt.want)
			}
		})
	}
}
