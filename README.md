# PSEUDOCODE

## struktur

```go
program <nama program>
...
kamus
...
algoritma
...
endprogram
```

## `kamus`

### deklarasi variable

untuk mendeklarasikan variable gunakan

```go
{var_name} : {var_type}
```

tipe variable yang tersedia

| pseudocode | golang |
| --- | --- |
| boolean | bool |
| string | string |
| char | rune |
| integer | int |
| real | float64, float32 |

### constant, type declaration, dan alias var

constant variable gunakan  

```go
constant {var_name} : {var_type} <- {value} # catatan gunakan selalu "<-" di pseudocode untuk insert value
```

type declaration gunakan

```go
type {var_name} <
  {var_name} : {var_type}
  {var_name} : {var_type}
  {var_name} : {var_type}
>
```

alias variable gunakan

```go
{new_var_alias} : {exist_var_dec}
```

### CONTOH `kamus`

```go
kamus
  type waktu <
    jam   : int
    menit : int
    detik : int
  >
  constant jumlah_menit : int <- 1234 
  b1, b2, b3  : integer
```

## `algoritma`

### input dan output

```go
input(var_name)
output(var_name)
```

### insert value ke variable

```go
{var_name} <- {var_value}
```

### `MOD` dan `DIV`

`MOD` <- sisa bagi  
`DIV` <- berapa kali bilangan ini dibagi sama bilangan pembagi  
selalu gunakan `MOD` dan `DIV` untuk melakukan perhitungan di pseudocode bukan menggunakan `%` dan `/`

```go
bil1 <- 100 MOD 8
bil2 <- 100 DIV 8
```

### operator logic

| logic | pseudocode | golang (logical) | golang (bitwise/comparison) |
| --- | --- | --- | --- |
| NOT | NOT | `!` | `!=` (Not Equal) |
| AND | AND | `&&` | `&` (Bitwise AND) |
| OR | OR | `\|\|` | `\|` (Bitwise OR) |
| XOR | XOR | `!=` | `^` (Bitwise XOR) |

### FOR LOOP

`i` yang ada di perulangan `for` bisa digunakan di dalam perulangan  
cth for loop dari 1 ke 9

```go
for i <- 1 to 9 do
  ...
endfor
```

cth for loop dari 9 ke 1

```go
for i <- 9 downto 1 do 
  ...
endfor
```

### IF ELSE

tunggu diajarin dulu biar gak salah dgn dosen
