Go commandline shorthand utility for Linux
------------------------------------------

- Run commands added to a data.json and refer to them by means of the name assigned to it.

- Can be used to run rudimentary batch commands in sequence, as specified, which otherwise need a bash script and setting up the environment.

- Purpose is to reduce time spent on running commands that are used often.

Usage
-----

- Pull module via go get.

- Run go build to compile. Setup the binary in a directory included in PATH.

- Add GOPLAY environment variable to point to the data.json file in ~/.bashrc or ~/.bash_profile.
