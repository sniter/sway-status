package battery

import (
	"bytes"
	"fmt"
	"log"
	"text/template"
)

type BatteryStatus int

const (
	BatteryCharging BatteryStatus = iota
	BatteryFull
	BatteryNotCharging
	BatteryDischarging
	BatteryUnknownStatus
)

func ParseBatteryStatus(statusStr string) (BatteryStatus, error) {
	switch statusStr {
	case "Charging":
		return BatteryCharging, nil
	case "Full":
		return BatteryFull, nil
	case "Not charging":
		return BatteryNotCharging, nil
	case "Discharging":
		return BatteryDischarging, nil
	default:
		return BatteryUnknownStatus, fmt.Errorf("unknown battery status: %s", statusStr)
	}
}

type BatteryInfo interface {
	Report() BatteryReport
}

type BatteryReport struct {
	Capacity int
	Status   BatteryStatus
}

type Battery struct {
	Provider    BatteryInfo
	Template    *template.Template
	LabelFormat string
	Name        string
	Instance    string
}

const DefaultTemplate string = " {{ .Status }}{{ .Capacity }}% "

type BatteryTemplate *template.Template

// func BatteryStatusIcon(bs BatteryStatus) string = ???

func MakeBatteryTemplate(blueprint string, toStatusIcon func(BatteryStatus) string, toCapacityIcon func(int) string) BatteryTemplate {
	funcMap := template.FuncMap{
		"toStatusIcon":   toStatusIcon,
		"toCapacityIcon": toCapacityIcon,
	}
	tplBuilder := template.New("Battery").Funcs(funcMap)
	if blueprint != "" {
		tpl, err := tplBuilder.Parse(blueprint)
		if err != nil {
			log.Fatal(err)
		}
		return BatteryTemplate(tpl)
	}
	tpl, err := tplBuilder.Parse(DefaultTemplate)
	if err != nil {
		log.Fatal(err)
	}
	return BatteryTemplate(tpl)
}

func (b Battery) GetName() string     { return b.Name }
func (b Battery) GetInstance() string { return b.Instance }
func (b Battery) GetFullText(_ []byte) (string, error) {
	var buf bytes.Buffer
	err := b.Template.Execute(&buf, b.Provider)
	if err != nil {
		return err.Error(), err
	}
	return buf.String(), nil

	// tpl.Execute(wr io.Writer, data any)
	// capacity, err := b.Provider.Capacity()
	// if err != nil {
	// 	return "", err
	// }
	// status, err := b.Provider.Status()
	// if err != nil {
	// 	return "", err
	// }
	// value, err := b.Renderer.Render(capacity, status)
	// if err != nil {
	// 	return "", err
	// }
	// return fmt.Sprintf(b.LabelFormat, value), nil
}
