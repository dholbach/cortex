package querier

import (
	"fmt"
	"testing"

	"github.com/prometheus/prometheus/promql"
)

var result *promql.Result

func BenchmarkChunkQueryable(b *testing.B) {
	for _, encoding := range encodings {
		store, from := makeMockChunkStore(b, 24*30, encoding.e)
		b.ResetTimer()

		for _, queryable := range queryables {
			for _, query := range queries {
				b.Run(fmt.Sprintf("%s/%s/%s", queryable.name, encoding.name, query.query), func(b *testing.B) {
					var r *promql.Result
					for n := 0; n < b.N; n++ {
						r = testQuery(b, queryable.f(store), from, query)
					}
					result = r
				})
			}
		}
	}
}
