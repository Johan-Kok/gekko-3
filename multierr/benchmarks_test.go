
package multierr

import (
	"errors"
	"fmt"
	"testing"
)

func BenchmarkAppend(b *testing.B) {
	errorTypes := []struct {
		name string
		err  error
	}{
		{
			name: "nil",
			err:  nil,
		},
		{
			name: "single error",
			err:  errors.New("test"),
		},
		{
			name: "multiple errors",
			err:  appendN(nil, errors.New("err"), 10),
		},
	}

	for _, initial := range errorTypes {
		for _, v := range errorTypes {
			msg := fmt.Sprintf("append %v to %v", v.name, initial.name)
			b.Run(msg, func(b *testing.B) {
				for _, appends := range []int{1, 2, 10} {
					b.Run(fmt.Sprint(appends), func(b *testing.B) {
						for i := 0; i < b.N; i++ {
							appendN(initial.err, v.err, appends)
						}
					})
				}
			})
		}
	}
}
