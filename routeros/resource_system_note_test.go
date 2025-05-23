package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// console - added "show-at-cli-login" option to display a note before telnet login;
const testSystemNoteTaskMinVersion = "7.14"
const testSystemNoteTask = "routeros_system_note.test"

const testSystemNoteNote = "For authorized use only."

func TestAccSystemNoteTest_basic(t *testing.T) {
	if !testCheckMinVersion(t, testSystemNoteTaskMinVersion) {
		t.Logf("Test skipped, the minimum required version is %v", testSystemNoteTaskMinVersion)
		return
	}

	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				Steps: []resource.TestStep{
					{
						Config: testAccSystemNoteConfig(testSystemNoteNote, false, false),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testSystemNoteTask),
							resource.TestCheckResourceAttr(testSystemNoteTask, "note", testSystemNoteNote),
							resource.TestCheckResourceAttr(testSystemNoteTask, "show_at_login", "false"),
							resource.TestCheckResourceAttr(testSystemNoteTask, "show_at_cli_login", "false"),
						),
					},
					{
						Config: testAccSystemNoteConfig(testSystemNoteNote, true, false),
						Check: resource.ComposeTestCheckFunc(
							resource.TestCheckResourceAttr(testSystemNoteTask, "note", testSystemNoteNote),
							resource.TestCheckResourceAttr(testSystemNoteTask, "show_at_login", "true"),
							resource.TestCheckResourceAttr(testSystemNoteTask, "show_at_cli_login", "false"),
						),
					},
					{
						Config: testAccSystemNoteConfig(testSystemNoteNote, false, true),
						Check: resource.ComposeTestCheckFunc(
							resource.TestCheckResourceAttr(testSystemNoteTask, "note", testSystemNoteNote),
							resource.TestCheckResourceAttr(testSystemNoteTask, "show_at_login", "false"),
							resource.TestCheckResourceAttr(testSystemNoteTask, "show_at_cli_login", "true"),
						),
					},
					{
						Config: testAccSystemNoteConfig(testSystemNoteNote, true, true),
						Check: resource.ComposeTestCheckFunc(
							resource.TestCheckResourceAttr(testSystemNoteTask, "note", testSystemNoteNote),
							resource.TestCheckResourceAttr(testSystemNoteTask, "show_at_login", "true"),
							resource.TestCheckResourceAttr(testSystemNoteTask, "show_at_cli_login", "true"),
						),
					},
				},
			})

		})
	}
}

func testAccSystemNoteConfig(note string, showAtLogin bool, showAtCliLogin bool) string {
	return fmt.Sprintf(`%v

resource "routeros_system_note" "test" {
	note              = "%v"
	show_at_login     = %v
	show_at_cli_login = %v
}
`, providerConfig, note, showAtLogin, showAtCliLogin)
}
