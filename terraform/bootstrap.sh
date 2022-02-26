set -e -o pipefail

BUCKET_NAME="${BUCKET_NAME:?BUCKET_NAME is not set}"
gsutil mb -l asia-northeast1 "gs://$BUCKET_NAME" >&2
gsutil versioning set on "gs://$BUCKET_NAME" >&2
echo $BUCKET_NAME
