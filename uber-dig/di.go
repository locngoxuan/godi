package main

func Initialize() error {
	err := container.Provide(func() string {
		return "Hi There"
	})
	if err != nil {
		return err
	}
	err = container.Provide(NewMessage)
	if err != nil {
		return err
	}
	err = container.Provide(NewEvent)
	if err != nil {
		return err
	}
	err = container.Provide(NewGreeter)
	if err != nil {
		return err
	}
	return nil
}
