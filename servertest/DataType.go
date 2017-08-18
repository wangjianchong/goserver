package main

type (
	Login struct {
		Username string
		Pwd      string
	}
	Music struct {
		Name      string
		Url       string
		Discuss   string
		Favourite string
		MusicId   string
	}
	Test struct {
		Username string
		Pwd      string
	}
	Contact struct {
		Phone  string
		Wechat string
		QICQ   string
	}
	Addr struct {
		Province string
		City     string
		Region   string
		Street   string
	}
)
