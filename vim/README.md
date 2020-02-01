# VIM 

Mencoba kembali mensetup vim yang dulu sempat dicoba tapi tidak nyaman. Akhirnya balik lagi ke VSCode. Melihat kawan dikantor code menggunakan vim jadi kembali ingin mencoba tapi dengan setup yang lebih proper. 

Feature:

Fitur fitur yang diinginkan ada di VIM

- [x] Go Language
- [x] Syntax higlighting
- [x] Tree Menu
- [x] Linter
- [ ] Autocomplete
- [ ] Extract Method
- [ ] Unit test generator

Keymap:

Keymap berdasarkan kebiasaan di VSCode, Jika tidak bisa disamakan paling tidak tau gantinya apa.

- [ ] SHIFT + up / down                    : untuk memblok beberapa baris
- [ ] Key map: SHIFT + CMD  + left / right : untuk memblok baris kode
- [ ] Key map: OPT + up / down             : untuk memindahkan baris kode ke atas atau bawah  
- [ ] Key map: OPT + CMD + up and down     : untuk membuat kursor ke beberapa baris 
- [ ] Key map: CMD + P                     : untuk memilih file
- [ ] Key map: CMD + W                     : untuk menghapus tab

## Setup

Install [vim-plug](https://github.com/junegunn/vim-plug). Rekomendasi dari @ri7nz <ri7nz.labs@gmail.com>

```bash
set encoding=utf8

if empty(glob('~/.vim/autoload/plug.vim'))
  silent !curl -fLo ~/.vim/autoload/plug.vim --create-dirs
    \ https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim
  autocmd VimEnter * PlugInstall --sync | source $MYVIMRC
endif

call plug#begin('~/.vim/plugged')

...

call plug#end()

"Basic Config
"

" Autowrite = auto save when type such as :make or :GoBuild
set autowrite

" Using backspace for removing
set bs=2

" Set show matching bracket
set showmatch

" Always shows location in file (line #)
set ruler

" Tabs
set ts=4
set sw=4
set smarttab
```
