package main

type User struct {
	ID      uint     `json:"id" gorm:"primary_key"`
	Name    string   `json:"name" gorm:"size:30"`
	Guilds  []*Guild `gorm:"many2many:user_guilds"`
	Friends []*User  `gorm:"many2many:user_friends"`
}

type Guild struct {
	ID       uint      `json:"id" gorm:"primary_key"`
	Name     string    `json:"name" gorm:"size:30"`
	Users    []*User   `gorm:"many2many:user_guilds"`
	Channels []Channel `gorm:"many2many:guild_channels"`
}

type Channel struct {
	ID      uint   `json:"id" gorm:"primary_key"`
	Name    string `json:"name" gorm:"size:30"`
	GuildID uint
}

type Message struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	Content   string `json:"content"`
	UserID    uint
	User      User
	ChannelID uint
	Channel   Channel
}
