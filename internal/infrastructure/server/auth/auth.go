package auth

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"sync"
	"time"
)

type MemoryAuth struct {
	mutex  *sync.RWMutex
	memory map[string]string
}

func NewMemoryAuth() *MemoryAuth {
	return &MemoryAuth{
		mutex:  &sync.RWMutex{},
		memory: make(map[string]string),
	}
}

func (m *MemoryAuth) CreateAuth(login, pass string) string {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	time64 := time.Now().Unix()
	timeInt := strconv.FormatInt(time64, 10)
	token := login + pass + timeInt
	hashToken := sha256.Sum256([]byte(token))
	hashedToken := hex.EncodeToString(hashToken[:])
	m.memory[hashedToken] = login

	return hashedToken
}

func (m *MemoryAuth) FetchAuth(token string) (string, bool) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	login, ok := m.memory[token]
	if !ok {
		return "", false
	}
	return login, true
}

func (m *MemoryAuth) DeleteAuth(token string) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	delete(m.memory, token)
}
