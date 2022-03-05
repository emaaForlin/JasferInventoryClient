package handlers

import "log"

type Client struct {
	l      *log.Logger
	addr   string
	apikey string
}

func NewClient(l *log.Logger, addr string, apikey string) *Client {
	return &Client{l, addr, apikey}
}
