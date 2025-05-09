package valueobject

type UserField string

var (
	UserFieldID          = UserField("id")
	UserFieldEmail       = UserField("email")
	UserFieldPassword    = UserField("password")
	UserFieldDisplayName = UserField("display_name")
	UserFieldAvatarURL   = UserField("avatar_url")
	UserFieldUpdatedAt   = UserField("updated_at")
)

func (f UserField) String() string {
	return string(f)
}
