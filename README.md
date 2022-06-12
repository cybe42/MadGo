# MadGo
![MadGo](https://cdn.discordapp.com/attachments/890628633981513729/899767979783376926/robot.png)

MadGo is a powerful boilerplate for you web applications providing fast prototyping and ease of deployability using [Docker](https://github.com/docker)

- **Still W.I.P**

Welcome to MadGo! MadGo gives you a template to base your web applications on with powerful tools.


First, Set up everything.

Requirements:
- Go
- Git
- Make
- NPM/PNPM

Run the following commands to any directory you wish to download the template to:
```
git clone http://github.com/cybe42/MadGo.git
```

after it is complete, set up everything by running the setup script for your OS of choice.

- On Linux:
```
chmod +x setup.sh
./setup.sh
```

- On Windows:

```
setup.bat
```

And after you set up, you can edit the folder `api` to write you API's and add it to `registerapi.go`
You can edit the frontend folder. If you don't know how, take a look at [SolidJS](https://www.solidjs.com/guides/getting-started#learn-solid)



and to run your app enter the following:
```
make run
```

and if you want to build your app
```
make build
```

to build with debug information stripped:
```
make release
```