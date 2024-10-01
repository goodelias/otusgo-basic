package fixapp

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPrintStaff(t *testing.T) {
	testCases := []struct {
		name    string
		staff   []Employee
		want    string
		wantErr error
	}{
		{
			name: "with multiply employees",
			staff: []Employee{
				NewEmployee(2, 30, "Alec", 12),
				NewEmployee(1, 23, "Ilya", 11),
			},
			want: "User ID: 1; Age: 23; Name: Ilya; Department ID: 11; \n" +
				"User ID: 2; Age: 30; Name: Alec; Department ID: 12; \n",
			wantErr: nil,
		},
		{
			name:    "with no employees",
			staff:   []Employee{},
			want:    "",
			wantErr: ErrEmptyStaff,
		},
		{
			name: "with single employee",
			staff: []Employee{
				NewEmployee(1, 52, "Jake", 101),
			},
			want:    "User ID: 1; Age: 52; Name: Jake; Department ID: 101; \n",
			wantErr: nil,
		},
		{
			name: "with zero values",
			staff: []Employee{
				NewEmployee(2, 0, "", 0),
			},
			want:    "",
			wantErr: ErrZeroOrNegativeField,
		},
		{
			name:    "with nil employee",
			staff:   nil,
			want:    "",
			wantErr: ErrEmptyStaff,
		},
		{
			name: "with duplicate employees",
			staff: []Employee{
				NewEmployee(1, 23, "Ilya", 11),
				NewEmployee(2, 30, "Alec", 12),
				NewEmployee(1, 23, "Ilya", 11),
			},
			want: "User ID: 1; Age: 23; Name: Ilya; Department ID: 11; \n" +
				"User ID: 2; Age: 30; Name: Alec; Department ID: 12; \n",
			wantErr: nil,
		},
		{
			name: "with negative values",
			staff: []Employee{
				NewEmployee(-1, -23, "Ilya", 11),
			},
			want:    "",
			wantErr: ErrZeroOrNegativeField,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			got, err := PrintStaff(test.staff)
			if test.wantErr != nil {
				assert.Error(t, err)
				assert.Empty(t, got)
				require.Contains(t, err.Error(), test.wantErr.Error())
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, test.want, got)
		})
	}
}
