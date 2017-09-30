package main

import (
	"flag"
	"fmt"
	"log"
	"sort"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
	api "github.com/zhirsch/destiny2-api"
	db "github.com/zhirsch/destiny2-db"
)

var (
	apiKey = flag.String("apikey", "", "the Bungie API key")
	user   = flag.String("user", "", "the user to query")
)

func containsInt32(haystack []int32, needle int32) bool {
	for _, x := range haystack {
		if x == needle {
			return true
		}
	}
	return false
}

type item struct {
	hash        int32
	name        string
	typeAndTier string
	owned       bool
}

type items map[int32]*item

func (is items) add(i *api.DestinyDefinitionsDestinyInventoryItemDefinition, owned bool) {
	is[int32(i.Hash)] = &item{
		hash:        int32(i.Hash),
		name:        i.DisplayProperties.Name,
		typeAndTier: i.ItemTypeAndTierDisplayName,
		owned:       owned,
	}
}

type byItemName []*item

func (b byItemName) Len() int {
	return len(b)
}

func (b byItemName) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b byItemName) Less(i, j int) bool {
	if b[i].name == b[j].name {
		return b[i].hash < b[j].hash
	}
	return b[i].name < b[j].name
}

func (is items) sortedKeys() []int32 {
	var entries []*item
	for _, v := range is {
		entries = append(entries, v)
	}
	sort.Sort(byItemName(entries))

	var keys []int32
	for _, e := range entries {
		keys = append(keys, e.hash)
	}
	return keys
}

type gearset struct {
	hash    int32
	name    string
	hunter  items
	titan   items
	warlock items
	weapons items
}

type gearsets map[int32]*gearset

func (gs gearsets) add(g, i *api.DestinyDefinitionsDestinyInventoryItemDefinition, owned bool) {
	if _, ok := gs[int32(g.Hash)]; !ok {
		gs[int32(g.Hash)] = &gearset{
			hash:    int32(g.Hash),
			name:    g.DisplayProperties.Name,
			hunter:  make(items),
			titan:   make(items),
			warlock: make(items),
			weapons: make(items),
		}
	}
	// TODO(zhirsch): Don't hard code these values.
	if containsInt32(i.ItemCategoryHashes, 23) {
		gs[int32(g.Hash)].hunter.add(i, owned)
	} else if containsInt32(i.ItemCategoryHashes, 22) {
		gs[int32(g.Hash)].titan.add(i, owned)
	} else if containsInt32(i.ItemCategoryHashes, 21) {
		gs[int32(g.Hash)].warlock.add(i, owned)
	} else if containsInt32(i.ItemCategoryHashes, 1) {
		gs[int32(g.Hash)].weapons.add(i, owned)
	}
}

type byGearsetName []*gearset

func (b byGearsetName) Len() int {
	return len(b)
}

func (b byGearsetName) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b byGearsetName) Less(i, j int) bool {
	if b[i].name == b[j].name {
		return b[i].hash < b[j].hash
	}
	return b[i].name < b[j].name
}

func (gs gearsets) sortedKeys() []int32 {
	var entries []*gearset
	for _, v := range gs {
		entries = append(entries, v)
	}
	sort.Sort(byGearsetName(entries))

	var keys []int32
	for _, e := range entries {
		keys = append(keys, e.hash)
	}
	return keys
}

func parseCharacterID(s string) int64 {
	if id, err := strconv.ParseInt(s, 10, 64); err != nil {
		panic(err)
	} else {
		return id
	}
}

func getCharacterItems(client *api.APIClient, destinyMembership api.UserUserInfoCard) []uint32 {
	args := map[string]interface{}{"components": []api.DestinyDestinyComponentType{100}}
	response, _, err := client.Destiny2Api.Destiny2GetProfile(destinyMembership.MembershipId, destinyMembership.MembershipType, args)
	if err != nil {
		panic(err)
	}
	var items []uint32
	for _, characterIDStr := range response.Response.Profile.Data.CharacterIds {
		characterID := parseCharacterID(characterIDStr)
		args := map[string]interface{}{"components": []api.DestinyDestinyComponentType{201, 205}}
		response, _, err := client.Destiny2Api.Destiny2GetCharacter(characterID, destinyMembership.MembershipId, destinyMembership.MembershipType, args)
		if err != nil {
			panic(err)
		}
		for _, item := range response.Response.Equipment.Data.Items {
			items = append(items, item.ItemHash)
		}
		if response.Response.Inventory.Data != nil {
			for _, item := range response.Response.Inventory.Data.Items {
				items = append(items, item.ItemHash)
			}
		}
	}
	return items
}

