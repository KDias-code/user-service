package app

import "diplom/user-service/pkg/configs"

func Start(conf *configs.Configs) error {
	s := new(server)

	s.generate(conf)

	err := s.app.Listen(":" + conf.Port)
	s.logger.Info("server started")
	if err != nil {
		s.logger.Error("Failed to start server", "error", err)
		return err
	}

	return err
}
