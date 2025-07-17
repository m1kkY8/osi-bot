# osi-bot

A Discord bot written in Go for team and user management, leaderboard display, and bookstack (documentation) integration. Designed primarily for Hack The Box (HTB) teams, it provides slash commands, interactive components, and admin workflows.

## Features

- **Discord Slash Commands**: Manage users, teams, and permissions via Discord's slash command interface.
- **HTB Team Integration**: Handles HTB team join requests and approvals, with admin role checks.
- **Bookstack Integration**: Lists and manages Bookstack users interactively within Discord.
- **Leaderboard**: Displays and pages through leaderboard data interactively.
- **Component Interactions**: Supports Discord buttons and component-based UI for richer interaction.
- **Admin Controls**: Role-based access for sensitive commands.

## Getting Started

### Prerequisites

- Go 1.24+
- Discord bot token
- (Optional) Docker for containerized deployment

### Environment Variables

Set the following variables (e.g., via `.env` or your deployment environment):

- `DISCORD_TOKEN` – Your Discord bot token (required)
- `GUILD_ID` – (Optional) Guild/server ID for registering commands server-specific
- `ADMIN_ROLE_ID` – Discord Role ID for bot admins (for admin command access)
- `HTB_TEAM_ID` – Team ID for HTB integration
- `HTB_TOKEN` – API token for Hack The Box (if using team integration)

### Running Locally

```sh
go run ./cmd/bot/main.go
```

Or use the Makefile:

```sh
make run
```

### Docker

Build and run the container:

```sh
docker build -t osi-bot .
docker run -e DISCORD_TOKEN=... -e ADMIN_ROLE_ID=... -e HTB_TEAM_ID=... -e HTB_TOKEN=... osi-bot
```

### CI/CD

A `Jenkinsfile` is included for pipeline automation and Docker Hub deployment.

## Usage

Once invited to a Discord server:

- Use `/alexandria` commands for Bookstack user management.
- Use `/team` commands for viewing and handling HTB join requests.
- Admins can approve/reject join requests and manage user permissions.
- Interactive components (buttons, paginated lists) are used for navigation.

## Development

- Main entry: [`cmd/bot/main.go`](cmd/bot/main.go)
- Customization is via [`pkg/models`], [`pkg/bot/handlers`], and [`pkg/api/htb`].
- All bot commands and interactions are registered at startup.

## Dependencies

- [bwmarrin/discordgo](https://github.com/bwmarrin/discordgo) for Discord API
- [joho/godotenv](https://github.com/joho/godotenv) for env files

## License

_This project currently does not have an explicit license. Please add one if you intend to distribute or use in production._

---

> Maintained by [m1kkY8](https://github.com/m1kkY8)

# osi-bot Commands Guide

This bot provides slash commands for managing your HTB Team and Alexandria book stack directly from Discord.  
Below is a guide for using each command.

---

## HTB Team Management Commands

All HTB team commands are grouped under the `/team` command.

### `/team invitations`

- **Description:** Get all pending join requests to your HTB team.
- **Usage:**  
  `/team invitations`  
  _Replies with a list of users who have requested to join your team._

### `/team accept`

- **Description:** Accept a user’s join request.
- **Usage:**  
  `/team accept request_id:12345`  
  _Replace `12345` with the actual request ID._

### `/team reject`

- **Description:** Reject a user’s join request.
- **Usage:**  
  `/team reject request_id:12345`  
  _Replace `12345` with the actual request ID._

### `/team kick`

- **Description:** Remove a user from your team.
- **Usage:**  
  `/team kick user_id:54321`  
  _Replace `54321` with the user’s ID you want to remove._

### `/team leaderboard`

- **Description:** Show the team leaderboard, sorted by points.
- **Usage:**  
  `/team leaderboard`  
  _Replies with a paginated leaderboard of your team members._

---

## Alexandria Book Stack Commands

All Alexandria commands are grouped under the `/alexandria` command.

### `/alexandria register`

- **Description:** Register a Discord user in Alexandria.
- **Usage:**  
  `/alexandria register username:@discorduser`  
  _Select the user from the Discord user picker._

### `/alexandria update`

- **Description:** Change a user's permission level.
- **Usage:**  
  `/alexandria update user_id:123456789 role:viewer`  
  `/alexandria update user_id:123456789 role:editor`  
  _Replace `123456789` with the user's ID and choose the desired role._

### `/alexandria remove`

- **Description:** Remove a user from Alexandria.
- **Usage:**  
  `/alexandria remove user_id:123456789`  
  _Replace `123456789` with the user's ID._

### `/alexandria users`

- **Description:** List all registered users in Alexandria.
- **Usage:**  
  `/alexandria users`  
  _Replies with a list of all Alexandria users._

---

## General Notes

- All commands are available as Discord slash commands in servers where the bot is installed.
- You may need appropriate permissions to use some administrative commands (accept, reject, kick, update, remove).
- When using commands that require IDs, you can get request or user IDs from the `/team invitations` or `/alexandria users` output.

---

## Example Usage

```
/team invitations
/team accept request_id:39622
/team kick user_id:1184581
/alexandria register username:@m1kkY8
/alexandria update user_id:1184581 role:editor
/alexandria users
```

---

If you have any issues or questions, check the source code or open an issue in the repository.
