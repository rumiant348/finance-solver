package models

type listValidator struct {
	ListDB
}

func (lv *listValidator) Create(list *List) error {
	err := runListValidatorFns(list,
		lv.titleNotEmpty,
	)
	if err != nil {
		return err
	}
	return lv.ListDB.Create(list)
}

func (lv *listValidator) Update(list *List) error {
	err := runListValidatorFns(list,
		lv.titleNotEmpty,
	)
	if err != nil {
		return err
	}
	return lv.ListDB.Update(list)
}

func (lv *listValidator) Delete(id uint) error {
	var list List
	list.ID = id

	err := runListValidatorFns(&list,
		lv.idNotNull,
		lv.titleNotEmpty,
	)
	if err != nil {
		return err
	}
	return lv.ListDB.Delete(id)
}

func (lv *listValidator) titleNotEmpty(list *List) error {
	if list.Title == "" {
		return ErrTitleEmpty
	}
	return nil
}

func (lv *listValidator) idNotNull(list *List) error {
	if list.ID <= 0 {
		return ErrIDInvalid
	}
	return nil
}

type listValFn func(*List) error

func runListValidatorFns(list *List, fns ...listValFn) error {
	for _, fn := range fns {
		if err := fn(list); err != nil {
			return err
		}
	}
	return nil
}
