# BEEGO
This is an example beego web app, to demonstrate how quick and easy is to develop a web applicaiton in beego.

## Prerequisites
* Go
* Bee tool
You can install the bee tool using `go install github.com/beego/bee/v2@latest`.

Once the bee tool is installed you can type `bee` to see all the available options that bee tool provides.

## To run this app
* Update the configuration file in `conf/app.conf`.
	* You can set the http port on which our server will be running, by default it is 8080.
	* We are using Mysql database. You will have to specify `dbuser`, `dbpass`, `dbhost`, `dbport`,`dbname` in the conf file.
* Once the confifurations are set you have to do `bee run`, which will build our Beego application and will the execute the generated binary. You should see http server running on port 8080. You can visit `localhost:8080` to see our app.
* Admin dashboard is running on  `localhost:8088` to monitor our application.

