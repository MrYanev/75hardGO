package api

// define routes for endpoints

// flow of logic/inputs should go:
// main -> api -> models -> storage (txt files)

func routing(s *Service) {
	s.Router.Run("localhost:8080")
}
