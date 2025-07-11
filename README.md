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
