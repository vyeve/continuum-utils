package rest

type Environment int

const (
	DTEnvironment Environment = iota + 1
	QAEnvironment
)

func (e Environment) String() string {
	switch e {
	case DTEnvironment:
		return "DT Environment"
	case QAEnvironment:
		return "QA Environment"
	}
	return "Undefined"
}
