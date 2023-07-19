# paperops


## Overleaf

1. go into the overleaf dir
2. bin/init
3. change `overleaf.rc`'s SHARELATEX_LISTEN_IP
4. run `bin/shell`, and input `tlmgr option repository https://mirrors.tuna.tsinghua.edu.cn/CTAN/systems/texlive/tlnet`
5. tlmgr install scheme-full

or install all the [dependencies](https://github.com/overleaf/overleaf/wiki/Quick-Start-Guide#latex-environment):

```bash
apt update
apt -y upgrade
apt install texlive-full -y
apt -y install texlive-bibtex-extra
apt -y install texlive-xetex
apt -y install context
wget https://mirror.ctan.org/systems/texlive/tlnet/update-tlmgr-latest.sh update-tlmgr-latest.sh
sh update-tlmgr-latest.sh
tlmgr update --self --all
tlmgr install scheme-full
tlmgr install biblatex
tlmgr install biber
mtxrun --generate
mktexlsr
updmap-sys
fmtutil-sys --all
apt -y install ttf-mscorefonts-installer
```