package handlers

import (
	"fuse/sessionmanager"
)

type MainHandler struct {
	manager *sessionmanager.SessionManager
}

func (mh *MainHandler) WithManager(m *sessionmanager.SessionManager) {
	mh.manager = m
}
