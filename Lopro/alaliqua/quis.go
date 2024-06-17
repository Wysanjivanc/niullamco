import (
	"context"
	"fmt"
	"io"

	datacatalog "cloud.google.com/go/datacatalog/apiv1"
	"cloud.google.com/go/datacatalog/apiv1/datacatalogpb"
)

// addFarm demonstrates creating an entry group and entry for a farm.
func addFarm(w io.Writer, projectID string) error {
	// projectID := "my-project-id"
	location := "us"
	entryGroupID := "farms"
	entryID := "my_farm"
	ctx := context.Background()
	policyTag := "my-data-access-policy"
	policyTag2 := "my-data-access-policy-2"
	policyTag3 := "my-data-access-policy-3"

	client, err := datacatalog.NewClient(ctx)
	if err != nil {
		return fmt.Errorf("datacatalog.NewClient: %v", err)
	}
	defer client.Close()

	req := &datacatalogpb.CreateEntryGroupRequest{
		Parent: fmt.Sprintf("projects/%s/locations/%s", projectID, location),
		EntryGroupId: entryGroupID,
		EntryGroup: &datacatalogpb.EntryGroup{
			DisplayName: "My Farms",
			Description: "A group of farms.",
		},
	}
	entryGroup, err := client.CreateEntryGroup(ctx, req)
	if err != nil {
		return fmt.Errorf("CreateEntryGroup: %v", err)
	}
	fmt.Fprintf(w, "Created entry group: %s\n", entryGroup.Name)

	req = &datacatalogpb.CreateEntryRequest{
		Parent: entryGroup.Name,
		EntryId: entryID,
		Entry: &datacatalogpb.Entry{
			DisplayName: "My Farm",
			Description: "A farm in Iowa.",
			Type_:       "Farm",
			Schema: &datacatalogpb.Schema{
				Columns: []*datacatalogpb.ColumnSchema{
					{
						Column: "name",
						Type_:   "STRING",
						Mode:    "REQUIRED",
						Description: "The name of the farm.",
					},
					{
						Column: "acres",
						Type_:   "DOUBLE",
						Mode:    "REQUIRED",
						Description: "The number of acres in the farm.",
					},
					{
						Column: "owner",
						Type_:   "STRING",
						Mode:    "NULLABLE",
						Description: "The name of the owner of the farm.",
					},
				},
			},
			SourceSystemTimestamps: &datacatalogpb.SystemTimestamps{
				CreateTime: &datacatalogpb.Timestamp{
					Seconds: 1585241600,
				},
				UpdateTime: &datacatalogpb.Timestamp{
					Seconds: 1585241600,
				},
			},
			UserSpecifiedSystem: map[string]string{
				"custom_label": "custom_value",
			},
			UserSpecifiedLabels: map[string]string{
				"label_a": "value_a",
				"label_b": "value_b",
				"label_c": "value_c",
			},
			PolicyTags: []string{
				policyTag,
				policyTag2,
				policyTag3,
			},
		},
	}
	entry, err := client.CreateEntry(ctx, req)
	if err != nil {
		return fmt.Errorf("CreateEntry: %v", err)
	}
	fmt.Fprintf(w, "Created entry: %s\n", entry.Name)
	return nil
}
  
