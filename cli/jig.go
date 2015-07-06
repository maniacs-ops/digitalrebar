package main

import (
	"encoding/json"
	"fmt"
	"log"

	crowbar "github.com/VictorLowther/crowbar-api"
	"github.com/spf13/cobra"
)

func addJiggerCommands(singularName string,
	maker func() crowbar.Crudder,
	res *cobra.Command) {
	if _, ok := maker().(crowbar.Jigger); !ok {
		return
	}
	cmd := &cobra.Command{
		Use:   "jigs [id]",
		Short: fmt.Sprintf("List all jigs for a specifc %v", singularName),
		Run: func(c *cobra.Command, args []string) {
			if len(args) != 1 {
				log.Fatalf("%v requires 1 argument\n", c.UseLine())
			}
			obj := maker().(crowbar.Jigger)
			if crowbar.SetId(obj, args[0]) != nil {
				log.Fatalf("Failed to parse ID %v for an %v\n", args[0], singularName)
			}
			objs, err := crowbar.Jigs(obj)
			if err != nil {
				log.Fatalf("Failed to get jigss for %v(%v)\n", singularName, args[0])
			}
			fmt.Println(prettyJSON(objs))
		},
	}
	res.AddCommand(cmd)
}

func init() {
	lister := func() ([]crowbar.Crudder, error) {
		jigs, err := crowbar.Jigs()
		if err != nil {
			return nil, err
		}
		res := make([]crowbar.Crudder, len(jigs))
		for i := range jigs {
			res[i] = jigs[i]
		}
		return res, nil
	}
	matcher := func(sample string) (string, error) {
		obj := &crowbar.Jig{}
		err := json.Unmarshal([]byte(sample), obj)
		if err != nil {
			return "", fmt.Errorf("Error unmarshalling jig\nError: %v\n", err.Error())
		}
		objs, err := obj.Match()
		if err != nil {
			return "", fmt.Errorf("Error fetching matches for %v", sample)
		}
		return prettyJSON(objs), nil
	}
	maker := func() crowbar.Crudder { return &crowbar.Jig{} }
	singularName := "jig"
	app.AddCommand(makeCommandTree(singularName, lister, matcher, maker))
}
