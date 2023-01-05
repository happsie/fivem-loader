package updater

type Updater interface {
	Update(name, url, destinationFolder string, skipConfig, forceUpdate bool) error
}
