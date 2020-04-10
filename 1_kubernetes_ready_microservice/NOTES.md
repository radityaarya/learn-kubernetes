# kompose installation
```
# Linux
curl -L https://github.com/kubernetes/kompose/releases/download/v1.20.0/kompose- linux-amd64 -o kompose

# macOS
curl -L https://github.com/kubernetes/kompose/releases/download/v1.20.0/kompose- darwin-amd64 -o kompose

chmod +x kompose
sudo mv ./kompose /usr/local/bin/kompose
```

autocompletion
```
# Bash (add to .bashrc for persistence)
source <(kompose completion bash)

# Zsh (add to .zshrc for persistence)
source <(kompose completion zsh)
```

Convert the Compose file into Kubernetes resources files
```
kompose convert -f docker-compose.yml
```
