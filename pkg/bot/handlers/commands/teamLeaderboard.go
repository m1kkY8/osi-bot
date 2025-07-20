package commands

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/gubarz/gohtb/services/teams"
	"github.com/m1kkY8/osi-bot/pkg/bot/embeds"
	"github.com/m1kkY8/osi-bot/pkg/types"
	"github.com/m1kkY8/osi-bot/pkg/util"
)

func customModel(members teams.MembersResponse) []types.TeamMember {
	var teamMembers []types.TeamMember
	for _, member := range members.Data {
		teamMember := types.TeamMember{
			ID:       member.Id,
			Name:     member.Name,
			UserOwns: member.UserOwns,
			RootOwns: member.RootOwns,
			Points:   member.Points,
			RankText: member.RankText,
		}
		teamMembers = append(teamMembers, teamMember)
	}

	// Sort by points descending
	sort.Slice(teamMembers, func(i, j int) bool {
		return teamMembers[i].Points > teamMembers[j].Points
	})
	return teamMembers
}

func teamLeaderboardSlashHandler(client *types.Client, pages *types.Page) func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		// convert string to int
		teamID, _ := strconv.Atoi(client.GetTeamID())

		members, err := client.HTBClient.Teams.Team(int(teamID)).Members(client.Context)
		if err != nil {
			_ = util.RespondEphemeral(s, i.Interaction, "Error fetching leaderboard")
			return
		}

		client.TeamMembers = customModel(members)

		userID := i.Member.User.ID
		pages.PageMap[userID] = 0
		embed, buttons := embeds.LeaderboardEmbed(0, client.TeamMembers)

		err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds:     []*discordgo.MessageEmbed{embed},
				Components: buttons,
			},
		})
		if err != nil {
			fmt.Println("error sending a message:", err)
			return
		}
	}
}
