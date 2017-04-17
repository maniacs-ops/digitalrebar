#!/usr/bin/env bash

if which sudo 2>/dev/null >/dev/null ; then
    SUDO=sudo
fi

SED="sed -i"
if [[ $(uname -s) == Darwin ]] ; then
	SED="sed -i .bak"
fi

SERVICES="network-server logging-service"

function usage {
    echo "Usage: $0 <flags> [options docker-compose flags/commands]"
    echo "  -h or --help - help (this)"
    echo "  --clean - cleans up directory and exits"
    echo "  --dhcp # Adds the dhcp component"
    echo "  --provisioner # Adds the provisioner component"
    echo "  --dns # Adds the dns component"
    echo "  --ntp # Adds the ntp component"
    echo "  --chef # Adds the chef component"
    echo "  --webproxy # Adds the webproxy component"
    echo "  --revproxy # Adds the revproxy component"
    echo "  --logging # Adds the logging (kibana,elasticsearch+) components"
    echo "  --debug # Adds the cadviser components"
    echo "  --node # Adds the node component"
    echo "  --tag <TAG> # Uses that tag for builds and trees. default: latest"
    echo "  --external_ip <CIDR Address, default: 192.168.124.11/24> "
    echo
    echo " If additional arguments are provided, they are passed to docker-compose"
    echo " Otherwise nothing is run and just files are setup."
}

#
# Sets a value for a variable.
# The variable must exist.
#
function set_var_in_common_env {
  local var=$1
  local value=$2

  $SED -e "s/^${var}=.*/${var}=${value}/" common.env
}

FILES=(base.yml trust-me.yml)
REMOVE_FILES=""
ACCESS_MODE="HOST"
PROVISION_IT="NO"
PROXY_IT="NO"

if [[ -f tag ]]; then
    DR_TAG="$(cat tag)"
elif [[ ! $DR_TAG ]]; then
    branch="$(git symbolic-ref -q HEAD)"
    branch="${branch##refs/heads/}"
    branch="${branch:-latest}"
    DR_TAG="${DR_TAG:-${branch}}"
fi
ADD_DNS=false
RUN_NTP="NO"

while [[ $1 == -* ]] ; do
  arg=$1
  shift

  case $arg in
      --help)
      usage
      exit 0
      ;;
    -h)
      usage
      exit 0
      ;;
    --clean)
        rm -f access.env services.env dc docker-compose.yml config-dir/api/config/networks/the_admin.json config-dir/api/config/networks/the_bmc.json tag
        $SUDO rm -rf data-dir
      exit 0
      ;;
    --tag)
      DR_TAG=$1
      shift
      ;;
    --external_ip)
      EXTERNAL_IP=$1
      shift
      ;;
    --forwarder_ip)
      FORWARDER_IP=$1
      shift
      ;;
    --provisioner)
      FILES+=(provisioner.yml)
      PROVISION_IT="YES"
      SERVICES+=" provisioner-service"
      ;;
    --ntp)
      FILES+=(ntp.yml)
      SERVICES+=" ntp-service"
      RUN_NTP="YES"
      ;;
    --chef)
      FILES+=(chef.yml)
      SERVICES+=" chef-service"
      ;;
    --dhcp)
      FILES+=(dhcp.yml)
      SERVICES+=" dhcp-mgmt_service dhcp-service"
      ;;
    --dns)
      if [[ $ADD_DNS != true ]] ; then
          FILES+=(dns.yml)
      fi
      SERVICES+=" dns-service"
      ADD_DNS=true
      ;;
    --dns-mgmt)
      if [[ $ADD_DNS != true ]] ; then
          FILES+=(dns.yml)
      fi
      SERVICES+=" dns-mgmt_service"
      ADD_DNS=true
      ;;
    --webproxy)
      FILES+=(webproxy.yml)
      SERVICES+=" proxy-service"
      PROXY_IT="YES"
      ;;
    --revproxy)
      FILES+=(revproxy.yml)
      ;;
    --debug)
      FILES+=(debug.yml)
      ;;
    --logging)
      FILES+=(logging.yml)
      ;;
    --node)
      FILES+=(node.yml)
      ;;
  esac

