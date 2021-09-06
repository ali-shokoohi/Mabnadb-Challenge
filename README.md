# mabnadp-challenge

My answer to the Mabnadp's challenge.
Written in Go, Gin, Gorm, PostgreSQL and etc...

## Documents

You can see all of API's documents in [docs/](./docs/) directory

## Configure

### Configure your Go:

If you didn't set your $GOPATH address It's okay.
Let's do this!
First, create a directory in your home.
Like this:

```bash
mkdir -p $HOME/go/
```

Then we must set this directory's address into $GOPATH.
All you have to do is just it:

```bash
export GOPATH=$HOME/go
```

Congratulations!!! You now configure your Go language.
Let's go next together...

### Configure the app:

Okay. Let me explain.
We have a file named .env-simple
Can you see that?
Great.
Let's look inside this
As you can see, we have some variables. Right?
Excellent! We call those variables the environments;
That can define in the whole system.
You must change those values as your values.
Like dbhost must be your Postgresql host address,
Or PORT must be a port that you like our app launch on this.
After changes my dear, It's time to rename this .env-simple file to just _.env_. Now our app can execute those environments.
Well done!!
Let's run our app...

## Run

### Run as docker way:

If you have docker on your machine, So you are a legend!
After configuring the app, For running it with docker we must enter a command:

```bash
docker-compose up --build -d
```

After a while, you can see you did it!
You can open your browser as localhost:_PORT_
The PORT is that value you defined in the _.env_ file,
Or 8090 as default.

### Run as classic way:

Oh, I see! We both like the classic way. Oh, wait! I don't mean that!
Never mind! Running our app is a better choice!
At first, We must install dependencies and build the app then we can run this.

#### Install and build:

To install dependencies and build our app as a single command we can enter this command:

```bash
go install
```

If You have some issues with installing dependencies, Maybe this command can help you:

```bash
go env -w GO111MODULE=auto
```

Very well! After that, we now have dependencies and a build file that you can't see here. Don't worry!
Our app has been compiled at the bin/ directory in your _$GOPATH_ directory.
The exact address of our binary file is:
_$GOPATH/bin/mabnadp-challenge_

#### Run the compiled file:

Well, After install and build our app we have a binary file,
To execute this file we only have to enter the address of the file in command line, Like the:

```bash
$GOPATH/bin/mabnadp-challenge
```

And there We go! Our app is now working very well!

#### Systemd

You can manage the app with systemd.
Oh fellow, Do We have a mabnadp-challenge.service file?
That's good!
Its contents are below:

```
[Unit]
Description=mabnadp-challenge

[Service]
Type=simple
Restart=always
RestartSec=5s
User=USER
WorkingDirectory=/home/USER/go/bin
ExecStart=/home/USER/go/bin/mabnadp-challenge
# sudo mkdir -p /var/log/mabnadp-challenge
StandardOutput=/var/log/mabnadp-challenge/good.log
StandardError=/var/log/mabnadp-challenge/bad.log

[Install]
WantedBy=multi-user.target
```

So You must change the USER value to your username.
After that, we must copy this file to systemd configures directory that systemd can run this.
Let's copy:

```
sudo cp mabnadp-challenge.service /etc/systemd/system/
```

Now we can manage the app!
Commands that you need:
For start:

```
sudo systemctl start mabnadp-challenge.service
```

For stop:

```
sudo systemctl stop mabnadp-challenge.service
```

For restart:

```
sudo systemctl restart mabnadp-challenge.service
```

For auto-start:

```
sudo systemctl enable --now mabnadp-challenge.service
```

That's it. We are done! We now have a rest API with Go!
I would be honored if you contribute to this project.
