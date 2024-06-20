package electronic

type Phone interface {
	Brand() string
	Model() string
	Type() string
}

type StationPhone interface {
	ButtonsCount() int
}

type Smartphone interface {
	OS() string
}

type applePhone struct {
	model string
}

func NewApplePhone(model string) *applePhone {
	return &applePhone{model: model}
}

func (ap *applePhone) Brand() string {
	return "apple"
}

func (ap *applePhone) Model() string {
	return ap.model
}

func (ap *applePhone) Type() string {
	return "smartphone"
}

func (ap *applePhone) OS() string {
	return "IOS"
}

type androidPhone struct {
	brand string
	model string
}

func NewAndroidPhone(brand, model string) *androidPhone {
	return &androidPhone{brand: brand, model: model}
}

func (an *androidPhone) Brand() string {
	return an.brand
}

func (an *androidPhone) Model() string {
	return an.model
}

func (an *androidPhone) Type() string {
	return "smartphone"
}

func (an *androidPhone) OS() string {
	return "Android"
}

type radioPhone struct {
	model        string
	buttonsCount int
}

func NewRadioPhone(model string, buttonsCount int) *radioPhone {
	return &radioPhone{model: model, buttonsCount: buttonsCount}
}

func (rp *radioPhone) Brand() string {
	return "Generic"
}

func (rp *radioPhone) Model() string {
	return rp.model
}

func (rp *radioPhone) Type() string {
	return "station"
}

func (rp *radioPhone) ButtonsCount() int {
	return rp.buttonsCount
}
