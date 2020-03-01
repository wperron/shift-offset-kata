# usage: echo '5 7 9 1 3 4' | tr ' ' '\n' | awk -f shift_offset.awk
$0 > last {last = $0; lastrow = NR}
END {print lastrow}