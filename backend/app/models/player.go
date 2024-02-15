package models

type Player struct {
	Name          string `json:"name"`
	Gender        string `json:"gender"`
	Id            string `json:"id"`
	PartnerId     string `json:"partner_id"`
	Event         string `json:"event"`
	Confederation string `json:"confederation"`
	Country       string `json:"country"`
	Points        int    `json:"points"`
	Tournaments   int    `json:"tournaments"`
	Rank          int    `json:"rank"`
}
