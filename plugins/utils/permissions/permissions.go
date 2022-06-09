package permissions

import "github.com/bwmarrin/discordgo"

// MemberHasPermission checks if a member has the given permission
func MemberHasPermission(s *discordgo.Session, channelID string, userID string, permission int64) (bool, error) {
	p, err := s.UserChannelPermissions(userID, channelID)
	if err != nil {
		return false, err
	}

	if p&permission == permission {
		return true, nil
	}

	return false, nil
}

// ComesFromDM returns true if a message comes from a DM channel
func ComesFromDM(s *discordgo.Session, m *discordgo.MessageCreate) (bool, error) {
	channel, err := s.State.Channel(m.ChannelID)
	if err != nil {
		if channel, err = s.Channel(m.ChannelID); err != nil {
			return false, err
		}
	}

	return channel.Type == discordgo.ChannelTypeDM, nil
}
