curl -X 'POST' \
  'https://dct101.dlpxdc.co/v3/management/engines' \
  -H 'accept: application/json' \
  -H 'Authorization: apk 1.pvR2JlMe9MEWHHV38yhNPbGMOHV9W1R2iiGYguXXSgskSIlAlyeNxiDmESFGNBLC' \
  -H 'Content-Type: application/json' \
  -d '{ 
    "name": "dlpx-hckthn",    
    "hostname": "dlpx-hckthn.dlpxdc.co",
    "username": "admin",
    "password": "delphix",
    "insecure_ssl": true
}' -o output.json -k | jq '.'