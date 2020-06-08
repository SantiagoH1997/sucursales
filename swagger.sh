BASEDIR=$(dirname "$0")
GO111MODULE=off swagger generate spec -o $BASEDIR/swagger/swagger.yaml --scan-models 