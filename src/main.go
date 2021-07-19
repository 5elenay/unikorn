package main

// import "fmt"

var currentVersion string = "0.2.0"

func main() {
	// Handle Commands
	HandleCommands()

	// Generates command documentation for me, Because i am too lazy to write myself.
	/*
		for _, item := range allCommands {
			fmt.Printf("## %s\n%s\n", item.Name, item.Description)
			fmt.Print("### Example(s)\n")

			for _, usage := range item.Usage {
				fmt.Printf("```bash\n%s\n```\n", usage)
			}

			if len(item.Options) > 0 {
				fmt.Print("### Options\n")

				for _, option := range item.Options {
					fmt.Printf("- `%s`: %s\n", option.Name, option.Description)
				}
			}
		}
	*/
}
