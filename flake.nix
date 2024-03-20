{
  description = "A simple Go package";

  # Nixpkgs / NixOS version to use.
  #inputs.nixpkgs.url = "nixpkgs/nixos-21.11";

  #inputs = {
  #  nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";
  #  flake-utils.url = "github:numtide/flake-utils";
  #};

  outputs = { self, nixpkgs }:
    let

      # to work with older version of flakes
      lastModifiedDate =
        self.lastModifiedDate or self.lastModified or "19700101";
      # Generate a user-friendly version number.
      version = builtins.substring 0 12 lastModifiedDate;

      # System types to support.
      supportedSystems =
        [ "x86_64-linux" "x86_64-darwin" "aarch64-linux" "aarch64-darwin" ];

      # Helper function to generate an attrset '{ x86_64-linux = f "x86_64-linux"; ... }'.
      forAllSystems = nixpkgs.lib.genAttrs supportedSystems;

      # Nixpkgs instantiated for supported system types.
      nixpkgsFor = forAllSystems (system: import nixpkgs { inherit system; });

    in {
      # Provide some binary packages for selected system types.
      # `nix build`
      packages = forAllSystems (system:
        let pkgs = nixpkgsFor.${system};
        in {
          wechat_hook = pkgs.buildGoModule {
            pname = "wechat_hook";
            inherit version;
            # In 'nix develop', we don't need a copy of the source tree
            # in the Nix store.
            src = ./.;

            # This hash locks the dependencies of this package. It is
            # necessary because of how Go requires network access to resolve
            # VCS.  See https://www.tweag.io/blog/2021-03-04-gomod2nix/ for
            # details. Normally one can build with a fake sha256 and rely on native Go
            # mechanisms to tell you what the hash should be or determine what
            # it should be "out-of-band" with other tooling (eg. gomod2nix).
            # To begin with it is recommended to set this, but one must
            # remember to bump this hash when your dependencies change.
            #vendorSha256 = pkgs.lib.fakeSha256;
            doCheck = false;
            vendorHash = "sha256-kMhEs7CdcGtyyurW4ufptoEIZ8GmcTGGl1JUjxk84us=";
          };
          # Add entry to build a docker image with wechat_hook
          # caveat: only works on Linux
          #
          # Usage:
          # nix build .#wechat_hook-docker
          # docker load < result
          wechat_hook-docker = pkgs.dockerTools.buildLayeredImage {
            name = "wechat_hook";
            tag = version;
            contents = [ pkgs.cacert self.packages.${system}.wechat_hook ];
            config.Entrypoint = [ "/bin/wechat_hook" ];
            config.Created = builtins.substring 0 12 lastModifiedDate;
          };
        });

      # Add dependencies that are only needed for development
      # `nix develop`
      devShells = forAllSystems (system:
        let pkgs = nixpkgsFor.${system};
        in {
          default = pkgs.mkShell {
            buildInputs = with pkgs; [ go gopls gotools go-tools ];
          };
        });

      # The default package for 'nix build'. This makes sense if the
      # flake provides only one package or there is a clear "main"
      # package.
      defaultPackage =
        forAllSystems (system: self.packages.${system}.wechat_hook-docker);
    };
}
