# Golang CLI Task!

The app is a command line app written in go using Cobra package.The app searches available Github repositories based on a given string.

## Building from source

    - go build -o app main.go

## Search Command

The search command accepts a search string to query Github API.

### Example:

    - app search go

Running this command will fetch a limited number of repositories that it s name contains the given param which is "go".

The search command has 3 optional flags which are:

- Sort flag: sort parameter will sort results by repository name in Ascending or Descending order.
- Ignore flag: ignore parameter will ignore repositories, where the name of the repository includes the provided string.
- Page flag: page parameter will set the page number of the results to fetch.

### Example

    -  app search go -s desc -i master -p 3

Running this command will fetch the third page of the repositories that has "go" in the name of the repository in a descending order where the name of the repository doesn't contain the word "master" .

For more information about the flags run app search --help

## Run application using docker

### Build Image:

Run the following command:

    docker build . -t cli

### Run Image:

The application logs sent request in a directory inside the image, the image shares a volume on where you can get the log file.
To run the app run the following command:

    docker run --volume=$HOME/logs:/logs -i cli help

This will run the CLI app and show help menu.
