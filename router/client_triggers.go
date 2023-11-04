package main

import "log"

func (c *Client) triggerSearch(message []byte) {
	log.Printf("Search for %s\n", string(message))

}
