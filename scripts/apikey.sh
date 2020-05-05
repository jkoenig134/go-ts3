read -p "host: " HOST
read -p "username: " USER
read -p "password: " PASSWORD

nc -q1 $HOST 10011 <<EOF
login $USER $PASSWORD
apikeyadd scope=manage lifetime=0
EOF