func getVaultItems(client *api.APIClient, destinyMembership api.UserUserInfoCard) []uint32 {
	args := map[string]interface{}{"components": []api.DestinyDestinyComponentType{102}}
	response, _, err := client.Destiny2Api.Destiny2GetProfile(destinyMembership.MembershipId, destinyMembership.MembershipType, args)
	if err != nil {
		panic(err)
	}
	var items []uint32
	if response.Response.ProfileInventory.Data != nil {
		for _, item := range response.Response.ProfileInventory.Data.Items {
			items = append(items, item.ItemHash)
		}
	}
	return items
}

func getItems(client *api.APIClient, destinyMembership api.UserUserInfoCard) []uint32 {
	items := getCharacterItems(client, destinyMembership)
	items = append(items, getVaultItems(client, destinyMembership)...)
	return items
}

func printItems(items items, name string) {
	if len(items) == 0 {
		return
	}
	fmt.Printf("  %v\n", name)
	for _, ik := range items.sortedKeys() {
		i := items[ik]
		owned := " "
		if i.owned {
			owned = "✔"
		}
		fmt.Printf("   %v %v (%v) [%v]\n", owned, i.name, i.typeAndTier, i.hash)
	}
}

func main() {
	flag.Parse()

	// Create an API client.
	cfg := api.NewConfiguration()
	cfg.AddDefaultHeader("X-API-Key", *apiKey)
	client := api.NewAPIClient(cfg)

	// Open the manifest database.
	db, err := db.Open(client)
	if err != nil {
		log.Fatal(err)
	}

	// Lookup the Bungie user.
	args := map[string]interface{}{"q": *user}
	response, _, err := client.UserApi.UserSearchUsers(args)
	if err != nil {
		panic(err)
	}

	// Get all the items owned by the user.
	ownedItems := make(map[uint32]bool)
	for _, user := range response.Response {
		resp, _, err := client.UserApi.UserGetMembershipDataById(user.MembershipId, -1)
		if err != nil {
			panic(err)
		}
		for _, destinyMembership := range resp.Response.DestinyMemberships {
			for _, itemHash := range getItems(client, destinyMembership) {
				ownedItems[itemHash] = true
			}
		}
	}

	// Build the list of gearsets, and mark which items in them are owned by the user.
	gearsetDefinitions, err := db.GetAll("DestinyInventoryItemDefinition", &api.DestinyDefinitionsDestinyInventoryItemDefinition{})
	if err != nil {
		log.Fatal(err)
	}
	gs := make(gearsets)
	for _, gearsetDefinition := range gearsetDefinitions.([]*api.DestinyDefinitionsDestinyInventoryItemDefinition) {
		if gearsetDefinition.Gearset == nil {
			continue
		}
		for _, itemHash := range gearsetDefinition.Gearset.ItemList {
			itemDefinition, err := db.Get("DestinyInventoryItemDefinition", itemHash, &api.DestinyDefinitionsDestinyInventoryItemDefinition{})
			if err != nil {
				log.Fatal(err)
			}
			gs.add(gearsetDefinition, itemDefinition.(*api.DestinyDefinitionsDestinyInventoryItemDefinition), ownedItems[itemHash])
		}
	}

	// Print out the results.
	for _, gk := range gs.sortedKeys() {
		g := gs[gk]
		fmt.Printf("%v [%v]\n", g.name, g.hash)
		printItems(g.hunter, "Hunter Armor")
		printItems(g.titan, "Titan Armor")
		printItems(g.warlock, "Warlock Armor")
		printItems(g.weapons, "Weapons")
	}
}
