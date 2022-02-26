apk add curl jq

i=0

# shellcheck disable=SC2153
# shellcheck disable=SC2001
for API_URL in $(echo "$API_URLS" | sed "s/,/ /g"); do
  i=$(($i+1))
  curl -X GET $API_URL -s -f > /openapi2krakend/swagger/$i.json
  if [ -s /openapi2krakend/swagger/$i.json ]; then
    echo "Downloaded $API_URL..."
    cat /openapi2krakend/swagger/$i.json
  else
    echo "$API_URL not found!"
    rm -f /openapi2krakend/swagger/$i.json
  fi
done

cd openapi2krakend || exit
./openapi2krakend
