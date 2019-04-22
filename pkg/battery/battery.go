package battery

import (
	"fmt"

	"github.com/distatus/battery"
)

// Percentage returns the main battery's current level
// TODO: State arrows (down discharging, up charging)
func Percentage() string {
	battery, err := battery.Get(0)
	if err != nil {
		return ""
	}

	percentage := (battery.Current / battery.Full) * float64(100)
	if percentage < 20 {
		return fmt.Sprintf("%.0f%% ", percentage)
	}
	return ""
}
