package main

import (
	"github.com/maxence-charriere/go-app/v8/pkg/app"
)

// Generates the GUI objects for the columns in the table
func (p *MagicTable) UpdateColumnsNodes() error {
	var err error
	err = nil

	p.tableColumnNodes = []app.UI{}
	for columnCounter, columnMetadataResponse := range p.testDataAndMetaData.magicTableMetaData {
		// Check if column should be shown
		columnHeader := app.Th().
			Scope("col").
			Body(app.Text(columnMetadataResponse.GetColumnHeaderName() + p.SetSortIconForTableHeader(columnCounter))).
			OnClick(p.MyOnColumnClickWrapper(columnCounter))
		shouldBeShown := columnMetadataResponse.GetShouldBeVisible()
		if shouldBeShown == true {
			p.tableColumnNodes = append(p.tableColumnNodes, columnHeader)
		}
	}

	return err
}

// Generates the GUI objects for the rows in the table
func (p *MagicTable) UpdateRowNodes() error {
	var err error
	err = nil

	p.tableRowsNodes = []app.UI{}
	//Loop every row
	for rowCounter, i := range p.testDataAndMetaData.originalTestdataInstances {
		rowNodes := []app.UI{}
		// Loop every column and get value and check if column should be shown
		for _, columnMetadataResponse := range p.testDataAndMetaData.magicTableMetaData {
			columnName := columnMetadataResponse.GetColumnDataName()
			columnData := p.formatColumnData(i, columnName)
			shouldBeShown := columnMetadataResponse.GetShouldBeVisible()
			if shouldBeShown == true {
				rowNodes = append(rowNodes, app.Td().Body(app.Text(columnData)))
			}
		}

		// Set background color on selected row
		var rowSelectedColor string
		if i.UniqueId != p.uniqueRowSelected {
			rowSelectedColor = "#FFFFFF"
		} else {
			rowSelectedColor = "#E1EAFA"
			p.rowSelected = rowCounter
		}

		// Append values for each row
		p.tableRowsNodes = append(p.tableRowsNodes,
			app.Tr().
				Body(rowNodes...).
				OnClick(p.MyOnClickOnRowWrapper(rowCounter)).
				OnDblClick(p.MyOnDblClickOnRowWrapper(rowCounter)).
				Style("background-color", rowSelectedColor))

	}

	return err
}
