package types

// TeamMember represents a member of a Hack The Box team
type TeamMember struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	RankText string `json:"rank_text"`
	UserOwns int    `json:"user_owns"`
	RootOwns int    `json:"root_owns"`
	Points   int    `json:"points"`
}

// TeamJoinRequest represents a request to join a team
type TeamJoinRequest struct {
	ID   int             `json:"id"`
	User TeamRequestUser `json:"user"`
}

// TeamRequestUser represents a user in a team join request
type TeamRequestUser struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// TeamInvitationsResponse represents the API response for team invitations
type TeamInvitationsResponse struct {
	Data []TeamJoinRequest `json:"data"`
}

// TeamMembersResponse represents the API response for team members
type TeamMembersResponse struct {
	Data []TeamMember `json:"data"`
}

// TeamActionResponse represents a generic team action response
type TeamActionResponse struct {
	Data struct {
		Message string `json:"message"`
	} `json:"data"`
}
