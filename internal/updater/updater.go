package updater

type Updater interface {
	Update(name, url string) error
}