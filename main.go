package main

import (
	"encoding/json"

	"github.com/CodeGeek04/Nextjs-OpenAPI-Converter/handler"
	"github.com/CodeGeek04/Nextjs-OpenAPI-Converter/types"
	"github.com/CodeGeek04/Nextjs-OpenAPI-Converter/utils"
)

func main() {
	item := handler.HandleFolder("C:/Users/Shivam Mittal/Desktop/IgniteTech/ignitetech-eloquens-ai/web/app/api", "C:/Users/Shivam Mittal/Desktop/IgniteTech/ignitetech-eloquens-ai/web/app/")
	postmanCollection := types.PostmanCollection{
		Info: types.PostmanInfo{
			Name:   "Nextjs to API Collection",
			Schema: "https://schema.getpostman.com/json/collection/v2.1.0/collection.json",
		},
		Item: []types.Item{item},
	}

	indentedMarshal, err := json.MarshalIndent(postmanCollection, "", "    ")
	utils.CheckError(err)

	err = utils.WriteToFile(indentedMarshal, "PostmanSchema.json")
	utils.CheckError(err)
}
