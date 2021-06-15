#!/bin/zsh

# echo $(cat /Users/jonathanpalacio/Desktop/go-notion-test/curl/token.txt)

curl 'https://api.notion.com/v1/pages/10c14d577a53437397528e3e8a4ee404' \
  -H 'Notion-Version: 2021-05-13' \
  -H "Authorization: Bearer $(cat /Users/jonathanpalacio/Desktop/go-notion-test/curl/token.txt)" \
  -v | jq

