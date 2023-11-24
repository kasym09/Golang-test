curl -s https://zero.academie.one/assets/superhero/all.json curl -s https://zero.academie.one/assets/superhero/all.json | jq -r '.[] | select(.id==170)| .name'

