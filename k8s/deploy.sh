#!/bin/bash

# Betiğin bulunduğu klasörün dizini
script_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
echo "🚀🚀🚀 PostgresDB going to initialize cluster 🚀🚀🚀"
source "$script_dir/postgres/apply.sh"
sleep 20 #Give some time to pod bootstrap because i dont have starter checks.
echo "🚀🚀🚀 Redis going to initialize cluster 🚀🚀🚀"
source "$script_dir/redis/apply.sh"
sleep 20 #Give some time to pod bootstrap because i dont have starter checks.
echo "🚀🚀🚀 Application going to initialize cluster 🚀🚀🚀"
source "$script_dir/app/apply.sh"

