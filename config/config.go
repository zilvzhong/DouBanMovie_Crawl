package config



var Defaults map[string]string

func init()  {
	Defaults = map[string]string{
		"doubian_url": "https://movie.douban.com/top250?start=%s&filter=",
	}
}

