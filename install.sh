#!/bin/bash

# Mengganti nama binary dengan nama proyek Anda
binary_name="xerador"

# Membuat direktori instalasi
install_dir="/usr/local/bin"

# Menyalin file binary ke direktori instalasi
cp "$binary_name" "$install_dir/"

echo "Finish installation. run with '$binary_name'."
