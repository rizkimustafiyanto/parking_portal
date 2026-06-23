package constans

type PaymentScenario string

const (
	ScenarioSuccess PaymentScenario = "SUCCESS"
	ScenarioFailure PaymentScenario = "FAILURE"
	ScenarioTimeout PaymentScenario = "TIMEOUT"
)