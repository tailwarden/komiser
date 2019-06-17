package ovh

import (
	"fmt"

	ovhClient "github.com/ovh/go-ovh/ovh"
)

type Ticket struct {
	State string `json:"state"`
}

type TicketStat struct {
	Open  int `json:"open"`
	Close int `json:"close"`
}

func (ovh OVH) GetTickets() ([]int, error) {
	client, err := ovhClient.NewDefaultClient()
	if err != nil {
		return []int{}, err
	}

	tickets := []int{}
	err = client.Get("/support/tickets", &tickets)
	return tickets, err
}

func (ovh OVH) GetTicketsStats() (TicketStat, error) {
	stat := TicketStat{}

	client, err := ovhClient.NewDefaultClient()
	if err != nil {
		return stat, err
	}

	tickets, err := ovh.GetTickets()
	if err != nil {
		return stat, err
	}
	for _, ticketId := range tickets {
		ticket := Ticket{}
		err = client.Get(fmt.Sprintf("/support/tickets/%d", ticketId), &ticket)
		if err != nil {
			return stat, err
		}

		if ticket.State == "open" {
			stat.Open++
		} else {
			stat.Close++
		}

	}
	return stat, err
}
