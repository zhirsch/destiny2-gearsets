package main

import (
	"flag"
	"fmt"
	"log"
	"sort"
	"strconv"

	"github.com/go-openapi/runtime"
	runtime_client "github.com/go-openapi/runtime/client"
	"github.com/zhirsch/destiny2-api/client"
	"github.com/zhirsch/destiny2-api/client/destiny2"
	"github.com/zhirsch/destiny2-api/client/user"
	"github.com/zhirsch/destiny2-api/models"
	db "github.com/zhirsch/destiny2-db"
)

var (
	flagAPIKey = flag.String("apikey", "", "the Bungie API key")
	flagUser   = flag.String("user", "", "the user to query")
)

func contains(haystack []uint32, needle uint32) bool {
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

func (is items) add(i *models.DestinyDefinitionsDestinyInventoryItemDefinition, owned bool) {
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
	other   items
}

type gearsets map[int32]*gearset

func (gs gearsets) add(g, i *models.DestinyDefinitionsDestinyInventoryItemDefinition, owned bool) {
	if _, ok := gs[int32(g.Hash)]; !ok {
		gs[int32(g.Hash)] = &gearset{
			hash:    int32(g.Hash),
			name:    g.DisplayProperties.Name,
			hunter:  make(items),
			titan:   make(items),
			warlock: make(items),
			weapons: make(items),
			other:   make(items),
		}
	}
	// TODO(zhirsch): Don't hard code these values.
	if contains(i.ItemCategoryHashes, 23) {
		gs[int32(g.Hash)].hunter.add(i, owned)
	} else if contains(i.ItemCategoryHashes, 22) {
		gs[int32(g.Hash)].titan.add(i, owned)
	} else if contains(i.ItemCategoryHashes, 21) {
		gs[int32(g.Hash)].warlock.add(i, owned)
	} else if contains(i.ItemCategoryHashes, 1) {
		gs[int32(g.Hash)].weapons.add(i, owned)
	} else {
		gs[int32(g.Hash)].other.add(i, owned)
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

func isError(errorCode models.ExceptionsPlatformErrorCodes) error {
	if errorCode != 1 {
		return fmt.Errorf("request failed: %v", errorCode)
	}
	return nil
}

func getCharacterItems(api *client.BungieNet, auth runtime.ClientAuthInfoWriter, destinyMembership *models.UserUserInfoCard) ([]uint32, error) {
	params := destiny2.NewDestiny2GetProfileParams()
	params.SetDestinyMembershipID(destinyMembership.MembershipID)
	params.SetMembershipType(int32(destinyMembership.MembershipType))
	params.SetComponents([]int64{100})
	resp, err := api.Destiny2.Destiny2GetProfile(params, auth)
	if err != nil {
		return nil, err
	}
	if err := isError(resp.Payload.ErrorCode); err != nil {
		return nil, err
	}
	var items []uint32
	for _, characterIDStr := range resp.Payload.Response.Profile.Data.CharacterIds {
		characterID, err := strconv.ParseInt(characterIDStr, 10, 64)
		if err != nil {
			return nil, err
		}
		params := destiny2.NewDestiny2GetCharacterParams()
		params.SetCharacterID(characterID)
		params.SetDestinyMembershipID(destinyMembership.MembershipID)
		params.SetMembershipType(int32(destinyMembership.MembershipType))
		params.SetComponents([]int64{201, 205})
		resp, err := api.Destiny2.Destiny2GetCharacter(params, auth)
		if err != nil {
			return nil, err
		}
		if err := isError(resp.Payload.ErrorCode); err != nil {
			return nil, err
		}
		for _, item := range resp.Payload.Response.Equipment.Data.Items {
			items = append(items, item.ItemHash)
		}
		if resp.Payload.Response.Inventory.Data != nil {
			for _, item := range resp.Payload.Response.Inventory.Data.Items {
				items = append(items, item.ItemHash)
			}
		}
	}
	return items, nil
}

func getVaultItems(api *client.BungieNet, auth runtime.ClientAuthInfoWriter, destinyMembership *models.UserUserInfoCard) ([]uint32, error) {
	params := destiny2.NewDestiny2GetProfileParams()
	params.SetDestinyMembershipID(destinyMembership.MembershipID)
	params.SetMembershipType(int32(destinyMembership.MembershipType))
	params.SetComponents([]int64{102})
	resp, err := api.Destiny2.Destiny2GetProfile(params, auth)
	if err != nil {
		return nil, err
	}
	if err := isError(resp.Payload.ErrorCode); err != nil {
		return nil, err
	}
	var items []uint32
	if resp.Payload.Response.ProfileInventory.Data != nil {
		for _, item := range resp.Payload.Response.ProfileInventory.Data.Items {
			items = append(items, item.ItemHash)
		}
	}
	return items, nil
}

func getItems(api *client.BungieNet, auth runtime.ClientAuthInfoWriter, destinyMembership *models.UserUserInfoCard) ([]uint32, error) {
	characterItems, err := getCharacterItems(api, auth, destinyMembership)
	if err != nil {
		return nil, err
	}
	vaultItems, err := getVaultItems(api, auth, destinyMembership)
	if err != nil {
		return nil, err
	}
	return append(characterItems, vaultItems...), nil
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

	// Create the authentication and API client.
	auth := runtime_client.APIKeyAuth("X-API-Key", "header", *flagAPIKey)
	api := client.Default

	// Open the manifest database.
	db, err := db.Open(api, auth)
	if err != nil {
		log.Fatal(err)
	}

	// Lookup the Bungie user.
	params := user.NewUserSearchUsersParams()
	params.SetQ(flagUser)
	resp, err := api.User.UserSearchUsers(params, auth)
	if err != nil {
		log.Fatal(err)
	}

	// Get all the items owned by the user.
	ownedItems := make(map[uint32]bool)
	for _, u := range resp.Payload.Response {
		params := user.NewUserGetMembershipDataByIDParams()
		params.SetMembershipID(u.MembershipID)
		params.SetMembershipType(-1)
		resp, err := api.User.UserGetMembershipDataByID(params, auth)
		if err != nil {
			log.Fatal(err)
		}
		for _, destinyMembership := range resp.Payload.Response.DestinyMemberships {
			items, err := getItems(api, auth, destinyMembership)
			if err != nil {
				log.Fatal(err)
			}
			for _, itemHash := range items {
				ownedItems[itemHash] = true
			}
		}
	}

	// Build the list of gearsets, and mark which items in them are owned by the user.
	gearsetDefinitions, err := db.GetAll("DestinyInventoryItemDefinition", &models.DestinyDefinitionsDestinyInventoryItemDefinition{})
	if err != nil {
		log.Fatal(err)
	}
	gs := make(gearsets)
	for _, gearsetDefinition := range gearsetDefinitions.([]*models.DestinyDefinitionsDestinyInventoryItemDefinition) {
		if gearsetDefinition.Gearset == nil {
			continue
		}
		for _, itemHash := range gearsetDefinition.Gearset.ItemList {
			itemDefinition, err := db.Get("DestinyInventoryItemDefinition", itemHash, &models.DestinyDefinitionsDestinyInventoryItemDefinition{})
			if err != nil {
				log.Fatal(err)
			}
			gs.add(gearsetDefinition, itemDefinition.(*models.DestinyDefinitionsDestinyInventoryItemDefinition), ownedItems[itemHash])
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
		printItems(g.other, "Other")
	}
}
