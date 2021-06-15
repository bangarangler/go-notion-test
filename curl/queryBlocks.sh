#!/bin/zsh

curl 'https://api.notion.com/v1/blocks/10c14d577a53437397528e3e8a4ee404/children?page_size=100' \
  -H "Authorization: Bearer $(cat /Users/jonathanpalacio/Desktop/go-notion-test/curl/token.txt)" \
  -H "Notion-Version: 2021-05-13" \
  -v | jq
