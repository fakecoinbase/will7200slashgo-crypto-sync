package main

import (
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/alecthomas/kong"
	"github.com/pelletier/go-toml"
	"github.com/pelletier/go-toml/query"

	_ "github.com/will7200/go-crypto-sync/internal/coinbase"
	_ "github.com/will7200/go-crypto-sync/internal/holdings"
)

type Context struct {
	Debug  bool
	Tree   *toml.Tree
	Config Config
}

var cli struct {
	tree  *toml.Tree
	Debug bool `help:"Enable debug mode."`

	Sync SyncCmd `cmd help:"Sync holdings to another account"`
}

// Config holds details when syncing
type Config struct {
	// Debug
	// Prints Debug Information
	Debug bool `toml:"debug"`
	// Destination will the aggregated information will go
	// Supported: person capital
	Destination string `toml:"destination"`
	// PriceDataSource will fetch the currency pricing data from this
	// Supported: coinbase
	PriceDataSource string `toml:"priceDataSource"`
	// DestinationCurrencyAs will fetch the converted pricing data of the concurrency in the specified format
	// Data Matrix:
	// Coinbase: USD, many others look at their api
	DestinationCurrencyAs string `toml:"destinationCurrencyAs"`
	// List of crypto currency holdings with their configurations
	// Supported: coinbase
	Holdings map[string]map[string]interface{} `toml:"holdings"`
	// List of Destinations holding their configuration
	// Supported: personalcapital
	Destinations map[string]map[string]interface{} `toml:"destinations"`
}

// TOML returns a Resolver that retrieves values from a TOML source.
//
// Hyphens in flag names are replaced with underscores.
func TOML(r io.Reader) (kong.Resolver, error) {
	config, err := toml.LoadReader(r)
	if err != nil {
		return nil, err
	}
	cli.tree = config
	var f kong.ResolverFunc = func(context *kong.Context, parent *kong.Path, flag *kong.Flag) (interface{}, error) {
		name := strings.Replace(flag.Name, "-", "_", -1)
		raw, err := query.CompileAndExecute(fmt.Sprintf("$.%s", name), config)
		if err != nil || len(raw.Values()) == 0 {
			return nil, nil
		}
		values := raw.Values()
		if flag.IsBool() {
			return values[0], nil
		}
		if flag.IsSlice() {
			return raw.Values(), nil
		}
		return raw.Values()[0], nil
	}

	return f, nil
}

func main() {
	ctx := kong.Parse(&cli, kong.Configuration(TOML, "config.toml"))
	log.SetFlags(log.Lshortfile | log.Ldate)
	conf := Config{
		Holdings:     map[string]map[string]interface{}{},
		Destinations: map[string]map[string]interface{}{},
	}
	err := cli.tree.Unmarshal(&conf)
	if err != nil {
		panic(err)
	}
	err = ctx.Run(&Context{Debug: cli.Debug, Tree: cli.tree, Config: conf})
	ctx.FatalIfErrorf(err)
}
