package main

import (
	"cloud.google.com/go/firestore"
	"context"
	"encoding/json"
	firebase "firebase.google.com/go"
	"github.com/urfave/cli/v2"
	"google.golang.org/api/option"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Commands = append(app.Commands,
		&cli.Command{
			Name:    "create",
			Aliases: []string{"c"},
			Flags: []cli.Flag{
				credentialsFile(),
				collection(),
				document(),
				jsonFile(),
			},
			Usage: "create a document on firebase",
			Action: func(c *cli.Context) error {
				return create(
					c.String("credentials_file"),
					c.String("collection"),
					c.String("document"),
					c.String("file"),
				)
			},
		},
		&cli.Command{
			Name:    "set",
			Aliases: []string{"s"},
			Flags: []cli.Flag{
				credentialsFile(),
				collection(),
				document(),
				jsonFile(),
			},
			Usage: "updates a document on firebase",
			Action: func(c *cli.Context) error {
				return set(
					c.String("credentials_file"),
					c.String("collection"),
					c.String("document"),
					c.String("file"),
				)
			},
		},
		&cli.Command{
			Name:    "delete",
			Aliases: []string{"d"},
			Flags: []cli.Flag{
				credentialsFile(),
				collection(),
				document(),
			},
			Usage: "deletes a document on firebase",
			Action: func(c *cli.Context) error {
				return delete(
					c.String("credentials_file"),
					c.String("collection"),
					c.String("document"),
				)
			},
		},
		&cli.Command{
			Name:    "read",
			Aliases: []string{"r"},
			Flags: []cli.Flag{
				credentialsFile(),
				collection(),
				document(),
				jsonFile(),
			},
			Usage: "read a document from firebase",
			Action: func(c *cli.Context) error {
				return read(
					c.String("credentials_file"),
					c.String("collection"),
					c.String("document"),
					c.String("file"),
				)
			},
		})
	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func create(credentialsFile, collection, documentName, fileName string) error {
	ctx := context.Background()
	cli, err := new(ctx, credentialsFile)
	if err != nil {
		return err
	}
	content, err := readJSONFile(fileName)
	if err != nil {
		return err
	}
	col := cli.Collection(collection)
	cal := col.Doc(documentName)
	_, err = cal.Create(ctx, content)
	return err
}

func delete(credentialsFile, collection, documentName string) error {
	ctx := context.Background()
	cli, err := new(ctx, credentialsFile)
	if err != nil {
		return err
	}
	col := cli.Collection(collection)
	cal := col.Doc(documentName)
	_, err = cal.Delete(ctx)
	return err
}

func set(credentialsFile, collection, documentName, file string) error {
	ctx := context.Background()
	cli, err := new(ctx, credentialsFile)
	if err != nil {
		return err
	}
	content, err := readJSONFile(file)
	if err != nil {
		return err
	}
	col := cli.Collection(collection)
	cal := col.Doc(documentName)
	_, err = cal.Set(ctx, content)
	return err
}

func read(credentialsFile, collection, documentName, fileName string) error {
	ctx := context.Background()
	cli, err := new(ctx, credentialsFile)
	if err != nil {
		return err
	}
	col := cli.Collection(collection)
	doc := col.Doc(documentName)
	snap, err := doc.Get(context.Background())
	if err != nil {
		return err
	}
	res, err := json.Marshal(snap.Data())
	if err != nil {
		return err
	}
	return ioutil.WriteFile(fileName, res, 0644)
}

func readJSONFile(file string) (map[string]interface{}, error) {
	body, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	out := map[string]interface{}{}
	err = json.Unmarshal(body, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func new(ctx context.Context, file string) (*firestore.Client, error) {
	opt := option.WithCredentialsFile(file)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, err
	}
	client, err := app.Firestore(ctx)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func collection() cli.Flag {
	return &cli.StringFlag{
		Name:     "collection",
		Aliases:  []string{"c"},
		Usage:    "--" + "collection" + " name",
		EnvVars:  []string{"COLLECTION"},
		Required: true,
	}
}

func credentialsFile() cli.Flag {
	return &cli.StringFlag{
		Name:     "authentication",
		Aliases:  []string{"auth"},
		Usage:    "--" + "authentication" + " name",
		EnvVars:  []string{"AUTHENTICATION"},
		Required: true,
	}
}

func document() cli.Flag {
	return &cli.StringFlag{
		Name:     "document",
		Aliases:  []string{"d"},
		Usage:    "--" + "document" + " name",
		EnvVars:  []string{"DOCUMENT"},
		Required: true,
	}
}

func jsonFile() cli.Flag {
	return &cli.StringFlag{
		Name:     "file",
		Aliases:  []string{"f"},
		Usage:    "--" + "file" + " name",
		Required: true,
	}
}
