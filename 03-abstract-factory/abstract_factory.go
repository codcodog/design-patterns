package abstractfactory

type AbstractFactory interface {
	createButton() Button
	createBorder() Border
}

type Button interface {
	createButton() string
}
type Border interface {
	createBorder() string
}

type WinFactory struct {
}

func (w *WinFactory) createButton() Button {
	return &WinButton{}
}
func (w *WinFactory) createBorder() Border {
	return &WinBorder{}
}

type MacFactory struct {
}

func (m *MacFactory) createButton() Button {
	return &MacButton{}
}
func (m *MacFactory) createBorder() Border {
	return &MacBorder{}
}

type WinButton struct {
}
type WinBorder struct {
}

func (w *WinButton) createButton() string {
	return "win button"
}

func (w *WinBorder) createBorder() string {
	return "win border"
}

type MacButton struct {
}
type MacBorder struct {
}

func (m *MacButton) createButton() string {
	return "mac button"
}
func (m *MacBorder) createBorder() string {
	return "mac border"
}
