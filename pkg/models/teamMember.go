package models

type TeamMember struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	RankText string `json:"rank_text"`
	UserOwns int    `json:"user_owns"`
	RootOwns int    `json:"root_owns"`
	Points   int    `json:"points"`
}
