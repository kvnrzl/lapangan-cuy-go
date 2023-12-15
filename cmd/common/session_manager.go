package common

import (
	"github.com/google/uuid"
)

type SessionManager struct{}

func (sm *SessionManager) CreateSession(userID uuid.UUID) {

}

func (sm *SessionManager) DeleteSession(sessionID string) {

}

func (sm *SessionManager) GetUserFromSession(sessionID string) {}

func (sm *SessionManager) RenewSession(sessionID string) {}

func (sm *SessionManager) SessionExist(sessionID string) {}
