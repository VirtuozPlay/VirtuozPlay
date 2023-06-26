# Welcome to VirtuozPlay

## Development setup

### Dependencies

You need to have the following dependencies installed on your system:

- [Go](https://go.dev/doc/install)
- [Node.js](https://nodejs.org/en/download/)
- [Yarn](https://classic.yarnpkg.com/en/docs/install)
- [PostgreSQL](https://www.postgresql.org/download/)
- [Buffalo](https://gobuffalo.io/en/docs/installation)

#### ArchLinux

```sh
# Install packages
pacman -S go npm nodejs yarn

# Install buffalo from AUR /!\ version 0.18.5
yay -S buffalo-bin
# OR paru -S buffalo-bin

# install js things
yarn install

# Install plugins
buffalo plugins install
```

#### MacOS

**_TODO_**

#### Nix

1. Install and setup [Direnv](https://direnv.net/)
2. Run `direnv allow $PROJECT_DIR`, where `$PROJECT_DIR` is the directory where you cloned this repository
3. Run `cd $PROJECT_DIR`
4. It should automatically install all the dependencies, wait for it to finish

#### Windows

1. install Go 1.20
   https://go.dev/dl/

2. Install Postgres
   https://www.postgresql.org/download/windows/
   open pgAdmin

3. Install a package manager
   run powershell in administrator
   follow the instructions -> https://chocolatey.org/install
   run `choco`
   
4. change password in database.yml
   password: ...

5. install Buffalo + plugins
   ```cmd
   choco install buffalo
   buffalo plugins install
   ```

6. PATH
   put the yarn path in environment variable

7. create db
   `buffalo pop create -a`

### Recommended IDE Plugins

#### VSCode

- [Go](https://code.visualstudio.com/docs/languages/go)
- [ESLint](https://marketplace.visualstudio.com/items?itemName=dbaeumer.vscode-eslint)
- [Prettier](https://marketplace.visualstudio.com/items?itemName=esbenp.prettier-vscode)
- [Material Icon Theme](https://marketplace.visualstudio.com/items?itemName=PKief.material-icon-theme)
- [Volar](https://marketplace.visualstudio.com/items?itemName=johnsoncodehk.volar) (Vue.js 3 support)

#### GoLand / IntelliJ

- [Go plugin](https://plugins.jetbrains.com/plugin/9568-go) (not needed for GoLand)
- [Atom Material Icons](https://plugins.jetbrains.com/plugin/10044-atom-material-icons)
- [Prettier](https://plugins.jetbrains.com/plugin/10456-prettier)
- [Tailwind CSS](https://plugins.jetbrains.com/plugin/12075-tailwind-css)
- [Vue.js](https://plugins.jetbrains.com/plugin/9442-vue-js)
- [Vite](https://plugins.jetbrains.com/plugin/14632-vite)
  
### Database Setup

You may need to do is open up the "database.yml" file and edit it to use the correct usernames, passwords, hosts, etc... that are appropriate for your environment.
You will also need to make sure that **you** start/install the Postgres. Buffalo **won't** install and start it for you.

#### Create Your Databases

Ok, so you've edited the "database.yml" file and started your database, now Buffalo can create the databases in that file for you:

```sh
buffalo pop create -a
```

### Starting the Application

Buffalo ships with a command that will watch your application and automatically rebuild the Go binary and any assets for you. To do that run the "buffalo dev" command:

If you point your browser to [http://localhost:5173](http://localhost:5173) you should see a "You did it!" page.

#### Using Nix

```sh
# Automatically starts the database
devenv up
# Run once to create the base directories
buffalo build
# Make sure the database is running
buffalo dev
```

#### Using Docker

```sh
# If using docker for database
docker compose up -d
# Run once to create the base directories
buffalo build
# Make sure the database is running
buffalo dev
```

### Useful Commands

```shell
# Runs ESLint against the project
yarn lint
# Runs Prettier against the project
yarn format
# Runs TypeScript type-checking
yarn type-check
```

---

**Congratulations!** You now have VirtuozPlay up and running.
