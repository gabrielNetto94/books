package domain

import (
	"testing"
)

func TestBookValidate(t *testing.T) {
	tests := []struct {
		name    string
		book    Book
		wantErr bool
	}{
		{
			name: "valid book",
			book: Book{
				Id:     "123e4567-e89b-12d3-a456-426614174000",
				Title:  "Test Title",
				Author: "Test Author",
				Desc:   "Test Description",
			},
			wantErr: false,
		},
		{
			name: "invalid author",
			book: Book{
				Id:     "123e4567-e89b-12d3-a456-426614174000",
				Title:  "Test Title",
				Author: "",
				Desc:   "Test Description",
			},
			wantErr: true,
		},
		{
			name: "invalid title",
			book: Book{
				Id:     "123e4567-e89b-12d3-a456-426614174000",
				Title:  "",
				Author: "Test Author",
				Desc:   "Test Description",
			},
			wantErr: true,
		},
		{
			name: "invalid description",
			book: Book{
				Id:     "123e4567-e89b-12d3-a456-426614174000",
				Title:  "Test Title",
				Author: "Test Author",
				Desc:   "",
			},
			wantErr: true,
		},
		{
			name: "invalid ID",
			book: Book{
				Id:     " invalid-id",
				Title:  "Test Title",
				Author: "Test Author",
				Desc:   "Test Description",
			},
			wantErr: true,
		},
		{
			name: "multiple invalid fields",
			book: Book{
				Id:     " invalid-id",
				Title:  "",
				Author: "",
				Desc:   "",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.book.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Book.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
