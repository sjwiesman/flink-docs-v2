# Apache Flink Docs V2

This is an experimental rewrite of Flink docs using Hugo.
The end goal is to replace Flinks Jekyll infrastructure with a more modern system. 

### Build the site locally

Make sure you have installed
[Hugo](https://gohugo.io/getting-started/installing/) on your
system. Follow the instructions to clone this repository and build the
docs locally.

  * Clone the repository
	```sh
	git clone https://github.com/sjwiesman/flink-docs-v2/
	cd flink-docs-v2/
	```
  * Fetch the theme submodule
	```sh
	git submodule update --init --recursive
	```
  * Start local server
	```sh
	hugo serve
	```
	Site can be viewed at http://localhost:1313
