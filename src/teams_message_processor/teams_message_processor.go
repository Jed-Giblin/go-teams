package teams_message_processor

import "io"

type TeamsMessageProcessor interface {
	onMessage( )
	parseMessage(body io.ReadCloser, v interface{})
}