done

EXTERNAL_IP=${EXTERNAL_IP:-192.168.99.100/24}
FORWARDER_IP=
MYPWD=`pwd`
cd config-dir/api/config/networks
rm -f the_admin.json || :
rm -f the_bmc.json || :
if [[ $PROVISION_IT = YES ]] ; then
    [[ -f the_admin.json.mac ]] && ln -s the_admin.json.mac the_admin.json
fi
cd $MYPWD

if [[ -x ../../go/bin/$DR_TAG/linux/amd64/rebar ]]; then
    mkdir -p data-dir/bin
    cp "../../go/bin/$DR_TAG/linux/amd64/"* data-dir/bin
fi

# Process templates and build one big yml file for now.
rm -f docker-compose.yml
for i in "${FILES[@]}"; do
    fname="$i"
    if [[ $i != /* ]] ; then
        fname="yaml_templates/$i"
    fi
    # Fix Access Mode
    cat "$fname" >> docker-compose.yml
done

if [[ ! $DEV_MODE = Y ]]; then
    $SED -e "/START DEV_MODE/,/END DEV_MODE/d" docker-compose.yml
fi
$SED -e "/ACCESS_MODE/d" -e '/DEV_MODE/d' docker-compose.yml

if [[ $REMOVE_FILES ]] ; then
	rm -f $REMOVE_FILES
fi

# Find the IP address we should have Consul advertise on
if [[ $(uname -s) == Darwin ]]; then
    CONSUL_ADVERTISE=${DOCKER_HOST%:*}
    CONSUL_ADVERTISE=${CONSUL_ADVERTISE##*/}
elif [[ $(uname -s) == "MINGW64_NT-10.0" ]]; then
    CONSUL_ADVERTISE=${DOCKER_HOST%:*}
    CONSUL_ADVERTISE=${CONSUL_ADVERTISE##*/}
else
    gwdev=$(/sbin/ip -o -4 route show default |head -1 |awk '{print $5}')
    if [[ $gwdev ]]; then
        # First, advertise the address of the device with the default gateway
        CONSUL_ADVERTISE=$(/sbin/ip -o -4 addr show scope global dev "$gwdev" |head -1 |awk '{print $4}')
        CONSUL_ADVERTISE="${CONSUL_ADVERTISE%/*}"
    else
        # Hmmm... we have no access to the Internet.  Pick an address with
        # global scope and hope for the best.
        CONSUL_ADVERTISE=$(/sbin/ip -o -4 addr show scope global |head -1 |awk '{print $4}')
        CONSUL_ADVERTISE="${CONSUL_ADVERTISE%/*}"
    fi
fi
# If we did not get and address to listen on, we are pretty much boned anyways
if [[ ! $CONSUL_ADVERTISE ]]; then
    echo "Could not find an address for Consul to listen on!"
    exit 1
fi
# CONSUL_JOIN is separate from CONSUL_ADVERTISE as futureproofing
CONSUL_JOIN="$CONSUL_ADVERTISE"
# Make access.env for Variables.
cat >access.env <<EOF
USE_OUR_PROXY=$PROXY_IT
EXTERNAL_IP=$EXTERNAL_IP
FORWARDER_IP=$FORWARDER_IP
CONSUL_JOIN=$CONSUL_JOIN
DR_START_TIME=$(date +%s)
RUN_NTP=$RUN_NTP
EOF

# Add proxies from this environment to the containers.
# Need to do similar things
if [[ $http_proxy && $PROXY_IT = NO ]]; then
    cat >>access.env <<EOF
UPSTREAM_HTTP_PROXY=$http_proxy
UPSTREAM_HTTPS_PROXY=$https_proxy
UPSTREAM_NO_PROXY=$no_proxy
EOF
fi

cat >services.env <<EOF
SERVICES=$SERVICES
EOF

cat >config-dir/consul/server-advertise.json <<EOF
{"advertise_addr": "${CONSUL_ADVERTISE}"}
EOF

echo "$DR_TAG" >tag

# With remaining arguments
if [ "$#" -gt 0 ] ; then
    docker-compose $@
fi
