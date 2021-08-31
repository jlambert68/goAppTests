package server

import (
	"context"
	"fmt"
	"goAppTest1/cmd/070_AddListDelete_TestDomains/protos/api"
)

func (server *Server) GetMagicTableMetadata(ctx context.Context, in *api.MagicTableMetadataRequest) (*api.MagicTableMetadataRespons, error) {

	var magicTableMetadataRespons *api.MagicTableMetadataRespons

	var magicTableColumnsMetadataResponse []*api.MagicTableColumnMetadata
	var err error

	magicTableColumnsMetadataFromDB, err := listMagicTableColumnMetadataForjsonFile("2")
	if err != nil {
		fmt.Println(err.Error())
		return magicTableMetadataRespons, err
	}

	//fmt.Println(magicTableColumnsMetadataFromDB)

	for _, magicTableColumnsMetadata := range magicTableColumnsMetadataFromDB {

		magicTableColumnMetadata := &api.MagicTableColumnMetadata{
			ColumnHeaderName:       magicTableColumnsMetadata.ColumnHeaderName,
			ColumnDataName:         magicTableColumnsMetadata.ColumnDataName,
			ColumnDataType:         magicTableColumnsMetadata.ColumnDataType, //api.MagicTableColumnDataType_Int,
			Sortable:               magicTableColumnsMetadata.Sortable,
			FormatPresentationType: magicTableColumnsMetadata.FormatPresentationType, //api.MagicTableColumnPresentationType_Simple,
			ShouldBeVisible:        magicTableColumnsMetadata.ShouldBeVisible,
		}

		magicTableColumnsMetadataResponse = append(magicTableColumnsMetadataResponse, magicTableColumnMetadata)

	}

	magicTableMetadataRespons = &api.MagicTableMetadataRespons{
		MagicTableColumnsMetadata: magicTableColumnsMetadataResponse,
		SearchbarIsVisible:        true,
	}

	return magicTableMetadataRespons, nil

	/*

		// Unique Id
		magicTableColumnMetadata := &api.MagicTableColumnMetadata{
			ColumnHeaderName:       "Unique Id",
			ColumnDataName:         "UniqueId",
			ColumnDataType:         api.MagicTableColumnDataType_Int,
			Sortable:               true,
			FormatPresentationType: api.MagicTableColumnPresentationType_Simple,
			ShouldBeVisible:        false,
		}
		magicTableColumnsMetadata = append(magicTableColumnsMetadata, magicTableColumnMetadata)

		// Name
		magicTableColumnMetadata = &api.MagicTableColumnMetadata{
			ColumnHeaderName:       "Name",
			ColumnDataName:         "Name",
			ColumnDataType:         api.MagicTableColumnDataType_String,
			Sortable:               true,
			FormatPresentationType: api.MagicTableColumnPresentationType_Simple,
			ShouldBeVisible:        true,
		}
		magicTableColumnsMetadata = append(magicTableColumnsMetadata, magicTableColumnMetadata)

		// Instance Type
		magicTableColumnMetadata = &api.MagicTableColumnMetadata{
			ColumnHeaderName:       "Instance Type",
			ColumnDataName:         "InstanceType",
			ColumnDataType:         api.MagicTableColumnDataType_String,
			Sortable:               false,
			FormatPresentationType: api.MagicTableColumnPresentationType_Simple,
			ShouldBeVisible:        true,
		}
		magicTableColumnsMetadata = append(magicTableColumnsMetadata, magicTableColumnMetadata)

		// ECU
		magicTableColumnMetadata = &api.MagicTableColumnMetadata{
			ColumnHeaderName:       "ECU",
			ColumnDataName:         "Ecu",
			ColumnDataType:         api.MagicTableColumnDataType_Float,
			Sortable:               true,
			FormatPresentationType: api.MagicTableColumnPresentationType_Simple,
			ShouldBeVisible:        true,
		}
		magicTableColumnsMetadata = append(magicTableColumnsMetadata, magicTableColumnMetadata)

		// Mem
		magicTableColumnMetadata = &api.MagicTableColumnMetadata{
			ColumnHeaderName:       "Mem",
			ColumnDataName:         "Memory",
			ColumnDataType:         api.MagicTableColumnDataType_Float,
			Sortable:               true,
			FormatPresentationType: api.MagicTableColumnPresentationType_Simple,
			ShouldBeVisible:        true,
		}
		magicTableColumnsMetadata = append(magicTableColumnsMetadata, magicTableColumnMetadata)

		// Network
		magicTableColumnMetadata = &api.MagicTableColumnMetadata{
			ColumnHeaderName:       "Network",
			ColumnDataName:         "Network",
			ColumnDataType:         api.MagicTableColumnDataType_String,
			Sortable:               true,
			FormatPresentationType: api.MagicTableColumnPresentationType_Simple,
			ShouldBeVisible:        true,
		}
		magicTableColumnsMetadata = append(magicTableColumnsMetadata, magicTableColumnMetadata)

		// Price
		magicTableColumnMetadata = &api.MagicTableColumnMetadata{
			ColumnHeaderName:       "Price",
			ColumnDataName:         "Price",
			ColumnDataType:         api.MagicTableColumnDataType_String,
			Sortable:               true,
			FormatPresentationType: api.MagicTableColumnPresentationType_Simple,
			ShouldBeVisible:        true,
		}
		magicTableColumnsMetadata = append(magicTableColumnsMetadata, magicTableColumnMetadata)

	*/

}
