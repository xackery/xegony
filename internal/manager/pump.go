package manager

import (
	"github.com/xackery/xegony/model"
)

// do not call directly, handled by manager
func (m *Manager) pump() {
	for {
		logger := model.NewLogger()
		select {
		case queryReq := <-m.queryChan:
			logger.Debug().Interface("req", queryReq).Msg("log")
		}
	}
}
