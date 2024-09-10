package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"strconv"
)

func main() {
	myApp := app.New()

	myApp.Settings().SetTheme(customTheme{})
	//myApp.Settings().SetTheme(&forcedVariant{Theme: theme.DefaultTheme(), variant: theme.VariantDark})
	myWindow := myApp.NewWindow("Тест на честность")
	myWindow.Resize(fyne.NewSize(800, 400))

	//text.Size().Min()
	elementCount := 1
	element := getElements(elementCount)

	score := 0

	//text := canvas.NewText(element["text"], color.Black)
	//text.TextSize = 36

	text := widget.NewLabel(element["text"])

	btnYes := widget.NewButton("Да", func() {
		if elementCount < 34 {
			elementCount += 1
			element := getElements(elementCount)
			text.SetText(element["text"])
			number, err := strconv.ParseInt(element["yes"], 8, 0)
			if err == nil {
				score += int(number)
				fmt.Println(score)
			}
		} else {
			text.SetText(getScore(score))
			myWindow.SetContent(container.NewVBox(
				text,
			))
		}
	})
	btnYes.Resize(fyne.NewSize(100, 70))
	btnYes.Move(fyne.NewPos(300, 200))
	contButtonYes := container.NewWithoutLayout(btnYes)

	btnNo := widget.NewButton("Нет", func() {
		if elementCount < 34 {
			elementCount += 1
			element := getElements(elementCount)
			text.SetText(element["text"])
			number, err := strconv.ParseInt(element["no"], 8, 0)
			if err == nil {
				score += int(number)
				fmt.Println(score)
			}
		} else {
			text.SetText(getScore(score))
			myWindow.SetContent(container.NewVBox(
				text,
			))
		}

	})
	btnNo.Resize(fyne.NewSize(100, 70))
	btnNo.Move(fyne.NewPos(420, 154))
	contButtonNo := container.NewWithoutLayout(btnNo)

	myWindow.SetContent(container.NewVBox(
		text,
		contButtonYes,
		contButtonNo,
	))
	myWindow.ShowAndRun()
	//myApp.Run()
}

func getScore(score int) string {
	switch {
	case 0 <= score && score <= 5:
		return "Очень низкий показатель по шкале \"Честность\". Свидетельствует о ярко выраженной склонности ко лжи, приукрашиванию себя. Также может свидетельствовать о низких показателях социального интеллекта. Если данный опросник применялся в составе батареи методик, то результаты по личностным опросникам следует признать недостоверными."
	case 6 <= score && score <= 13:
		return "Низкий показатель по шкале \"Честность\". Свидетельствует о значительной склонности ко лжи. Любит приукрашивать себя, своё поведение. Если опросник применялся в составе батареи методик, то результаты по личностным опросникам не обязательнос следует признавать недостоверными. Однако следует отнестить к их результатам вполне критично."
	case 14 <= score && score <= 29:
		return "Нормальный результат. Склонность ко лжи не выявлена. Может быть, изредка склонен приукрашивать себя, своё поведение, но в пределах нормы."
	case 30 <= score && score <= 34:
		return "Высокий результат по шкале \"Честность\". Такой высокий результат может быть связан не только с высокой личностной честностью, но и следствием других причин: преднамеренного искажения ответов, очень неверной самооценки. Следует осторожно отнестись к данному результату.\n\n"
	}
	return ""
}
