if exists('g:autoloaded_sstp')
  finish
endif
let g:autoloaded_sstp = 1

let s:sep = fnamemodify('.', ':p')[-1:]
let s:cmd_dir = expand('<script>:p:h:h') .. s:sep .. 'bin'

function! sstp#sendrange() range
  let l:has_wsl = !has('win32') && exists('$WSLENV')
  if !l:has_wsl && !has('win32')
    echoerr 'sstp#sendrange() is only supported on Windows or WSL'
    return
  endif
  let l:cmd = 'send_sstp_windows_amd64.exe'
  let l:selected = join(getline(a:firstline, a:lastline), "\\n")
  let l:arg = substitute(l:selected, "\"", "\\\"", "g")
  let l:arg = substitute(l:arg, "\'", "\\\"", "g")
  let l:arg = substitute(l:arg, "\\", "\\\\", "g")
  echo system(s:cmd_dir .. s:sep .. l:cmd .. ' "' .. l:arg .. '"')
endfunction
