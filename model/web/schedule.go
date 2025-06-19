package web

type CreateScheduleRequest struct {
	MovieID      int   `json:"movie_id" validate:"required"`
	CinemaID     int   `json:"cinema_id" validate:"required"`
	StartTime    string `json:"start_time" validate:"required,datetime=2006-01-02T15:04:05Z07:00"`
	EndTime      string `json:"end_time" validate:"required,datetime=2006-01-02T15:04:05Z07:00"`
	ScreenNumber int    `json:"screen_number" validate:"required"`
	IsCancel     bool   `json:"screan_number" validates:"required"`
}

type UpdateScheduleRequest struct {
	MovieID      int   `json:"movie_id" validate:"required"`
	CinemaID     int   `json:"cinema_id" validate:"required"`
	StartTime    string `json:"start_time" validate:"required,datetime=2006-01-02T15:04:05Z07:00"`
	EndTime      string `json:"end_time" validate:"required,datetime=2006-01-02T15:04:05Z07:00"`
	ScreenNumber int    `json:"screen_number" validate:"required"`
}
