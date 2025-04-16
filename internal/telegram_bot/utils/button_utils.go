package utils

import "fmt"

func GetCustomButtons(textList *[]string, rowLen int) *[][]string {
	fmt.Println(textList)
	var markupButtons [][]string
	var row = make([]string, 0)
	for index, item := range *textList {
		if index != 0 && index%rowLen == 0 {
			markupButtons = append(markupButtons, row)
			row = make([]string, 0)
		}
		row = append(row, item)
	}
	markupButtons = append(markupButtons, row)
	return &markupButtons
}

func AppendButtons(oldButtons *[][]string, newButtons *[][]string) *[][]string {
	outputButtons := make([][]string, 0)
	if oldButtons != nil && len(*oldButtons) > 0 {
		outputButtons = *oldButtons
	}

	for _, item := range *newButtons {
		outputButtons = append(outputButtons, item)
	}
	return &outputButtons
}
