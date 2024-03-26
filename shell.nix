{ pkgs ? import <nixpkgs> { } }:

pkgs.mkShell {
  buildInputs = [
    pkgs.beancount-black
    pkgs.beancount-language-server
    pkgs.beancount
    pkgs.fava
    # keep this line if you use bash
    pkgs.bashInteractive
  ];
}
