#!/bin/sh -e

CSVQ=/bin/csvq
VANS=/bin/vans

if [ $# -eq 0 ] ; then
    echo "$0 -csvfile <csvfile> -s <inputfile> -d <outputfile>"
    exit 1
fi

CSV_FILE=
VANS_OPTS=
while [[ $# -gt 0 ]] ; do
  case $1 in
    -csvfile)
      CSV_FILE=$2
      shift 2
      ;;
    *)
      VANS_OPTS="$VANS_OPTS $1"
      shift
      ;;
  esac
done

mkdir -p /tmp/rcsv
cp -f $CSV_FILE /tmp/rcsv/nodeconf.csv

CSV_ENV=$( $CSVQ -r /tmp/rcsv -f LTSV "select * from nodeconf where node = \"$NODE_NAME\"" )
CSV_ENV=$( echo $CSV_ENV | sed 's/:/=/g' )
if [ "$CSV_ENV" != "" ]; then
    export $CSV_ENV
fi

$VANS $VANS_OPTS
