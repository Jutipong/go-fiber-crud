package chart

type (
	Service interface {
	}
	service struct{ repo Repository }
)

func NewService(repo Repository) Service {
	return &service{repo}
}
