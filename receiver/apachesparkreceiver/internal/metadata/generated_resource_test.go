// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResourceBuilder(t *testing.T) {
	for _, test := range []string{"default", "all_set", "none_set"} {
		t.Run(test, func(t *testing.T) {
			cfg := loadResourceAttributesConfig(t, test)
			rb := NewResourceBuilder(cfg)
			rb.SetSparkApplicationID("spark.application.id-val")
			rb.SetSparkApplicationName("spark.application.name-val")
			rb.SetSparkExecutorID("spark.executor.id-val")
			rb.SetSparkJobID(12)
			rb.SetSparkStageAttemptID(22)
			rb.SetSparkStageID(14)

			res := rb.Emit()
			assert.Equal(t, 0, rb.Emit().Attributes().Len()) // Second call should return 0

			switch test {
			case "default":
				assert.Equal(t, 5, res.Attributes().Len())
			case "all_set":
				assert.Equal(t, 6, res.Attributes().Len())
			case "none_set":
				assert.Equal(t, 0, res.Attributes().Len())
				return
			default:
				assert.Failf(t, "unexpected test case: %s", test)
			}

			val, ok := res.Attributes().Get("spark.application.id")
			assert.True(t, ok)
			if ok {
				assert.EqualValues(t, "spark.application.id-val", val.Str())
			}
			val, ok = res.Attributes().Get("spark.application.name")
			assert.True(t, ok)
			if ok {
				assert.EqualValues(t, "spark.application.name-val", val.Str())
			}
			val, ok = res.Attributes().Get("spark.executor.id")
			assert.True(t, ok)
			if ok {
				assert.EqualValues(t, "spark.executor.id-val", val.Str())
			}
			val, ok = res.Attributes().Get("spark.job.id")
			assert.True(t, ok)
			if ok {
				assert.EqualValues(t, 12, val.Int())
			}
			val, ok = res.Attributes().Get("spark.stage.attempt.id")
			assert.Equal(t, test == "all_set", ok)
			if ok {
				assert.EqualValues(t, 22, val.Int())
			}
			val, ok = res.Attributes().Get("spark.stage.id")
			assert.True(t, ok)
			if ok {
				assert.EqualValues(t, 14, val.Int())
			}
		})
	}
}