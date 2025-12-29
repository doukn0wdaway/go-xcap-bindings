{
  description = "golang xcap bindings dev flake";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-25.11";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { nixpkgs, flake-utils, ... }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs { inherit system; };
      in
      {

        devShells.default = pkgs.mkShell {



          buildInputs = [
            pkgs.rustc
            pkgs.pkg-config
            pkgs.cargo
            pkgs.wayland

            pkgs.mesa
            pkgs.go
            pkgs.libgbm
            pkgs.pipewire
            pkgs.libGL
            pkgs.clang

            pkgs.libxcb
          ];


          # export LIBCLANG_PATH=${pkgs.llvmPackages_20.libclang.lib}/lib
          shellHook = ''
            export CC=clang
            export CXX=clang++
          '';
        };
      });
}
