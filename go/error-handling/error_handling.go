package erratum

const testVersion = 2

func Use(opener ResourceOpener, input string) error {
	r, err := opener()
	for err != nil {
		if _, isTransient := err.(TransientError); isTransient {
			r, err = opener()
		} else {
			return err
		}
	}
	defer r.Close()

	err = frob(r, input)
	if err != nil {
		return err
	}

	return nil
}

func frob(r Resource, input string) (err error) {
	defer func() {
		if recovErr := recover(); recovErr != nil {
			if e, ok := recovErr.(FrobError); ok {
				r.Defrob(e.defrobTag)
			}
			err = recovErr.(error)
		}
	}()

	r.Frob(input)

	return nil
}
