{
  description = "Assembly Flake";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
  };

  outputs = {nixpkgs, ...}: let
    system = "x86_64-linux";
    pkgs = import nixpkgs {
      inherit system;
    };
  in
    with pkgs; {
      devShells."${system}".default = mkShell {
        packages = [
		  nasmfmt
          just
        ];

        buildInputs = [
		  binutils
          nasm
        ];

        shellHook = ''
          go version
        '';
      };
    };
}
