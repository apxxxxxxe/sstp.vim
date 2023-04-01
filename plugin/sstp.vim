if exists('g:loaded_sstp')
  finish
endif
let g:loaded_sstp = 1

command! -range SendSSTP <line1>,<line2>call sstp#sendrange()
