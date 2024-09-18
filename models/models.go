package models

type Player struct {
	PlayerID   uint        `json:"player_id" gorm:"primaryKey"`
	Name       string      `json:"name" binding:"required"`
	Characters []Character `json:"characters"`
}

type Character struct {
	CharacterID uint   `json:"character_id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Race        string `json:"race"`
	Class       string `json:"class"`
	Level       int    `json:"level"`
	PlayerID    uint   `json:"player_id"`
}

type Adventure struct {
	AdventureID uint      `json:"adventure_id" gorm:"primaryKey"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartDate   string    `json:"start_date"`
	EndDate     string    `json:"end_date"`
	Sessions    []Session `json:"sessions"`
}

type Session struct {
	SessionID    uint   `json:"session_id" gorm:"primaryKey"`
	AdventureID  uint   `json:"adventure_id"`
	SessionDate  string `json:"session_date"`
	SessionNotes string `json:"session_notes"`
}

type SessionParticipant struct {
	SessionID   uint `json:"session_id"`
	PlayerID    uint `json:"player_id"`
	CharacterID uint `json:"character_id"`
}
