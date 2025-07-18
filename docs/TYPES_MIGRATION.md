# Types Package Migration

This document outlines the migration of types from the `models` package to a dedicated `types` package.

## New Structure

### `/pkg/types/`

- **`bookstack.go`** - BookStack user types and API request/response types
- **`team.go`** - HTB team member types and related structures
- **`discord.go`** - Discord client and pagination types
- **`constants.go`** - All application constants (domains, roles, button IDs, commands)
- **`http.go`** - HTTP response and error types
- **`factories.go`** - Factory functions for creating new instances
- **`methods.go`** - Methods for the Client type

## Key Types Moved

### Core Types

- `BookstackUser` → `types.BookstackUser`
- `BookstackUserResponse` → `types.BookstackUserResponse`
- `TeamMember` → `types.TeamMember`
- `Client` → `types.Client`
- `Page` → `types.Page`

### New Types Added

- `BookstackUserCreateRequest`
- `BookstackUserUpdateRequest`
- `TeamJoinRequest`
- `TeamRequestUser`
- `TeamInvitationsResponse`
- `TeamMembersResponse`
- `TeamActionResponse`
- `HTTPResponse`
- `AuthHeader`
- `APIError`

### Constants Consolidated

- Domain constants (`MAIL_DOMAIN`, `BOOKSTACK_DOMAIN`)
- Role constants (`ADMIN`, `EDITOR`, `VIEWER`, `PUBLIC`)
- Button IDs for Discord components
- Command and subcommand names

## Backward Compatibility

All models package files now re-export types from the types package:

```go
// pkg/models/bookstackUser.go
type BookstackUser = types.BookstackUser
type BookstackUserResponse = types.BookstackUserResponse

func CreateBookstackUser(name, email, password string) *BookstackUser {
    return types.NewBookstackUser(name, email, password)
}
```

This ensures existing code continues to work without modification.

## Benefits

1. **Separation of Concerns** - Pure types are separated from business logic
2. **Reusability** - Types can be used across packages without circular imports
3. **Centralized Constants** - All constants are in one location
4. **Better Documentation** - Types are clearly documented with their purpose
5. **Type Safety** - Dedicated request/response types for APIs

## Usage Examples

### Creating a new BookStack user:

```go
user := types.NewBookstackUser("john", "john@example.com", "password123")
```

### Using constants:

```go
user.Roles = []int{types.EDITOR}
url := types.BOOKSTACK_DOMAIN + "/api/users"
```

### Client initialization:

```go
client := types.NewClient(teamMembers, discordSession)
client.SetGuildID("123456789")
```

## Migration Status

✅ All core types migrated
✅ Backward compatibility maintained
✅ Build passes successfully
✅ Constants centralized
✅ Factory functions created
✅ Type-safe API request/response structures
