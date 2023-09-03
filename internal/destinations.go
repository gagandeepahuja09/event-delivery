package internal

import "fmt"

const (
	salesforce       = "salesforce"
	snowflake        = "snowflake"
	marketo          = "marketo"
	google_analytics = "google_analytics"
)

var supportedDestinations = []string{
	salesforce,
	snowflake,
	marketo,
	google_analytics,
}

func getChannelName(destination string) string {
	return fmt.Sprintf("%s_channel", destination)
}

var destinationToErrorPercentageMap = map[string]int{
	salesforce:       0,
	google_analytics: 50,
	marketo:          100,
	snowflake:        0,
}
