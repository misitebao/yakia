package commands

import (
	"fmt"

	ini18n "github.com/misitebao/standard-repository/cmd/stdrepo/internal/i18n"
)

var helpTemplate = `
{{with (or .Long .Short)}}{{. | trimTrailingWhitespaces}}{{end}}

%s:{{if .Runnable}}
  {{.UseLine}}{{end}}{{if .HasAvailableSubCommands}}{{end}} {{if gt (len .Aliases) 0}}

%s:
  {{.NameAndAliases}}{{end}}{{if .HasExample}}

%s:
{{.Example}}{{end}}{{if .HasAvailableSubCommands}}

%s:{{range .Commands}}{{if (or .IsAvailableCommand)}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableLocalFlags}}

%s:
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasAvailableInheritedFlags}}

%s:
{{.InheritedFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasHelpSubCommands}}

Additional help topics:{{range .Commands}}{{if .IsAdditionalHelpTopicCommand}}
  {{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableSubCommands}}	

Use "{{.CommandPath}} [command] --help" for more information about a command.{{end}}

If Stdrepo is useful to you or your company, please consider sponsoring the project:
https://github.com/sponsors/misitebao
`

func init() {
	cmd.SetHelpTemplate(fmt.Sprintf(helpTemplate,
		ini18n.T("Template.Usage"),
		ini18n.T("Template.Aliases"),
		ini18n.T("Template.Examples"),
		ini18n.T("Template.AvailableCommands"),
		ini18n.T("Template.Flags"),
		ini18n.T("Template.GlobalFlags"),
	))

}
