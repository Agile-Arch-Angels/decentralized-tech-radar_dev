# Fetcher

The fetcher is used to pull whitelisted files or folders from one or more git repositories. The progress of the fetcher can be tracked using the small spinner and progress bar displayed in the terminal, where each dot represents a repository.

The fetcher takes a string containing 3 values:  
1. A URL to a git-based repository
2. A branch name
3. A path to a whitelist file
These can be repeated in the string as shown below for multiple fetches  



## Functions

* `func FetchFiles(url, branch, whitelist_file string) error`
     * **What it is**: This function is a public function that takes three arguments, a git repo url, the branch you want to pull from, and a path to a whitelist_file [See ... for info on the syntax for this](). It can return an error if an error is encountered during the fetch
    * **What it does**: This function will pull the files and folders specified in the whitelist file from the repo url using on the branch that is also given. 
* `func appendUrlToCSV(filename, url, repoName string)`
    * **What it is**: This is a private function that takes a filename to a fetched CSV file, a url being the repo that the CSV file was fetched from, and a repoName being the last part of the url aka. the name of the repository that is being fetched from
    * **What it does**: This function will append the repo url to all blips in a CSV file with the name of the filename argument using an `<a href=...>` HTML tag, the name of the tag is defined by the repoName argument. 
* `func ListingReposForFetch(repos []string) error`
    * **What it is**: A public function that takes a list of strings being a list of repositories, their branches and a whitelist file that is used on that repo
    * **What it does**: The function creates as many asynchronous fetch calls as there are repos in the list, it then creates a progress bar that reflects how many repos have been successfully processed
* `func round(x, unit float64) float64` 
    * **What it is**: A private function that takes two float64, an x and a unit.
    * **What it does**: The function rounds any float64 x to the nearest unit number so if unit = 0.05 it will round x to the nearest 0.05 number
* `func progressBar(numOfFiles int)`
    * **What it is**: A private function that takes an int numOfFiles which is the number of files that is expected to be fetched 
    * **What it does**: The function creates a 20 dot long progress bar that will update with hashtags for every nearest 5% finished fetches. So if you have finished 3/5 fetches, it would look like so [############........] 60%. It also includes a spinner to signal that he program is still running
* `func errHandler(err error, params ...string)`
    * **What it is**: A private function that takes in an error err and 0-infinite strings being relevant context for the error
    * **What it does**: This function is to reduce the amount of if err != nil statement. Instead you can just call this function with an error and it will panic for you with the relevant information.
* `func executer(cmd *exec.Cmd, folder string) error`
    * **What it is**: A private function that takes in a pointer to an exec.Cmd cmd and a string folder. It can return an error if the function fails at any point 
    * **What it does**: This function is supposed to help call terminal commands by letting the user construct a cmd and call this function with the wished working directory and it will run the command and return an error if it fails.
* `func puller(url, branch, whitelist_file string) ([]string, error)`
    * **What it is**: A private function that takes a git repo url, a branch for that repo, and path to a whitelist_file 
    * **What it does**: This function will call a chain of terminal commands to sparsely checkout and pull the files specified in the whitelist_file to a temp folder, it will the recursively walk through these folders if necessary to find CSV files and add them to a list to be cached later.

* ### Global variables for the file

## Important notes:

### Asynchronous
The fetcher runs using Go functions, meaning it runs asynchronously. This allows simultaneous fetching of files from multiple repositories, greatly decreasing the total running time on fetch calls with many repositories.

### Automatic Caching
The fetcher caches fetched CSV files in a folder named `cache`. The caching works for any folder depth so long as a CSV file is present

### Auto append Repository URLs
The fetcher will automatically append urls to the CSV files after verifying that they have no formatting errors. 