umount /mnt/orangepi
sshfs -o allow_other,default_permissions root@192.168.1.94:/ /mnt/orangepi
GOARCH=arm64 GOOS=linux go build -ldflags "-s -w" -o app .
cp app /mnt/orangepi/root/.gam-app/maldan-gam-app-desktop-v0.0.11