package env

import "os"

func GetLeaderElectionNamespace() (string, bool) {
	return os.LookupEnv("LEADER_ELECTION_NAMESPACE")
}
