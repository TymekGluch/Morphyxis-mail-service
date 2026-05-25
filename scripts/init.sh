set -e

cp ./scripts/prePush.sh .git/hooks/pre-push
chmod +x .git/hooks/pre-push

echo "Pre-push hook installed successfully!"
