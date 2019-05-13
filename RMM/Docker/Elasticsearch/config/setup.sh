#!/bin/sh

echo Initiating Elasticsearch Custom Index
# move to the directory of this setup script
cd "$(dirname "$0")"

# for some reason even when port 9200 is open Elasticsearch is unable to be accessed as authentication fails
# a few seconds later it works
until $(curl -sSf -XGET --insecure --user elastic:changeme 'http://localhost:9200/_cluster/health?wait_for_status=yellow' > /dev/null); do
    printf 'AUTHENTICATION ERROR DUE TO X-PACK, trying again in 5 seconds \n'
    sleep 5
done

# create a new template for managed_endpoint indexes
curl -v --insecure --user elastic:changeme -XPUT '0.0.0.0:9200/_template/template_managed_endpoint' -H 'Content-Type: application/json' -d @me.json

# create a new template for installed_software_tree indexes
curl -v --insecure --user elastic:changeme -XPUT '0.0.0.0:9200/_template/template_installed_software' -H 'Content-Type: application/json' -d @is.json

# create a new template for os_tree indexes
curl -v --insecure --user elastic:changeme -XPUT '0.0.0.0:9200/_template/template_operating_system' -H 'Content-Type: application/json' -d @os.json