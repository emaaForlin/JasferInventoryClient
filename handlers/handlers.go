package handlers

import "log"

type Client struct {
	l    *log.Logger
	addr string
}

func NewClient(l *log.Logger, addr string) *Client {
	return &Client{l, addr}
}
