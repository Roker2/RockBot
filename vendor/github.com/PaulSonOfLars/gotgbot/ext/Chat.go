package ext

type Chat struct {
	Bot              Bot              `json:"-"`
	Id               int              `json:"id"`
	Type             string           `json:"type"`
	Title            string           `json:"title"`
	Username         string           `json:"username"`
	FirstName        string           `json:"first_name"`
	LastName         string           `json:"last_name"`
	Photo            *ChatPhoto       `json:"photo"`
	Bio              string           `json:"bio"`
	Description      string           `json:"description"`
	InviteLink       string           `json:"invite_link"`
	PinnedMessage    *Message         `json:"pinned_message"`
	Permissions      *ChatPermissions `json:"permissions"`
	SlowModeDelay    int              `json:"slow_mode_delay"`
	StickerSetName   string           `json:"sticker_set_name"`
	CanSetStickerSet bool             `json:"can_set_sticker_set"`
	LinkedChatId     int              `json:"linked_chat_id"`
	Location         ChatLocation     `json:"location"`
}

type ChatPermissions struct {
	CanSendMessages       *bool `json:"can_send_messages,omitempty"`
	CanSendMediaMessages  *bool `json:"can_send_media_messages,omitempty"`
	CanSendPolls          *bool `json:"can_send_polls,omitempty"`
	CanSendOtherMessages  *bool `json:"can_send_other_messages,omitempty"`
	CanAddWebPagePreviews *bool `json:"can_add_web_page_previews,omitempty"`
	CanChangeInfo         *bool `json:"can_change_info,omitempty"`
	CanInviteUsers        *bool `json:"can_invite_users,omitempty"`
	CanPinMessages        *bool `json:"can_pin_messages,omitempty"`
}

type ChatLocation struct {
	Location Location `json:"location"`
	Address  string   `json:"address"`
}

type ChatPhoto struct {
	SmallFileId       string `json:"small_file_id"`
	SmallFileUniqueId string `json:"small_file_unique_id"`
	BigFileId         string `json:"big_file_id"`
	BigFileUniqueId   string `json:"big_file_unique_id"`
}

type ChatMember struct {
	User                  *User  `json:"user"`
	Status                string `json:"status"`
	CustomTitle           string `json:"custom_title"`
	IsAnonymous           bool   `json:"is_anonymous"`
	CanBeEdited           bool   `json:"can_be_edited"`
	CanPostMessages       bool   `json:"can_post_messages"`
	CanEditMessages       bool   `json:"can_edit_messages"`
	CanDeleteMessages     bool   `json:"can_delete_messages"`
	CanRestrictMembers    bool   `json:"can_restrict_members"`
	CanPromoteMembers     bool   `json:"can_promote_members"`
	CanChangeInfo         bool   `json:"can_change_info"`
	CanInviteUsers        bool   `json:"can_invite_users"`
	CanPinMessages        bool   `json:"can_pin_messages"`
	IsMember              bool   `json:"is_member"`
	CanSendMessages       bool   `json:"can_send_messages"`
	CanSendMediaMessages  bool   `json:"can_send_media_messages"`
	CanSendPolls          bool   `json:"can_send_polls"`
	CanSendOtherMessages  bool   `json:"can_send_other_messages"`
	CanAddWebPagePreviews bool   `json:"can_add_web_page_previews"`
	UntilDate             int64  `json:"until_date"`
}

func (chat Chat) SendAction(action string) (bool, error) {
	return chat.Bot.SendChatAction(chat.Id, action)
}

func (chat Chat) KickMember(userId int) (bool, error) {
	return chat.Bot.KickChatMember(chat.Id, userId)
}

func (chat Chat) UnbanMember(userId int) (bool, error) {
	return chat.Bot.UnbanChatMember(chat.Id, userId)
}

func (chat Chat) RestrictMember(userId int) (bool, error) {
	return chat.Bot.RestrictChatMember(chat.Id, userId)
}

func (chat Chat) UnRestrictMember(userId int) (bool, error) {
	return chat.Bot.UnRestrictChatMember(chat.Id, userId)
}

func (chat Chat) PromoteMember(userId int) (bool, error) {
	return chat.Bot.PromoteChatMember(chat.Id, userId)
}

func (chat Chat) SetAdministratorCustomTitle(userId int, customTitle string) (bool, error) {
	return chat.Bot.SetChatAdministratorCustomTitle(chat.Id, userId, customTitle)
}

func (chat Chat) DemoteMember(userId int) (bool, error) {
	return chat.Bot.DemoteChatMember(chat.Id, userId)
}

func (chat Chat) SetChatPermissions(perms ChatPermissions) (bool, error) {
	return chat.Bot.SetChatPermissions(chat.Id, perms)
}

func (chat Chat) ExportInviteLink() (string, error) {
	return chat.Bot.ExportChatInviteLink(chat.Id)
}

func (chat Chat) SetChatPhoto(photo InputFile) (bool, error) {
	return chat.Bot.SetChatPhoto(chat.Id, photo)
}

func (chat Chat) DeletePhoto() (bool, error) {
	return chat.Bot.DeleteChatPhoto(chat.Id)
}

func (chat Chat) SetTitle(title string) (bool, error) {
	return chat.Bot.SetChatTitle(chat.Id, title)
}

func (chat Chat) SetDescription(description string) (bool, error) {
	return chat.Bot.SetChatDescription(chat.Id, description)
}

func (chat Chat) PinMessage(messageId int) (bool, error) {
	return chat.Bot.PinChatMessage(chat.Id, messageId)
}

func (chat Chat) PinMessageQuiet(messageId int) (bool, error) {
	return chat.Bot.PinChatMessageQuiet(chat.Id, messageId)
}

func (chat Chat) UnpinMessage() (bool, error) {
	return chat.Bot.UnpinChatMessage(chat.Id)
}

func (chat Chat) UnpinMessageById(messageId int) (bool, error) {
	return chat.Bot.UnpinChatMessageById(chat.Id, messageId)
}

func (chat Chat) UnpinAll() (bool, error) {
	return chat.Bot.UnpinAllChatMessages(chat.Id)
}

func (chat Chat) Leave() (bool, error) {
	return chat.Bot.LeaveChat(chat.Id)
}

func (chat Chat) Get() (*Chat, error) {
	return chat.Bot.GetChat(chat.Id)
}

func (chat Chat) GetAdministrators() ([]ChatMember, error) {
	return chat.Bot.GetChatAdministrators(chat.Id)
}

func (chat Chat) GetMembersCount() (int, error) {
	return chat.Bot.GetChatMembersCount(chat.Id)
}

func (chat Chat) GetMember(userId int) (*ChatMember, error) {
	return chat.Bot.GetChatMember(chat.Id, userId)
}

func (chat Chat) SetStickerSet(stickerSetName string) (bool, error) {
	return chat.Bot.SetChatStickerSet(chat.Id, stickerSetName)
}

func (chat Chat) DeleteStickerSet() (bool, error) {
	return chat.Bot.DeleteChatStickerSet(chat.Id)
}

func (chat Chat) DeleteMessage(messageId int) (bool, error) {
	return chat.Bot.DeleteMessage(chat.Id, messageId)
}
