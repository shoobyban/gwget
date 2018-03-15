# Wget based on file list

## Usage

```sh
source$ cd ~/web/media
source$ find . -type f | grep -v '.thumb' | grep -v product/cache | sed 's#^./##' > ~/media.txt
```

scp media.txt onto the target server with a gwget binary

```sh
target$ cd ~/web/media
target$ ~/gwget ~/media.txt https://www.somesite.com/media/
```
