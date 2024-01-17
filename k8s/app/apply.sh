function main() {
  local script_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
  #kubectl apply -f "$script_dir/ubuntu-test.yaml"
  kubectl apply -f "$script_dir/secret.yaml"
  kubectl apply -f "$script_dir/jobs.yaml"
  kubectl apply -f "$script_dir/deployment.yaml"
  kubectl apply -f "$script_dir/service.yaml"
}
main