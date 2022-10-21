package squarecloud

type RawUser struct {
	Id            string `json:"id"`
	Tag           string `json:"tag"`
	Email         string `json:"email"`
	Plan          Plan   `json:"plan"`
	IsBlocklisted bool   `json:"blocklist"`
}

type User struct {
	Id            string        `json:"id"`
	Tag           string        `json:"tag"`
	Email         string        `json:"email"`
	Plan          Plan          `json:"plan"`
	IsBlocklisted bool          `json:"blocklist"`
	Applications  []Application `json:"applications"`
}

type Plan struct {
	Name     string       `json:"ultimate"`
	Duration PlanDuration `json:"duration"`
}

type PlanDuration struct {
	Formatted string `json:"formatted"`
	Raw       int    `json:"raw"`
}

type Application struct {
	Id        string `json:"id"`
	Tag       string `json:"tag"`
	Ram       int    `json:"ram"`
	Lang      string `json:"lang"`
	Type      string `json:"type"`
	Cluster   string `json:"cluster"`
	IsWebsite bool   `json:"isWebsite"`
	Avatar    string `json:"avatar"`
}
