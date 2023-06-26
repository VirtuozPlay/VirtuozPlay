# Setup a fully-featrued development environment for Buffalo and Vue
{
  inputs = {
    devenv.url = "github:cachix/devenv";
    flake-utils.url = "github:numtide/flake-utils";
    nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";
  };

  outputs = { self, devenv, flake-utils, nixpkgs, ... } @ inputs:
  flake-utils.lib.eachDefaultSystem (system:
  let
    pkgs = nixpkgs.legacyPackages.${system};
    buffalo-version = "0.18.4";
  in
  {
    devShells.default = devenv.lib.mkShell {
      inherit inputs pkgs;

      modules = [
        ({ lib, ... }: let
          neededPackages = with pkgs; [
            alsa-lib
            alsaLib
            atk
            at-spi2-atk
            at-spi2-core
            autoPatchelfHook
            autoreconfHook
            cairo
            cups
            dbus
            expat
            ffmpeg
            gcc
            gdk-pixbuf
            glib
            gnome2.GConf
            gtk2
            gtk3
            libdrm
            libxkbcommon
            mesa
            nodejs
            nspr
            nss
            pango
            unzip
            xorg.libX11
            xorg.libxcb
            xorg.libXcomposite
            xorg.libXcursor
            xorg.libXdamage
            xorg.libXext
            xorg.libXfixes
            xorg.libXi
            xorg.libXrandr
            xorg.libXrender
            xorg.libXScrnSaver
            xorg.libXtst
            xorg_sys_opengl
            xorg.xauth
            yarn
            (lib.getLib udev)
          ];   
        in
        {
          packages = neededPackages;

          env.NIX_LD = pkgs.lib.fileContents "${pkgs.stdenv.cc}/nix-support/dynamic-linker";
          env.NIX_LD_LIBRARY_PATH="${pkgs.lib.makeLibraryPath neededPackages}";

          languages.go.enable = true;
          languages.javascript.enable = true;
          languages.typescript.enable = true;

          services.postgres = {
            enable = true;

            initialDatabases = [
              {
                name = "postgres";
              }
            ];

            initialScript = ''
                CREATE USER postgres WITH SUPERUSER PASSWORD 'postgres';
            '';

            settings = {
              listen_addresses = lib.mkForce "*";
            };
          };

          # Run `install-buffalo` in your shell to install the Buffalo CLI
          scripts.install-buffalo.exec = ''
            go install -tags sqlite github.com/gobuffalo/cli/cmd/buffalo@v${buffalo-version}
          '';

          enterShell = ''
            export buildInputs="$buildInputs"
            install-buffalo
          '';
        })
      ];
    };
  });
}
