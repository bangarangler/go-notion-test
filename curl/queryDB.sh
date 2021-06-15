#!/bin/zsh

curl -X POST 'https://api.notion.com/v1/databases/229f3080be484e9cbd9b649c2c045ec5/query' \
	-H "Authorization: Bearer $(cat /Users/jonathanpalacio/Desktop/go-notion-test/curl/token.txt)" \
  -H 'Notion-Version: 2021-05-13' \
  -H "Content-Type: application/json" \ -v | jq
	# --data '{
	#   "filter": {
	#     "or": [
	#       {
	#         "property": "In stock",
	# 				"checkbox": {
	# 					"equals": true
	# 				}
	#       },
	#       {
	# 				"property": "Cost of next trip",
	# 				"number": {
	# 					"greater_than_or_equal_to": 2
	# 				}
	# 			}
	# 		]
	# 	},
	#   "sorts": [
	#     {
	#       "property": "Last ordered",
	#       "direction": "ascending"
	#     }
	#   ]
	# }'
