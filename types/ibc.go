package types

type IBC struct {
	SourceChain      ChainInfo        `json:"chain_1"`
	DestinationChain ChainInfo        `json:"chain_2"`
	Channels         []ChannelElement `json:"channels"`
}

type ChainInfo struct {
	ChainName    string `json:"chain_name"`
	ClientID     string `json:"client_id"`
	ConnectionID string `json:"connection_id"`
}

type ChannelElement struct {
	SourceChain      ChannelInfo `json:"chain_1"`
	DestinationChain ChannelInfo `json:"chain_2"`
	Ordering         Ordering    `json:"ordering"`
	Version          string      `json:"version"`
	Description      *string     `json:"description,omitempty"`
	Tags             *Tags       `json:"tags,omitempty"`
}

type ChannelInfo struct {
	ChannelID    string  `json:"channel_id"`
	PortID       string  `json:"port_id"`
	ClientID     *string `json:"client_id,omitempty"`
	ConnectionID *string `json:"connection_id,omitempty"`
}

type Ordering string

const (
	Ordered   Ordering = "ordered"
	Unordered Ordering = "unordered"
)

type Tags struct {
	Status     *Status `json:"status,omitempty"`
	Preferred  *bool   `json:"preferred,omitempty"`
	Dex        *string `json:"dex,omitempty"`
	Properties *string `json:"properties,omitempty"`
}
