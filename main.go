package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/go-github/v52/github"
	"github.com/spf13/cobra"
)

var (
	projectName string
	orgnization string
	token       string
)

func main() {
	// Create a new root command for our CLI
	var rootCmd = &cobra.Command{
		Use:   "delete-github-project",
		Short: "Deletes a GitHub project",
		Run:   deleteProject,
	}

	// // Add a flag for the GitHub access token
	// rootCmd.PersistentFlags().StringVar(&token, "token", "", "GitHub access token")

	// Add a required argument for the GitHub project name
	rootCmd.Flags().StringVarP(&projectName, "name", "n", "", "GitHub project name (required)")
	err := rootCmd.MarkFlagRequired("name")
	if err != nil {
		log.Fatalf("Failed to mark 'name' flag as required: %v", err)
	}

	// Add a flag for the GitHub project orgnization
	rootCmd.Flags().StringVarP(&orgnization, "orgnization", "s", "vivsoftorg2", "GitHub project orgnization")

	// Execute the root command
	err = rootCmd.Execute()
	if err != nil {
		log.Fatalf("Failed to execute command: %v", err)
	}
}

func deleteProject(cmd *cobra.Command, args []string) {

	token = os.Getenv("GITHUB_TOKEN")
	if token == "" {
		log.Fatal("GitHub access token not provided")
	}

	ctx := context.Background()

	client := github.NewTokenClient(ctx, token)
	repo, _, err := client.Repositories.Get(ctx, orgnization, projectName)
	if err == nil {

		// Prompt the user for confirmation before deleting the project
		var confirm string
		fmt.Print("Do you want to Delete the Repository " + repo.GetName() + " with URL is " + repo.GetHTMLURL())
		_, err = fmt.Scanln(&confirm)
		if err != nil {
			log.Fatalf("Failed to read user input: %v", err)
		}
		if confirm != "yes" {
			fmt.Println("Project deletion aborted.")
			return
		}
		// Delete the project
		_, err := client.Repositories.Delete(ctx, orgnization, projectName)
		if err != nil {
			fmt.Print(err)
		}
	}

	if err != nil {
		fmt.Print(err)
	}

}

// func getAllProject(cmd *cobra.Command, args []string) {

// 	token = os.Getenv("GITHUB_TOKEN")
// 	if token == "" {
// 		log.Fatal("GitHub access token not provided")
// 	}

// 	ctx := context.Background()

// 	client := github.NewTokenClient(ctx, token)

// 	// list public repositories for org "github"
// 	opt := &github.RepositoryListByOrgOptions{Type: "private"}
// 	repos, _, err := client.Repositories.ListByOrg(ctx, orgnization, opt)
// 	if err != nil {
// 		fmt.Print(err)
// 	}
// 	for repo := range repos {
// 		fmt.Print("\n")
// 		fmt.Print(repos[repo].GetName())
// 	}

// }
