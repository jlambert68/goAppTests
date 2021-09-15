package server

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"goAppTest1/cmd/070_AddListDelete_TestDomains/protos/api"
)

func (server *Server) GetMagicTableMetadata(ctx context.Context, in *api.MagicTableMetadataRequest) (*api.MagicTableMetadataRespons, error) {

	var magicTableMetadataRespons *api.MagicTableMetadataRespons

	var magicTableColumnsMetadataResponse []*api.MagicTableColumnMetadata
	var err error

	// Initiate logger if not already done
	server.checkIfLoggerIsInitiated()

	server.logger.WithFields(logrus.Fields{
		"Id":    "becb4612-2f92-433d-9553-b64eae6af3f5",
		"Trace": server.trace(false),
	}).Debug("Entering: GetMagicTableMetadata()")

	defer func() {
		server.logger.WithFields(logrus.Fields{
			"Id":    "278a5a9c-f3f5-4a42-a6ba-3de26dc6ce9f",
			"Trace": server.trace(false),
		}).Debug("Exiting: GetMagicTableMetadata()")
	}()

	magicTableColumnsMetadataFromDB, err := server.listMagicTableColumnMetadataForjsonFile(in.TableTypeGuid)
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
