package vault_collector

type vaultCollector struct {
}

func (*vaultCollector) Collect() {

}

func New() *vaultCollector {
	return &vaultCollector{}
}
