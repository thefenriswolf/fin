{ stdenvNoCC, lib, fetchFromGitHub, makeWrapper, buildGoModule }:

buildGoModule rec {
  pname = "gopta";
  version = "v20240331";

  src = fetchFromGitHub {
    owner = "thefenriswolf";
    repo = "fin";
    rev = "${version}";
    hash = "sha256-vZcCOl2V6sP10nbdiaF85CiPJaPY+QG+ERKRAbPx3yY=";
  };
  vendorHash = "";

  meta = with lib; {
    description = "Time Tracker Tool written in Golang";
    homepage = "https://github.com/thefenriswolf/ttt";
    license = licenses.mit;
    maintainers = with maintainers; [ thefenriswolf ];
  };
}
